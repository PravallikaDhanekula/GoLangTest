package model

type Product struct {
	Id       int    `json:"id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
