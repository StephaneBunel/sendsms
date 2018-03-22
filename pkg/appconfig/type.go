package appconfig

import (
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Viper *viper.Viper
	}
)

const (
	APP_NAME       = "senssms"
	CONF_FILE_NAME = "config"
)
