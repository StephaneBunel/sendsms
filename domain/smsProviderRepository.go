package domain

import "strings"

type (
	ISmsProviderRepository interface {
		Add(ISmsProvider)
		Exists(name string) bool
		FindByName(name string) ISmsProvider
		ListByName() []string
	}

	SmsProviderRepository struct {
		repo map[string]ISmsProvider
	}
)

var (
	smsProviderRepository = &SmsProviderRepository{
		repo: make(map[string]ISmsProvider),
	}
)

func NewProviderRepository() ISmsProviderRepository {
	return smsProviderRepository
}

func (prepo *SmsProviderRepository) Add(provider ISmsProvider) {
	nameLow := strings.ToLower(provider.Info().Name)
	prepo.repo[nameLow] = provider
}

func (prepo *SmsProviderRepository) Exists(name string) bool {
	nameLow := strings.ToLower(name)
	_, exists := prepo.repo[nameLow]
	return exists
}

func (prepo *SmsProviderRepository) FindByName(name string) ISmsProvider {
	if prepo.Exists(name) {
		return prepo.repo[name]
	}
	return nil
}

func (prepo *SmsProviderRepository) ListByName() []string {
	providerList := make([]string, 0)

	for providerName, _ := range prepo.repo {
		providerList = append(providerList, providerName)
	}

	return providerList
}
