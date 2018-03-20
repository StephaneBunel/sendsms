package cmdLineAdapter

import (
	"fmt"
	"os"

	"github.com/StephaneBunel/sendsms/infra"

	"github.com/spf13/cobra"
)

type (
	cmdLineAdapter struct {
		config  *infra.AppConfig
		rootCmd *cobra.Command
		flags   struct {
			profile string
			message string
			phones  []string
			stdin   bool
		}
	}
)

var (

/*
	cliProviderCmd = &cobra.Command{
		Use:   "provider",
		Short: "Get information about available providers (need subcommand)",
		Args:  cobra.MinimumNArgs(1),
		Run:   providerRun,
	}

	cliProviderListCmd = &cobra.Command{
		Use:   "list",
		Short: "List available providers",
		Long:  "List and show version of each availabale providers",
		Run:   providerListRun,
	}

	cliProviderInfoCmd = &cobra.Command{
		Use:   "info <provider>",
		Short: "Show full informations about specified provider",
		Args:  cobra.MinimumNArgs(1),
		Run:   providerInfoRun,
	}
*/

)

func NewCmdLineAdapter(config *infra.AppConfig) *cmdLineAdapter {
	cli := new(cmdLineAdapter)
	cli.config = config
	cli.rootCmd = &cobra.Command{
		Use:   "sendsms",
		Short: "sendsms is a simple tool to send SMS",
	}

	cli.sendCmdInit()
	cli.providerCmdInit()
	return cli
}

func (cli *cmdLineAdapter) Execute() {
	if err := cli.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
