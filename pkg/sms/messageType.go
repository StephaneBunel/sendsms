package sms

type (
	IMessageService interface {
		SetText(string) error
		GetText() string
		SetOption(name string, value interface{})
		FindOption(name string) (value interface{}, exists bool)
		ListOptionByName() []string
	}

	smsMessage struct {
		text    string
		options map[string]interface{}
	}
)
