package gdcn

import (
	"github.com/CoopHive/coophive/pkg/executor/bacalhau"
	optionsfactory "github.com/CoopHive/coophive/pkg/options"
	"github.com/CoopHive/coophive/pkg/resourceprovider"
	"github.com/CoopHive/coophive/pkg/system"
	"github.com/CoopHive/coophive/pkg/web3"
	"github.com/spf13/cobra"
)

func newResourceProviderCmd() *cobra.Command {
	options := optionsfactory.NewResourceProviderOptions()

	resourceProviderCmd := &cobra.Command{
		Use:     "resource-provider",
		Short:   "Start the gdcn resource-provider service.",
		Long:    "Start the gdcn resource-provider service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := optionsfactory.ProcessResourceProviderOptions(options)
			if err != nil {
				return err
			}
			return runResourceProvider(cmd, options)
		},
	}

	optionsfactory.AddResourceProviderCliFlags(resourceProviderCmd, &options)

	return resourceProviderCmd
}

func runResourceProvider(cmd *cobra.Command, options resourceprovider.ResourceProviderOptions) error {
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

	resourceProviderService, err := resourceprovider.NewResourceProvider(options, web3SDK, executor)
	if err != nil {
		return err
	}

	resourecProviderErrors := resourceProviderService.Start(commandCtx.Ctx, commandCtx.Cm)
	for {
		select {
		case err := <-resourecProviderErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
