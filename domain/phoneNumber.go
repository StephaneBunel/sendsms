package domain

import (
	"errors"
	"unicode"

	"github.com/romana/rlog"
)

type (
	PhoneNumber struct {
		phone string
	}

	PhoneNumberList []*PhoneNumber
)

var (
	phoneNumberErr_BAD = errors.New("BAD Phone number")
)

func NewPhoneNumber() *PhoneNumber {
	return new(PhoneNumber)
}

func (pn *PhoneNumber) Get() string {
	return pn.phone
}

func (pn *PhoneNumber) Set(phone string) error {
	// Check if phone number contains only '+' and digital number
	var plus bool = false
	var compose string

	for _, r := range phone {
		switch {
		case unicode.IsSpace(r) == true:
			break
		case r == '+':
			if plus == false {
				compose += "+"
				plus = true
			} else {
				rlog.Debug(compose)
				return phoneNumberErr_BAD
			}
		case unicode.IsDigit(r) == false:
			rlog.Debug(compose)
			return phoneNumberErr_BAD
		default:
			compose += string(r)
		}
	}

	if compose == "" {
		rlog.Debug(compose)
		return phoneNumberErr_BAD
	}

	pn.phone = compose
	return nil
}
