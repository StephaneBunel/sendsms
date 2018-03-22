package sms

import (
	"github.com/romana/rlog"
)

type (
	smsService struct {
		provider ProviderService
	}
)

func NewSmsService(provider ProviderService) *smsService {
	s := new(smsService)
	s.provider = provider
	return s
}

func (svc *smsService) SendRaw(text string, phoneNumbers ...string) error {
	var phoneNumberSet = make(map[string]PhonenumberService)
	var finalRecipients = make([]PhonenumberService, 0)

	for _, phone := range phoneNumbers {
		p := NewPhonenumber()
		err := p.Set(phone)
		if err != nil {
			return err
		}
		phoneNumberSet[p.Get()] = p
	}
	rlog.Debug("phoneNumberSet = ", phoneNumberSet)

	for _, phone := range phoneNumberSet {
		finalRecipients = append(finalRecipients, phone)
	}
	rlog.Debug("final recipient = ", finalRecipients)

	message := NewMessage()
	message.SetText(text)
	return svc.provider.Send(message, finalRecipients...)
}
