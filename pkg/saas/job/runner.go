package job

import (
	"context"
	"time"

	"github.com/CoopHive/coophive/pkg/data"
	"github.com/CoopHive/coophive/pkg/jobcreator"
	optionsfactory "github.com/CoopHive/coophive/pkg/options"
	"github.com/CoopHive/coophive/pkg/saas/types"
	coophivesystem "github.com/CoopHive/coophive/pkg/system"
	"github.com/CoopHive/coophive/pkg/web3"
	"github.com/spf13/cobra"
)

type JobRunner struct {
	Ctx        context.Context
	Options    jobcreator.JobCreatorOptions
	Web3SDK    *web3.Web3SDK
	JobCreator *jobcreator.JobCreator
	ErrorChan  chan error
}

func NewJobRunner(ctx context.Context) (*JobRunner, error) {
	// get options without a job to bootstrap the sdk and jobcreator
	options, err := GetJobOptions(types.JobSpec{
		Module: "",
		Inputs: map[string]string{},
	}, false)
	if err != nil {
		return nil, err
	}
	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return nil, err
	}
	jobCreatorService, err := jobcreator.NewJobCreator(options, web3SDK)
	if err != nil {
		return nil, err
	}
	tmpCommand := &cobra.Command{}
	tmpCommand.SetContext(ctx)
	cmdCtx := coophivesystem.NewCommandContext(tmpCommand)

	jobCreatorErrors := jobCreatorService.Start(cmdCtx.Ctx, cmdCtx.Cm)

	// wait a short period because we've just started the job creator service
	time.Sleep(100 * time.Millisecond)
	return &JobRunner{
		Options:    options,
		Web3SDK:    web3SDK,
		JobCreator: jobCreatorService,
		ErrorChan:  jobCreatorErrors,
	}, nil
}

func getCommandContext(ctx context.Context) *coophivesystem.CommandContext {
	tmpCommand := &cobra.Command{}
	tmpCommand.SetContext(ctx)
	return coophivesystem.NewCommandContext(tmpCommand)
}

func (runner *JobRunner) Subscribe(ctx context.Context, callback jobcreator.JobOfferSubscriber) {
	runner.JobCreator.SubscribeToJobOfferUpdates(callback)
}

func (runner *JobRunner) GetJobOffer(ctx context.Context, request types.JobSpec) (data.JobOffer, error) {
	options, err := GetJobOptions(request, true)
	if err != nil {
		return data.JobOffer{}, err
	}
	return runner.JobCreator.GetJobOfferFromOptions(options.Offer)
}

func (runner *JobRunner) GetJobContainer(ctx context.Context, request types.JobSpec) (data.JobOfferContainer, error) {
	jobOffer, err := runner.GetJobOffer(ctx, request)
	if err != nil {
		return data.JobOfferContainer{}, err
	}
	id, err := data.GetJobOfferID(jobOffer)
	if err != nil {
		return data.JobOfferContainer{}, err
	}
	jobOffer.ID = id
	container := data.GetJobOfferContainer(jobOffer)
	return container, nil
}

func (runner *JobRunner) RunJobSync(ctx context.Context, request types.JobSpec, handler func(evOffer data.JobOfferContainer)) (*jobcreator.RunJobResults, error) {
	commandCtx := getCommandContext(ctx)
	options := optionsfactory.NewJobCreatorOptions()

	options.Offer.Module.Name = request.Module
	options.Offer.Inputs = request.Inputs

	options, err := optionsfactory.ProcessJobCreatorOptions(options, []string{})
	if err != nil {
		return nil, err
	}

	return jobcreator.RunJob(commandCtx, options, handler)
}

func (runner *JobRunner) RunJobAsync(ctx context.Context, request types.JobSpec) (string, error) {
	jobOffer, err := runner.GetJobOffer(ctx, request)
	if err != nil {
		return "", err
	}
	container, err := runner.JobCreator.AddJobOffer(jobOffer)
	if err != nil {
		return "", err
	}
	return container.DealID, nil
}
