package entities

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

const addressTable = "addresses"

type Address struct {
	Id        uuid.UUID
	City      string
	Street    string
	Building  string
	Apartment string
	UserId    *uuid.UUID `gorm:"column:user_id"`
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func (Address) TableName() string {
	return addressTable
}
