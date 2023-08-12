package web

// model action request
type PostCreateRequest struct {
	FeaturedImage string `validate:"required"`
	Title         string `validate:"required,max=200"`
	Description   string `validate:"required,max=250"`
	CategoryId    int    `validate:"required,max=50"`
}
