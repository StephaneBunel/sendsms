package sms

import (
	"github.com/spf13/viper"
)

type (
	ProviderService interface {
		Info() ProviderInfo
		Config(*viper.Viper)
		Send(*SmsMessage, ...PhonenumberService) error
	}

	ProviderInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}
)
