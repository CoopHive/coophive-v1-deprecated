package options

import (
	"github.com/CoopHive/coophive/pkg/inspect"
	"github.com/CoopHive/coophive/pkg/system"
	"github.com/spf13/cobra"
)

func NewInspectOptions() inspect.InspectOptions {
	options := inspect.InspectOptions{
		Web3: GetDefaultWeb3Options(),
	}
	options.Web3.Service = system.JobCreatorService
	return options
}

func ProcessInspectOptions(options inspect.InspectOptions, args []string) (inspect.InspectOptions, error) {
	dealId := ""
	if len(args) == 1 {
		dealId = args[0]
	}

	if dealId != "" {
		options.DealID = dealId
	}
	newWeb3Options, err := ProcessWeb3Options(options.Web3)
	if err != nil {
		return options, err
	}
	options.Web3 = newWeb3Options

	return options, CheckInspectOptions(options)
}

func CheckInspectOptions(options inspect.InspectOptions) error {
	err := CheckWeb3Options(options.Web3)
	if err != nil {
		return err
	}

	// TODO validate dealID
	return nil
}

func AddInspectCliFlags(cmd *cobra.Command, options *inspect.InspectOptions) {
	AddWeb3CliFlags(cmd, &options.Web3)
}
