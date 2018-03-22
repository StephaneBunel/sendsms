package sms

func NewMessage() *SmsMessage {
	m := new(SmsMessage)
	m.options = make(map[string]interface{})
	return m
}

func (m *SmsMessage) SetText(text string) error {
	m.text = text
	return nil
}

func (m *SmsMessage) GetText() string {
	return m.text
}

func (m *SmsMessage) SetOption(name string, value interface{}) error {
	m.options[name] = value
	return nil
}

func (m *SmsMessage) FindOption(name string) (interface{}, bool) {
	v, exists := m.options[name]
	return v, exists
}

func (m *SmsMessage) ListOptionsByName() []string {
	options := make([]string, 0)
	for name, _ := range m.options {
		options = append(options, name)
	}
	return options
}
