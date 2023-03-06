package database

import "github.com/google/uuid"

func GenerateUUID() uint32{
	id, _ := uuid.NewRandom()
	return id.ID()
}