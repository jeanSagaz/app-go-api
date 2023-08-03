package entity

import uuid "github.com/google/uuid"

//type ID = uuid.UUID

type ID = string

func NewId() ID {
	return ID(uuid.New().String())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id.String()), err
}
