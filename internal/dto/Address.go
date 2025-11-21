package dto

type Address struct {
	City      string `json:"city" binding:"required"`
	Street    string `json:"street" binding:"required"`
	Building  string `json:"building" binding:"required"`
	Apartment string `json:"apartment" binding:"required"`
}
