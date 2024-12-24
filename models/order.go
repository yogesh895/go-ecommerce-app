package models

// Order struct represents an e-commerce order
type Order struct {
	ID      int    `json:"id"`
	ProductID int   `json:"product_id"`
	Quantity int   `json:"quantity"`
	Status  string `json:"status"`
}

