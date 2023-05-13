package entity

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"min=2,max=100,required"`
	Description string `json:"description" binding:"max=200"`
	Thumbnail   string `json:"thumbnail"`
}
