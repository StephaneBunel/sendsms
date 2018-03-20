package domain

type (
	SmsMessage struct {
		Text              string
		SmsMessageOptions map[string]interface{}
	}
)

func NewSmsMessage() *SmsMessage {
	m := new(SmsMessage)
	m.SmsMessageOptions = make(map[string]interface{})
	return m
}

func (m *SmsMessage) SetText(test string) {
	m.Text = test
}
