package id

import "github.com/google/uuid"

type Id struct {
	value string
}

func New() Id {
	return Id{uuid.New().String()}
}

func FromString(value string) (Id, error) {
	return Id{value}, uuid.Validate(value)
}

func (id Id) String() string {
	return id.value
}

func (id Id) Equal(other Id) bool {
	return id.value == other.value
}
