package sms

func NewSmsService(provider IProviderService) ISmsService {
	s := new(smsService)
	s.provider = provider
	return s
}
