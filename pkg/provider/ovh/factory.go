package ovh

import (
	domain "github.com/StephaneBunel/sendsms/pkg/sms"

	"github.com/ovh/go-ovh/ovh"
)

func ovhNewProvider() (domain.ProviderService, error) {
	p := new(ovhProvider)
	p.Name = "OVH"
	p.Version = "0.2"
	p.Authors = "(c) 2018 Stéphane Bunel"
	p.Help = "Configuration example (in YAML)\n" +
		"  providerConfig:\n" +
		"    api:\n" +
		"      location:          ovh-eu\n" +
		"      appKey:            <appKey>\n" +
		"      appSecret:         <appSecret>\n" +
		"      consumerKey:       <consumerKey>\n" +
		"      serviceName:       <service endpoint>\n" +
		"    smsOptions:\n" +
		"      charset:           UTF-8\n" +
		"      class:             phoneDisplay\n" +
		"      coding:            8bit\n" +
		"      nostopclause:      true\n" +
		"      priority:          high\n" +
		"      senderforresponse: false\n" +
		"      sender:            <sender>\n" +
		"    smsOptionsCaps:\n" +
		"      nostopclause:      noStopClause\n" +
		"      servicename:       serviceName\n" +
		"      senderforresponse: senderForResponse\n"
	return p, nil
}

func (p *ovhProvider) open() (err error) {
	if p.client != nil {
		return nil
	}

	p.client, err = ovh.NewClient(p.apiConfig.location, p.apiConfig.appKey, p.apiConfig.appSecret, p.apiConfig.consumerKey)
	return err
}
