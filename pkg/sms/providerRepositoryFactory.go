package sms

var (
	smsProviderRepository = &providerRepository{
		repo: make(map[string]IProviderService),
	}
)

func GetProviderRepository() IProviderRepositoryService {
	return smsProviderRepository
}
