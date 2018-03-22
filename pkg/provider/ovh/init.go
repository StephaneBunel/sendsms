package ovh

import (
	domain "github.com/StephaneBunel/sendsms/pkg/sms"
)

func init() {
	ovh, _ := ovhNewProvider()
	domain.GetProviderRepository().Add(ovh)
}
