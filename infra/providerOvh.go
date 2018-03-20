package infra

import (
	"errors"
	"strings"

	"github.com/StephaneBunel/sendsms/domain"

	"github.com/ovh/go-ovh/ovh"
	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

type (
	ovhProvider struct {
		domain.SmsProviderInfo
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

	ovhNewSMS map[string]interface{}
)

func init() {
	ovh, _ := newOvhProvider()
	domain.NewProviderRepository().Add(ovh)
}

func newOvhProvider() (domain.ISmsProvider, error) {
	p := new(ovhProvider)
	p.Name = "OVH"
	p.Version = "0.2"
	p.Authors = "(c) 2018 Stéphane Bunel"
	p.Help = "Configuration example (in YAML)\n" +
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

func (p *ovhProvider) prepareNewSMS() ovhNewSMS {
	sms := make(ovhNewSMS)
	for k, v := range p.smsConfig {
		sms[k] = v
	}
	return sms
}

func (p *ovhProvider) open() (err error) {
	if p.client != nil {
		return nil
	}

	p.client, err = ovh.NewClient(p.apiConfig.location, p.apiConfig.appKey, p.apiConfig.appSecret, p.apiConfig.consumerKey)
	return err
}

func (p *ovhProvider) Info() domain.SmsProviderInfo {
	return p.SmsProviderInfo
}

func (p *ovhProvider) Send(msg *domain.SmsMessage, phoneNumbers ...*domain.PhoneNumber) error {
	if p.client == nil {
		err := p.open()
		if err != nil {
			return err
		}
	}

	// Make receivers list
	receivers := make([]string, 0)
	for _, phone := range phoneNumbers {
		receivers = append(receivers, phone.Get())
	}

	sms := p.prepareNewSMS()
	for opt, val := range msg.SmsMessageOptions {
		sms[opt] = val
	}

	// Check if serviceName is not empty
	if len(p.apiConfig.serviceName) == 0 {
		return errors.New("ServiceName is empty")
	}

	sms["message"] = msg.Text
	sms["receivers"] = receivers

	rlog.Debugf("sending sms: %#v\n", sms)

	response := make(map[string]interface{})
	err := p.client.Post("/sms/"+p.apiConfig.serviceName+"/jobs", sms, &response)
	// var err error
	if err != nil {
		return err
	}
	rlog.Debugf("response: %#v\n", response)
	if ir, exists := response["invalidReceivers"]; exists == true {
		if len(ir.([]interface{})) > 0 {
			rlog.Warnf("Invalid receivers: %v", ir)
		}
	}
	return nil
}

func (p *ovhProvider) Config(userConfig *viper.Viper) {
	p.apiConfig.location = userConfig.GetString("api.location")
	p.apiConfig.appKey = userConfig.GetString("api.appKey")
	p.apiConfig.appSecret = userConfig.GetString("api.appSecret")
	p.apiConfig.consumerKey = userConfig.GetString("api.consumerKey")
	p.apiConfig.serviceName = userConfig.GetString("api.serviceName")

	p.smsConfig = map[string]interface{}{
		"charset":           "UTF-8",
		"class":             "phoneDisplay",
		"coding":            "8bit",
		"nostopclause":      true,
		"priority":          "high",
		"senderforresponse": false,
	}

	// Set defaults
	for opt, val := range userConfig.GetStringMap("smsOptions") {
		rlog.Debug("Set user defaut sms configuration: ", opt, " = ", val)
		p.smsConfig[opt] = val
	}

	// Fix options case
	for key, casedopt := range userConfig.GetStringMap("smsOptionsCaps") {
		lcaseopt := strings.ToLower(key)
		if content, exists := p.smsConfig[lcaseopt]; exists {
			p.smsConfig[casedopt.(string)] = content
			delete(p.smsConfig, lcaseopt)
		}
	}
}
