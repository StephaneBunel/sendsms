package sms

import (
	"github.com/spf13/viper"
)

type (
	IProviderService interface {
		Info() ProviderInfo
		Config(*viper.Viper)
		Send(IMessageService, ...IPhonenumberService) error
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
