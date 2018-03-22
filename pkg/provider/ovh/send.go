package ovh

import (
	"errors"

	domain "github.com/StephaneBunel/sendsms/pkg/sms"
	"github.com/romana/rlog"
)

func (p *ovhProvider) prepareNewSMS() ovhSMS {
	sms := make(ovhSMS)
	for k, v := range p.smsConfig {
		sms[k] = v
	}
	return sms
}

func (p *ovhProvider) Send(msg *domain.SmsMessage, phoneNumbers ...domain.PhonenumberService) error {
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
	for _, optionName := range msg.ListOptionsByName() {
		value, exists := msg.FindOption(optionName)
		if exists {
			sms[optionName] = value
		} else {
			rlog.Debug("Cannot finf optionName:", optionName, "!?!?")
		}
	}

	// Check if serviceName is not empty
	if len(p.apiConfig.serviceName) == 0 {
		return errors.New("ServiceName is empty")
	}

	sms["message"] = msg.GetText()
	sms["receivers"] = receivers

	rlog.Debugf("sending sms: %#v\n", sms)

	response := make(map[string]interface{})
	err := p.client.Post("/sms/"+p.apiConfig.serviceName+"/jobs", sms, &response)
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
