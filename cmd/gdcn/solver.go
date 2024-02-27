package gdcn

import (
	optionsfactory "github.com/CoopHive/coophive/pkg/options"
	"github.com/CoopHive/coophive/pkg/solver"
	memorystore "github.com/CoopHive/coophive/pkg/solver/store/memory"
	"github.com/CoopHive/coophive/pkg/system"
	"github.com/CoopHive/coophive/pkg/web3"
	"github.com/spf13/cobra"
)

func newSolverCmd() *cobra.Command {
	options := optionsfactory.NewSolverOptions()

	solverCmd := &cobra.Command{
		Use:     "solver",
		Short:   "Start the Generic Decentralized Compute Network solver service.",
		Long:    "Start the Generic Decentralized Compute Network solver service.",
		Example: "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			options, err := optionsfactory.ProcessSolverOptions(options)
			if err != nil {
				return err
			}
			return runSolver(cmd, options)
		},
	}

	optionsfactory.AddSolverCliFlags(solverCmd, &options)

	return solverCmd
}

func runSolver(cmd *cobra.Command, options solver.SolverOptions) error {
	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	solverStore, err := memorystore.NewSolverStoreMemory()
	if err != nil {
		return err
	}

	solverService, err := solver.NewSolver(options, solverStore, web3SDK)
	if err != nil {
		return err
	}

	solverErrors := solverService.Start(commandCtx.Ctx, commandCtx.Cm)

	for {
		select {
		case err := <-solverErrors:
			commandCtx.Cleanup()
			return err
		case <-commandCtx.Ctx.Done():
			return nil
		}
	}
}
