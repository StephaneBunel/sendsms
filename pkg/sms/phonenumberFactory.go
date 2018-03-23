package sms

func NewPhonenumber() IPhonenumberService {
	return new(phonenumber)
}
