package sms

type (
	IProviderRepositoryService interface {
		Add(IProviderService)
		Exists(name string) bool
		FindByName(name string) IProviderService
		ListByName() []string
	}

	providerRepository struct {
		repo map[string]IProviderService
	}
)
