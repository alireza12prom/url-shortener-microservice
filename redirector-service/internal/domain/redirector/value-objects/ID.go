package ValueObjects

import "github.com/google/uuid"

type ID struct {
	value uuid.UUID
}

func NewID(value string) (*ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return nil, err
	}

	return &ID{value: id}, nil
}

func (s *ID) GetValue() string {
	return s.value.String()
}
