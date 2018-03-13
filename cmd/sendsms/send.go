package main

import (
	"bufio"
	"io/ioutil"
	"os"

	providerModels "github.com/StephaneBunel/sendsms/provider"
	providerBiz "github.com/StephaneBunel/sendsms/provider/usecase"
	recipientModels "github.com/StephaneBunel/sendsms/recipient"
	smsModels "github.com/StephaneBunel/sendsms/sms"

	"github.com/romana/rlog"
	"github.com/spf13/cobra"
)

var (
	sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send SMS with selected profile",
		Run:   sendRun,
	}

	flagProfile   string
	flagMessage   string
	flagPhone     []string
	flagRecipient []string
	flagStdin     bool
)

func init() {
	sendCmd.Flags().StringVar(&flagProfile, "profile", "default", "Profile to use")
	sendCmd.Flags().StringVar(&flagMessage, "message", "Hello world !", "Message to send")
	sendCmd.Flags().StringArrayVar(&flagPhone, "phone", []string{}, "Phone number")
	// sendCmd.Flags().StringArrayVar(&flagRecipient, "recipient", "", "Recipient name")
	sendCmd.Flags().BoolVar(&flagStdin, "stdin", false, "Read message from stdin")

	rootCmd.AddCommand(sendCmd)
}

func selectAndConfigureProviderFromProfile() providerModels.Provider {
	if flagProfile == "" {
		rlog.Error("You must select the profile to use")
		os.Exit(1)
	}
	profileConfig := Config.Sub("sendsms.profiles." + flagProfile)
	if profileConfig == nil {
		rlog.Error("Cannot find profile:", flagProfile)
		os.Exit(1)
	}
	profileProvider := profileConfig.GetString("provider")
	if profileProvider == "" {
		rlog.Error("profile:", flagProfile, "Provider must be defined.")
		os.Exit(1)
	}
	useProvider, err := providerModels.NewProvider(profileProvider)
	if err != nil {
		rlog.Error(err)
		os.Exit(1)
	}
	useProvider.Config(profileConfig.Sub("providerConfig"))
	return useProvider
}

func sendRun(cmd *cobra.Command, args []string) {
	var message smsModels.Message

	if flagStdin {
		// Read message from stdin
		reader := bufio.NewReader(os.Stdin)
		text, err := ioutil.ReadAll(reader)
		if err != nil {
			rlog.Error(err)
			os.Exit(1)
		}
		message.Text = string(text)
	} else {
		message.Text = flagMessage
	}

	targets := make(recipientModels.RecipientList, 0)
	for _, phoneNumber := range flagPhone {
		targets = append(targets, recipientModels.Recipient{
			PhoneNumber: phoneNumber,
		})
	}

	if len(targets) == 0 {
		rlog.Error("Cannot send any SMS without at least one phone number (see --phone).")
		os.Exit(1)
	}

	provider := selectAndConfigureProviderFromProfile()
	providerBiz := providerBiz.NewProviderUsecase(provider)
	err := providerBiz.SendSMS(targets, &message)
	if err != nil {
		rlog.Error(err)
		os.Exit(1)
	}

}
