package main

import (
	"github.com/StephaneBunel/sendsms/pkg/appconfig"
	"github.com/StephaneBunel/sendsms/pkg/cmdLineAdapter"

	_ "github.com/StephaneBunel/sendsms/pkg/provider/ovh" // Force to init() provider
)

func main() {
	cnf := appconfig.New()
	cli := cmdLineAdapter.New(cnf)
	cli.Execute()
}
