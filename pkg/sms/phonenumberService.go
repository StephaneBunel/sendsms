package sms

import (
	"errors"
	"unicode"
)

var (
	phonenumberErr_BAD = errors.New("BAD Phone number")
)

func (pn *phonenumber) Get() string {
	return pn.phone
}

func (pn *phonenumber) Set(phone string) error {
	// Check if phone number contains only '+' and digital number
	// Remove spaces
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
				return phonenumberErr_BAD
			}
		case unicode.IsDigit(r) == false:
			return phonenumberErr_BAD
		default:
			compose += string(r)
		}
	}

	if compose == "" {
		return phonenumberErr_BAD
	}

	pn.phone = compose
	return nil
}
