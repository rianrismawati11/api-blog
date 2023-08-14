package web

type PostResponse struct {
	Id             int
	FeaturedImage  ImageDetail
	Title          string
	Description    string
	CategoryDetail CategoryDetail
}

type PostResponseWithoutCategoryDetail struct {
	Id            int
	FeaturedImage ImageDetail
	Title         string
	Description   string
	CategoryId    int
}

type CategoryDetail struct {
	Id   int
	Name string
}

type ImageDetail struct {
	Url   string
	Title string
	Type  string
}
