package entities

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

const usersTable = "users"

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Age       uint8
	Address   *Address // `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func (User) TableName() string {
	return usersTable
}
