package entity

type Product struct {
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	Gtin     string `json:"gtin"`
	Inclusao string `json:"inclusao"`
}

type Products []Product
