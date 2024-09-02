package model

type Customer struct {
	Name     string `json:"name"`
	Phone    int    `json:"phone"`
	Email    string `json:"email"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
