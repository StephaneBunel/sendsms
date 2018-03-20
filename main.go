package main

import (
	"github.com/StephaneBunel/sendsms/infra"
	"github.com/StephaneBunel/sendsms/infra/cmdLineAdapter"
)

func main() {
	Config := infra.NewAppConfig()
	cli := cmdLineAdapter.NewCmdLineAdapter(Config)
	cli.Execute()
}
