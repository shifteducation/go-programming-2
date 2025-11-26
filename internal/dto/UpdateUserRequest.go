package dto

type UpdateUserRequest struct {
	FirstName string
	LastName  string
	Age       uint8
	Address   *UpdateAddressRequest
}
