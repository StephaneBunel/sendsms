package service

import (
	"github.com/StephaneBunel/sendsms/domain"
	"github.com/romana/rlog"
)

type (
	SmsService struct {
		provider domain.ISmsProvider
	}
)

func NewSmsService(provider domain.ISmsProvider) *SmsService {
	s := new(SmsService)
	s.provider = provider
	return s
}

func (svc *SmsService) SendRaw(text string, phoneNumbers ...string) error {
	var phoneNumberSet = make(map[string]*domain.PhoneNumber)
	var finalRecipients = make([]*domain.PhoneNumber, 0)

	for _, phone := range phoneNumbers {
		p := domain.NewPhoneNumber()
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

	message := domain.NewSmsMessage()
	message.SetText(text)
	return svc.provider.Send(message, finalRecipients...)
}
