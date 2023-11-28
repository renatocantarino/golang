package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func Generate() ID {
	return ID(uuid.New())
}

func Validate(str string) (ID, error) {
	id, err := uuid.Parse(str)
	return ID(id), err
}
