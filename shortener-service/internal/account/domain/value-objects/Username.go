package ValueObjects

type Username struct {
	value string
}

func NewUsername(value string) (*Username, error) {
	// TODO: Check if email has the requirement
	return &Username{
		value: value,
	}, nil
}

func (id *Username) GetValue() string {
	return id.value
}
