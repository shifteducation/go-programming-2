package entities

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Address struct {
	Id        uuid.UUID
	City      string
	Street    string
	Building  string
	Apartment string
	UserId    *uuid.UUID
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
