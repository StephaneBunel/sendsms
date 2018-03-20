package domain

import (
	"github.com/spf13/viper"
)

type (
	SmsProviderInfo struct {
		Name             string
		Version          string
		Authors          string
		Site             string
		Help             string
		ShortDescription string
		LongDescription  string
	}

	ISmsProvider interface {
		Info() SmsProviderInfo
		Config(*viper.Viper)
		Send(*SmsMessage, ...*PhoneNumber) error
	}
)
