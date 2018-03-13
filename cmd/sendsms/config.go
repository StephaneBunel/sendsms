package main

import (
	"fmt"
	"os"

	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
)

func init() {
	Config = viper.New()
	Config.SetConfigType("yaml")
	Config.AddConfigPath("/etc/sendsms")
	Config.AddConfigPath("$HOME/.sendsms")
	Config.AddConfigPath(".")
	Config.SetConfigName(Config.GetString("config"))
	err := Config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		_ = fmt.Errorf("Error reading configuration: %s \n", err)
		os.Exit(1)
	}

	cfKey := "sendsms.logLevel"
	if Config.IsSet(cfKey) {
		os.Setenv("RLOG_LOG_LEVEL", Config.GetString(cfKey))
	}
	rlog.UpdateEnv()
}
