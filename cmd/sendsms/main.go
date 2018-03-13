package main

import (
	"os"

	recipientModels "github.com/StephaneBunel/sendsms/recipient"
	recipientRepository "github.com/StephaneBunel/sendsms/recipient/repository"
	recipientUsecase "github.com/StephaneBunel/sendsms/recipient/usecase"

	"github.com/romana/rlog"
)

// createInMemoryRecipientRepository creates a new in memory recipient repository
// and it adds some recipients into it.
func createInMemoryRecipientRepository() recipientRepository.RecipientRepository {
	recipientRepo, err := recipientRepository.NewRecipientRepository("InMemory")
	if err != nil {
		rlog.Error(err)
		os.Exit(1)
	}
	recipientRepo.Create(recipientModels.Recipient{
		Name: "Stéphane", PhoneNumber: "+33672967852",
	})
	recipientRepo.Create(recipientModels.Recipient{
		Name: "Stéphane", PhoneNumber: "+invalid",
	})
	recipientRepo.Create(recipientModels.Recipient{
		Name: "astreinte", PhoneNumber: "+33671835916",
	})
	recipientRepo.Create(recipientModels.Recipient{
		Name: "Stéphane", PhoneNumber: "+33672967852",
	})

	return recipientRepo
}

func scratch() {
	recipientRepo := createInMemoryRecipientRepository()
	recipientBiz := recipientUsecase.NewRecipientUsecase(recipientRepo)

	targets := recipientBiz.GetByName("Stéphane")
	if len(targets) == 0 {
		rlog.Error("Recipient not found.")
		os.Exit(1)
	}

	// hello := smsModels.Message{Text: "Hello world !!"}
	// rlog.Debugf("Send message (%v) to recipient (%v) via provider (%v)\n",
	// hello, targets, provider)
}

func main() {
	Execute()
}
