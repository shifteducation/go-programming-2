package dto

type CreateUserRequest struct {
	FirstName string   `json:"firstName" binding:"required"`
	LastName  string   `json:"lastName"  binding:"required"`
	Age       uint8    `json:"age" binding:"required"`
	Address   *Address `json:"address"`
}
