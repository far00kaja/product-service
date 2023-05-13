package entity

type ProductImage struct {
	Id        string `json:"id"`
	ProductId string `json:"productId"`
	ImageName string `json:"imageName"`
}
