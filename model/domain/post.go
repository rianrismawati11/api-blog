package domain

import "api-blog/model/web"

type Post struct {
	Id             int
	FeaturedImage  string
	Title          string
	Description    string
	CategoryId     int
	CategoryDetail web.CategoryDetail
}
