package sms

import "strings"

type (
	providerRepository struct {
		repo map[string]ProviderService
	}
)

var (
	smsProviderRepository = &providerRepository{
		repo: make(map[string]ProviderService),
	}
)

func GetProviderRepository() ProviderRepositoryService {
	return smsProviderRepository
}

func (prepo *providerRepository) Add(provider ProviderService) {
	nameLow := strings.ToLower(provider.Info().Name)
	prepo.repo[nameLow] = provider
}

func (prepo *providerRepository) Exists(name string) bool {
	nameLow := strings.ToLower(name)
	_, exists := prepo.repo[nameLow]
	return exists
}

func (prepo *providerRepository) FindByName(name string) ProviderService {
	if prepo.Exists(name) {
		return prepo.repo[name]
	}
	return nil
}

func (prepo *providerRepository) ListByName() []string {
	providerList := make([]string, 0)

	for providerName, _ := range prepo.repo {
		providerList = append(providerList, providerName)
	}

	return providerList
}
