package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/CoopHive/coophive/pkg/data"
	"github.com/CoopHive/coophive/pkg/jobcreator"
	jobutils "github.com/CoopHive/coophive/pkg/saas/job"
	"github.com/CoopHive/coophive/pkg/saas/store"
	"github.com/CoopHive/coophive/pkg/saas/system"
	"github.com/CoopHive/coophive/pkg/saas/types"
	"github.com/davecgh/go-spew/spew"
)

func (c *Controller) GetStatus(ctx types.RequestContext) (types.UserStatus, error) {
	balanceTransfers, err := c.Options.Store.GetBalanceTransfers(ctx.Ctx, store.OwnerQuery{
		Owner:     ctx.Owner,
		OwnerType: ctx.OwnerType,
	})
	if err != nil {
		return types.UserStatus{}, err
	}

	// add up the total value of all balance transfers
	credits := 0
	for _, balanceTransfer := range balanceTransfers {
		credits += balanceTransfer.Amount
	}
	return types.UserStatus{
		User:    ctx.Owner,
		Credits: credits,
	}, nil
}

func (c *Controller) GetJobs(ctx types.RequestContext) ([]*types.Job, error) {
	return c.Options.Store.GetJobs(ctx.Ctx, store.GetJobsQuery{
		Owner:     ctx.Owner,
		OwnerType: ctx.OwnerType,
	})
}

func (c *Controller) GetTransactions(ctx types.RequestContext) ([]*types.BalanceTransfer, error) {
	return c.Options.Store.GetBalanceTransfers(ctx.Ctx, store.OwnerQuery{
		Owner:     ctx.Owner,
		OwnerType: ctx.OwnerType,
	})
}

func (c *Controller) checkUserFundsForJob(ctx types.RequestContext, request types.JobSpec) error {
	// TODO: we should make checking for funds, running the job and then billing
	// the user in some sort of transaction so that user can't submit 100 jobs
	// in parallel and end up with a negative balance
	_, err := jobutils.GetModule(request.Module)
	if err != nil {
		return err
	}
	// module.Cost is negative, so we add it to the user's balance
	return nil
}

func (c *Controller) billUserForJob(ctx types.RequestContext, jobID string, request types.JobSpec) error {
	module, err := jobutils.GetModule(request.Module)
	if err != nil {
		return err
	}
	err = c.Options.Store.CreateBalanceTransfer(ctx.Ctx, types.BalanceTransfer{
		ID:          system.GenerateUUID(),
		Owner:       ctx.Owner,
		OwnerType:   ctx.OwnerType,
		PaymentType: types.PaymentTypeJob,
		Amount:      -module.Cost,
		Data: types.BalanceTransferData{
			JobID: jobID,
		},
	})

	return err
}

func (c *Controller) CreateJobAsync(ctx types.RequestContext, request types.JobSpec) (string, error) {
	err := c.checkUserFundsForJob(ctx, request)
	if err != nil {
		return "", err
	}
	// TODO: charge them money
	id, err := c.Options.JobRunner.RunJobAsync(ctx.Ctx, request)
	if err != nil {
		return "", err
	}
	err = c.billUserForJob(ctx, id, request)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) CreateJobSync(ctx types.RequestContext, request types.JobSpec) (*jobcreator.RunJobResults, error) {
	err := c.checkUserFundsForJob(ctx, request)
	if err != nil {
		return nil, err
	}

	result, err := c.Options.JobRunner.RunJobSync(ctx.Ctx, request, func(evOffer data.JobOfferContainer) {
		fmt.Printf("evOffer --------------------------------------\n")
		spew.Dump(evOffer)
	})
	if err != nil {
		return nil, err
	}

	err = c.billUserForJob(ctx, result.JobOffer.DealID, request)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *Controller) CreateAPIKey(ctx types.RequestContext, name string) (string, error) {
	apiKey, err := c.Options.Store.CreateAPIKey(ctx.Ctx, store.OwnerQuery{
		Owner:     ctx.Owner,
		OwnerType: ctx.OwnerType,
	}, name)
	if err != nil {
		return "", err
	}
	return apiKey, nil
}

func (c *Controller) GetAPIKeys(ctx types.RequestContext) ([]*types.ApiKey, error) {
	apiKeys, err := c.Options.Store.GetAPIKeys(ctx.Ctx, store.OwnerQuery{
		Owner:     ctx.Owner,
		OwnerType: ctx.OwnerType,
	})
	if err != nil {
		return nil, err
	}
	if apiKeys == nil {
		_, err := c.CreateAPIKey(ctx, "default")
		if err != nil {
			return nil, err
		}
		return c.GetAPIKeys(ctx)
	}
	return apiKeys, nil
}

func (c *Controller) DeleteAPIKey(ctx types.RequestContext, apiKey string) error {
	fetchedApiKey, err := c.Options.Store.CheckAPIKey(ctx.Ctx, apiKey)
	if err != nil {
		return err
	}
	if fetchedApiKey == nil {
		return errors.New("no such key")
	}
	// only the owner of an api key can delete it
	if fetchedApiKey.Owner != ctx.Owner || fetchedApiKey.OwnerType != ctx.OwnerType {
		return errors.New("unauthorized")
	}
	err = c.Options.Store.DeleteAPIKey(ctx.Ctx, *fetchedApiKey)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) CheckAPIKey(ctx context.Context, apiKey string) (*types.ApiKey, error) {
	key, err := c.Options.Store.CheckAPIKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}
	return key, nil
}
