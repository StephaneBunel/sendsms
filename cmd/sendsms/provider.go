package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/StephaneBunel/sendsms/provider"

	"github.com/romana/rlog"
	"github.com/spf13/cobra"
)

var (
	providerCmd = &cobra.Command{
		Use:   "provider",
		Short: "Get information about available providers (need subcommand)",
		Args:  cobra.MinimumNArgs(1),
		Run:   providerRun,
	}

	providerListCmd = &cobra.Command{
		Use:   "list",
		Short: "List available providers",
		Long:  "List and show version of each availabale providers",
		Run:   providerListRun,
	}

	providerInfoCmd = &cobra.Command{
		Use:   "info <provider>",
		Short: "Show full informations about specified provider",
		Args:  cobra.MinimumNArgs(1),
		Run:   providerInfoRun,
	}

	full bool
)

func init() {
	providerCmd.AddCommand(providerListCmd)
	providerCmd.AddCommand(providerInfoCmd)

	rootCmd.AddCommand(providerCmd)
}

func providerRun(cmd *cobra.Command, args []string) {
	rlog.Debug("provider Cmd args =", args)

}

func providerListRun(cmd *cobra.Command, args []string) {
	for name, factory := range provider.GetCatalog() {
		p, err := factory()
		if err != nil {
			rlog.Error(err)
			os.Exit(1)
		}
		info := p.Info()
		fmt.Printf("Provider: %s, full name: %s, version: %s, description: %s\n", name, info.Name, info.Version, info.ShortDescription)
		if full {
			fmt.Printf("    Authors: %s\n", info.Authors)
			fmt.Printf("    Site   : %s\n", info.Site)
			fmt.Printf("    Help   : %s\n", info.Help)
		}
	}
}

func providerInfoRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		for _, name := range args {
			if factory, exists := provider.GetCatalog()[name]; exists {
				p, err := factory()
				if err != nil {
					rlog.Error(err)
					os.Exit(1)
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
	}
}
