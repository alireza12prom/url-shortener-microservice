package ValueObjects

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	// TODO: Check if email has the requirement
	return &Email{
		value: value,
	}, nil
}

func (id *Email) GetValue() string {
	return id.value
}
