package sms

import "strings"

func (prepo *providerRepository) Add(provider IProviderService) {
	nameLow := strings.ToLower(provider.Info().Name)
	prepo.repo[nameLow] = provider
}

func (prepo *providerRepository) Exists(name string) bool {
	nameLow := strings.ToLower(name)
	_, exists := prepo.repo[nameLow]
	return exists
}

func (prepo *providerRepository) FindByName(name string) IProviderService {
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
