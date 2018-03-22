package sms

type (
	ProviderRepositoryService interface {
		Add(ProviderService)
		Exists(name string) bool
		FindByName(name string) ProviderService
		ListByName() []string
	}
)
