package sms

func NewMessage() IMessageService {
	m := new(smsMessage)
	m.options = make(map[string]interface{})
	return m
}
