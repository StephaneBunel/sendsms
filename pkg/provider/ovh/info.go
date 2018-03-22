package ovh

import domain "github.com/StephaneBunel/sendsms/pkg/sms"

func (p *ovhProvider) Info() domain.ProviderInfo {
	return p.ProviderInfo
}
