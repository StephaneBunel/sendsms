package sms

type (
	PhonenumberService interface {
		Get() string
		Set(string) error
	}

	phonenumber struct {
		phone string
	}

	PhonenumberList []PhonenumberService
)
