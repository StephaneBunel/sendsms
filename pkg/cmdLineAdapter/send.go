package cmdLineAdapter

import (
	"bufio"
	"io/ioutil"
	"os"

	domain "github.com/StephaneBunel/sendsms/pkg/sms"

	"github.com/romana/rlog"
	"github.com/spf13/cobra"
)

func (cli *cmdLineAdapter) sendCmdInit() {
	cliSendCmd := &cobra.Command{
		Use:   "send",
		Short: "Send SMS according to defined profile.",
		Run:   cli.sendCmd,
	}
	cliSendCmd.Flags().StringVar(&cli.flags.profile, "profile", "default", "Profile to use")
	cliSendCmd.Flags().StringVar(&cli.flags.message, "message", "Hello world !", "Message to send")
	cliSendCmd.Flags().StringArrayVar(&cli.flags.phones, "phone", []string{}, "Phone number")
	cliSendCmd.Flags().BoolVar(&cli.flags.stdin, "stdin", false, "Read message from stdin")
	cli.rootCmd.AddCommand(cliSendCmd)
}

func (cli *cmdLineAdapter) sendCmd(cmd *cobra.Command, args []string) {
	var smsService = domain.NewSmsService(cli.selectAndConfigureProviderFromProfile())
	var texto string

	if cli.flags.stdin {
		// Read message from stdin
		reader := bufio.NewReader(os.Stdin)
		text, err := ioutil.ReadAll(reader)
		if err != nil {
			rlog.Error(err)
			os.Exit(1)
		}
		texto = string(text)
	} else {
		texto = cli.flags.message
	}

	err := smsService.SendRaw(texto, cli.flags.phones...)
	if err != nil {
		rlog.Error(err)
		os.Exit(1)
	}
}

func (cli *cmdLineAdapter) selectAndConfigureProviderFromProfile() domain.IProviderService {
	if cli.flags.profile == "" {
		rlog.Error("You must select the profile to use")
		os.Exit(1)
	}
	profileConfig := cli.config.Viper.Sub("sendsms.profiles." + cli.flags.profile)
	if profileConfig == nil {
		rlog.Error("Cannot find profile:", cli.flags.profile)
		os.Exit(1)
	}
	profileProviderName := profileConfig.GetString("provider")
	if profileProviderName == "" {
		rlog.Error("profile:", cli.flags.profile, "A provider must be defined.")
		os.Exit(1)
	}

	provider := domain.GetProviderRepository().FindByName(profileProviderName)
	if provider == nil {
		rlog.Error("Provider: ", profileProviderName, "not found.")
		os.Exit(1)
	}
	provider.Config(profileConfig.Sub("providerConfig"))
	return provider
}
