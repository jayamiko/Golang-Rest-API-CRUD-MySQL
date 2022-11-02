package models

type Product struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Stock int32   `json:"stock"`
	Price float64 `json:"price"`
}
