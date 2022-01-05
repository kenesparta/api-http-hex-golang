package mooc

import "github.com/google/uuid"

type GoogleUuid struct {
	value string
}

type IUuid interface {
	Parse(string) ([16]byte, error)
}

func (u *GoogleUuid) Parse(value string) ([16]byte, error) {
	return uuid.Parse(value)
}
