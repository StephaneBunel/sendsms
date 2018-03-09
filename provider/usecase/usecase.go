package usecase

import (
	"unicode"

	"github.com/StephaneBunel/sendsms/provider"
	"github.com/StephaneBunel/sendsms/recipient"
	"github.com/StephaneBunel/sendsms/sms"

	"github.com/romana/rlog"
)

type (
	ProviderUsecase interface {
		SendSMS(recipient.RecipientList, *sms.Message) error
	}

	providerUsecase struct {
		provider provider.Provider
	}
)

func NewProviderUsecase(provider provider.Provider) ProviderUsecase {
	return &providerUsecase{provider}
}

func (p *providerUsecase) SendSMS(recipients recipient.RecipientList, message *sms.Message) error {
	var phoneNumberSet = make(map[string]struct{})
	var finalRecipients = make(recipient.RecipientList, 0)

	for _, reci := range recipients {
		var pn string
		for _, c := range reci.PhoneNumber {
			if unicode.IsDigit(c) {
				pn += string(c)
			}
		}
		if pn != "" {
			if _, exists := phoneNumberSet[pn]; !exists {
				phoneNumberSet[pn] = struct{}{}
				finalRecipients = append(finalRecipients, reci)
			} else {
				rlog.Infof("Skipping duplicate phone number %s (%s).\n", reci.PhoneNumber, reci.Name)
			}
		} else {
			rlog.Warnf("Skipping invalid phone number: %s (%s).\n", reci.PhoneNumber, reci.Name)
		}
	}
	rlog.Debug("phoneNumberSet = ", phoneNumberSet)
	rlog.Debug("final recipient = ", finalRecipients)
	return p.provider.Send(finalRecipients, message)
}
