package ovh

import (
	domain "github.com/StephaneBunel/sendsms/pkg/sms"

	"github.com/ovh/go-ovh/ovh"
)

type (
	ovhProvider struct {
		domain.ProviderInfo
		apiConfig struct {
			location    string
			appKey      string
			appSecret   string
			consumerKey string
			serviceName string // service endpoint
		}
		smsConfig map[string]interface{}
		client    *ovh.Client
	}

	ovhSMS map[string]interface{}
)
