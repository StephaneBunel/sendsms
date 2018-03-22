package sms

type (
	MessageService interface {
		SetText(string) error
		GetText() string
		SetOption(name string, value interface{})
		FindOption(name string) (value interface{}, exists bool)
	}

	SmsMessage struct {
		text    string
		options map[string]interface{}
	}
)
