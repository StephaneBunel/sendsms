package sms

func (m *smsMessage) SetText(text string) error {
	m.text = text
	return nil
}

func (m *smsMessage) GetText() string {
	return m.text
}

func (m *smsMessage) SetOption(name string, value interface{}) {
	m.options[name] = value
}

func (m *smsMessage) FindOption(name string) (interface{}, bool) {
	v, exists := m.options[name]
	return v, exists
}

func (m *smsMessage) ListOptionByName() []string {
	options := make([]string, 0)
	for name, _ := range m.options {
		options = append(options, name)
	}
	return options
}
