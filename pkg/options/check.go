package options

import (
	"github.com/CoopHive/coophive/pkg/check"
	_ "github.com/CoopHive/coophive/pkg/system"
	"github.com/spf13/cobra"
)

func NewCheckOptions() check.CheckOptions {
	options := check.CheckOptions{
	}
	//options.Web3.Service = system.JobCreatorService
	return options
}

func ProcessCheckOptions(options check.CheckOptions, args []string) (check.CheckOptions, error) {
	return options, CheckCheckOptions(options)
}

func CheckCheckOptions(options check.CheckOptions) error {
	return nil
}

func AddCheckCliFlags(cmd *cobra.Command, options *check.CheckOptions) {
	//AddWeb3CliFlags(cmd, &options.Web3)
}
