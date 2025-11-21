package models

import "github.com/google/uuid"

type Address struct {
	Id        uuid.UUID
	City      string
	Street    string
	Building  string
	Apartment string
}
