package repository

import (
	"errors"
)

import (
	model "github.com/StephaneBunel/sendsms/recipient"
)

type (
	RecipientRepository interface {
		Create(model.Recipient) error
		Fetch(name string) model.RecipientList
	}
)

var (
	catalogRecipientRepository = make(map[string]func() RecipientRepository)
	errNoFound                 = errors.New("respository not found.")
)

func NewRecipientRepository(name string) (RecipientRepository, error) {
	repo, exists := catalogRecipientRepository[name]
	if exists {
		return repo(), nil
	}
	return nil, errNoFound
}
