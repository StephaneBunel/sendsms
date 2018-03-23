package sms

import (
	"github.com/romana/rlog"
)

func (svc *smsService) SendRaw(text string, phoneNumbers ...string) error {
	var phoneNumberSet = make(map[string]IPhonenumberService)
	var finalRecipients = make([]IPhonenumberService, 0)

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
