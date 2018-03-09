package provider

import (
	"errors"

	modelRecipient "github.com/StephaneBunel/sendsms/recipient"
	modelSMS "github.com/StephaneBunel/sendsms/sms"

	"github.com/spf13/viper"
)

type (
	ProviderInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}

	Provider interface {
		Info() ProviderInfo
		Config(*viper.Viper)
		Send(modelRecipient.RecipientList, *modelSMS.Message) error
	}
)

var (
	availableProviderCatalog = make(map[string]func() (Provider, error))
)

func GetCatalog() map[string]func() (Provider, error) {
	return availableProviderCatalog
}

func NewProvider(name string) (Provider, error) {
	smsprovider, exists := availableProviderCatalog[name]
	if exists {
		return smsprovider()
	}
	return nil, errors.New("Provider not found.")
}
