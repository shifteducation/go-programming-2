package models

import "github.com/google/uuid"

//todo add CreatedAt, UpdatedAt

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Age       uint8
	Address   *Address
}
