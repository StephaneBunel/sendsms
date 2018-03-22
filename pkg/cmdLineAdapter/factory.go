package cmdLineAdapter

import (
	"fmt"
	"os"

	"github.com/StephaneBunel/sendsms/pkg/appconfig"

	"github.com/spf13/cobra"
)

type (
	cmdLineAdapter struct {
		config  *appconfig.AppConfig
		rootCmd *cobra.Command
		flags   struct {
			profile string
			message string
			phones  []string
			stdin   bool
		}
	}
)

func New(config *appconfig.AppConfig) *cmdLineAdapter {
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
