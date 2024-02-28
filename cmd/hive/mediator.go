package hive

import (
	"github.com/CoopHive/coophive/pkg/executor/bacalhau"
	"github.com/CoopHive/coophive/pkg/mediator"
	optionsfactory "github.com/CoopHive/coophive/pkg/options"
	"github.com/CoopHive/coophive/pkg/system"
	"github.com/CoopHive/coophive/pkg/web3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newMediatorCmd() *cobra.Command {
	options := optionsfactory.NewMediatorOptions()

	mediatorCmd := &cobra.Command{
		Use:     "mediator",
		Short:   "Start the hive mediator service.",
		Long:    "Start the hive mediator service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := optionsfactory.ProcessMediatorOptions(options)
			if err != nil {
				return err
			}
			return runMediator(cmd, options)
		},
	}

	optionsfactory.AddMediatorCliFlags(mediatorCmd, &options)

	return mediatorCmd
}

func runMediator(cmd *cobra.Command, options mediator.MediatorOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	executor, err := bacalhau.NewBacalhauExecutor(options.Bacalhau)
	if err != nil {
		return err
	}

	mediatorService, err := mediator.NewMediator(options, web3SDK, executor)
	if err != nil {
		return err
	}

	log.Debug().Msgf("Starting mediator service.")
	mediatorErrors := mediatorService.Start(commandCtx.Ctx, commandCtx.Cm)
	for {
		select {
		case err := <-mediatorErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
