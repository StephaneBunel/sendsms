package cmdLineAdapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/StephaneBunel/sendsms/domain"

	"github.com/romana/rlog"
	"github.com/spf13/cobra"
)

func (cli *cmdLineAdapter) providerCmdInit() {
	cliProviderCmd := &cobra.Command{
		Use:   "provider",
		Short: "Get information about available providers (need subcommand)",
		Args:  cobra.MinimumNArgs(1),
	}

	cliProviderListCmd := &cobra.Command{
		Use:   "list",
		Short: "List available providers",
		Long:  "List and show version of each availabale providers",
		Run:   cli.providerListCmd,
	}

	cliProviderInfoCmd := &cobra.Command{
		Use:   "info <provider>",
		Short: "Show informations about specified provider",
		Args:  cobra.MinimumNArgs(1),
		Run:   cli.providerInfoCmd,
	}

	cliProviderCmd.AddCommand(cliProviderListCmd)
	cliProviderCmd.AddCommand(cliProviderInfoCmd)
	cli.rootCmd.AddCommand(cliProviderCmd)
}

func (cli *cmdLineAdapter) providerListCmd(cmd *cobra.Command, args []string) {
	providerRepository := domain.NewProviderRepository()
	providerList := providerRepository.ListByName()
	fmt.Println(strings.Join(providerList, ", "))
}

func (cli *cmdLineAdapter) providerInfoCmd(cmd *cobra.Command, args []string) {
	providerRepository := domain.NewProviderRepository()
	if len(args) < 1 {
		return
	}
	for _, name := range args {
		p := providerRepository.FindByName(name)
		if p == nil {
			continue
		}
		info, err := json.Marshal(p.Info())
		if err != nil {
			rlog.Error(err)
			os.Exit(1)
		}
		var out bytes.Buffer
		json.Indent(&out, info, "", "  ")
		fmt.Println(out.String())
	}
}
