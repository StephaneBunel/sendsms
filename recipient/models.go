package recipient

type (
	Recipient struct {
		Name        string
		PhoneNumber string
		comment     string
		tags        []string
	}

	RecipientList []Recipient

	Group struct {
		Name       string
		Recipients RecipientList
		tags       []string
	}
)
