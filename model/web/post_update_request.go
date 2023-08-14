package web

type PostUpdateRequest struct {
	Id                 int
	FeaturedImage      string `validate:"max=200"`
	TitleFeaturedImage string `validate:"max=200"`
	TypeFeaturedImage  string `validate:"required,max=200"`
	Title              string `validate:"required,max=200"`
	Description        string `validate:"required,max=250"`
	CategoryId         int    `validate:"required,max=50"`
}
