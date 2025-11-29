package valueobjects

type Username struct {
	value string
}

func NewUsername(value string) (*Username, error) {
	// TODO: Check if email has the requirement
	return &Username{
		value: value,
	}, nil
}

func (u *Username) GetValue() string {
	return u.value
}
