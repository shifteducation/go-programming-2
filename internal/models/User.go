package models

import "github.com/google/uuid"

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	age       uint8
}
