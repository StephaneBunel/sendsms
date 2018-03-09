package repository

import (
	"container/list"

	model "github.com/StephaneBunel/sendsms/recipient"

	"github.com/romana/rlog"
)

type (
	inMemoryRecipientRepository struct {
		linkedRecipient *list.List
	}
)

func init() {
	catalogRecipientRepository["InMemory"] = NewInMemoryRecipientRepository
}

func NewInMemoryRecipientRepository() RecipientRepository {
	repo := new(inMemoryRecipientRepository)
	repo.linkedRecipient = list.New()
	return repo
}

func (repo *inMemoryRecipientRepository) Create(r model.Recipient) error {
	repo.linkedRecipient.PushBack(r)
	rlog.Debugf("Add recipient: %v\n", r)
	return nil
}

func (repo *inMemoryRecipientRepository) Fetch(name string) model.RecipientList {
	results := make(model.RecipientList, 0)

	for e := repo.linkedRecipient.Front(); e != nil; e = e.Next() {
		r := e.Value.(model.Recipient)
		if r.Name == name {
			rlog.Debugf("r = %v", r)
			results = append(results, e.Value.(model.Recipient))
		}
	}

	return results
}
