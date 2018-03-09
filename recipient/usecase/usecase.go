package usecase

import (
	models "github.com/StephaneBunel/sendsms/recipient"
	"github.com/StephaneBunel/sendsms/recipient/repository"
)

type (
	RecipientUsecase interface {
		GetByName(name string) models.RecipientList
	}

	recipientUsecase struct {
		repo repository.RecipientRepository
	}
)

func (ru *recipientUsecase) GetByName(name string) models.RecipientList {
	return ru.repo.Fetch(name)
}

func NewRecipientUsecase(repo repository.RecipientRepository) RecipientUsecase {
	return &recipientUsecase{repo}
}
