package sms

type (
	IPhonenumberService interface {
		Get() string
		Set(string) error
	}

	phonenumber struct {
		phone string
	}

	PhonenumberList []IPhonenumberService
)
