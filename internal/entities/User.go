package entities

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Age       uint8
	Address   *Address
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
