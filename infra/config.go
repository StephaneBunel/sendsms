package infra

import (
	"fmt"
	"os"

	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Viper *viper.Viper
	}
)

func NewAppConfig() *AppConfig {
	c := new(AppConfig)
	c.Viper = viper.New()

	c.Viper.SetConfigType("yaml")
	c.Viper.AddConfigPath("/etc/sendsms")
	c.Viper.AddConfigPath("$HOME/.sendsms")
	c.Viper.AddConfigPath(".")
	c.Viper.SetConfigName(c.Viper.GetString("config"))
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
