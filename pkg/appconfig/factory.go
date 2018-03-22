package appconfig

import (
	"fmt"
	"os"

	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

func New() *AppConfig {
	c := new(AppConfig)
	c.Viper = viper.New()

	c.Viper.SetConfigType("yaml")
	c.Viper.AddConfigPath("/etc/" + APP_NAME)
	c.Viper.AddConfigPath("$HOME/." + APP_NAME)
	c.Viper.AddConfigPath(".")
	c.Viper.SetConfigName(c.Viper.GetString(CONF_FILE_NAME))
	err := c.Viper.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("Error reading configuration: %s \n", err)
		os.Exit(1)
	}

	cfKey := "sendsms.logLevel"
	if c.Viper.IsSet(cfKey) {
		os.Setenv("RLOG_LOG_LEVEL", c.Viper.GetString(cfKey))
	}
	rlog.UpdateEnv()

	return c
}
