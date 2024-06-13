package hive

import (
	"encoding/json"
	"fmt"

	"github.com/CoopHive/coophive/pkg/inspect"
	optionsfactory "github.com/CoopHive/coophive/pkg/options"
	"github.com/CoopHive/coophive/pkg/web3"
	"github.com/CoopHive/coophive/pkg/web3/bindings/storage"
	"github.com/spf13/cobra"
)

func newInspectCmd() *cobra.Command {
	options := optionsfactory.NewInspectOptions()
	runCmd := &cobra.Command{
		Use:     "inspect",
		Short:   "View metadata for a previously run job.",
		Long:    "View metadata for a previously run job.",
		Example: "inspect cowsay:v0.0.1 -i Message=moo",
		RunE: func(cmd *cobra.Command, args []string) error {

			options, err := optionsfactory.ProcessInspectOptions(options, args)
			if err != nil {
				return err
			}
			return inspectJob(cmd, options)
		},
	}
	optionsfactory.AddInspectCliFlags(runCmd, &options)
	return runCmd
}

func inspectJob(cmd *cobra.Command, options inspect.InspectOptions) error {
	web3SDK, err := web3.NewContractSDK(options.Web3)
	if err != nil {
		return err
	}

	fmt.Printf("deal id: %s\n", options.DealID)

	dealData, err := web3SDK.Contracts.Storage.GetDeal(web3SDK.CallOpts, options.DealID)
	if err != nil {
		return err
	}

	resultData, err := web3SDK.Contracts.Storage.GetResult(web3SDK.CallOpts, options.DealID)
	if err != nil {
		return err
	}

	agreementData, err := web3SDK.Contracts.Storage.GetAgreement(web3SDK.CallOpts, options.DealID)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(struct {
		Deal      storage.SharedStructsDeal
		Result    storage.SharedStructsResult
		Agreement storage.SharedStructsAgreement
	}{
		Deal:      dealData,
		Result:    resultData,
		Agreement: agreementData,
	})
	if err != nil {
		return err
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	return nil
}
