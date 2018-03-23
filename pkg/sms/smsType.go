package sms

type (
	ISmsService interface {
		SendRaw(text string, phoneNumbers ...string) error
	}

	smsService struct {
		provider IProviderService
	}
)
