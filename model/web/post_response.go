package web

type PostResponse struct {
	Id             int
	FeaturedImage  string
	Title          string
	Description    string
	CategoryDetail CategoryDetail
}

type CategoryDetail struct {
	Id   int
	Name string
}

type ImageDetail struct {
	Title string
	Type  string
}
