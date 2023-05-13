package entity

type ProductDetail struct {
	Id        string  `json:"id"`
	ProductId string  `json:"productId"`
	Color     string  `json:"color"`
	Price     float32 `json:"price"`
}
