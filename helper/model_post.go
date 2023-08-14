package helper

import (
	"api-blog/model/domain"
	"api-blog/model/web"
)

func ToPostResponseWithoutCategoryDetail(post domain.Post) web.PostResponseWithoutCategoryDetail {
	return web.PostResponseWithoutCategoryDetail{
		Id: post.Id,
		FeaturedImage: web.ImageDetail{
			Url:   post.FeaturedImage,
			Title: post.TitleFeaturedImage,
			Type:  post.TypeFeaturedImage,
		},
		Title:       post.Title,
		Description: post.Description,
		CategoryId:  post.CategoryId,
	}
}

func ToPostResponse(post domain.Post) web.PostResponse {
	return web.PostResponse{
		Id: post.Id,
		FeaturedImage: web.ImageDetail{
			Url:   post.FeaturedImage,
			Title: post.TitleFeaturedImage,
			Type:  post.TypeFeaturedImage,
		},
		Title:       post.Title,
		Description: post.Description,
		CategoryDetail: web.CategoryDetail{
			Id:   post.CategoryDetail.Id,
			Name: post.CategoryDetail.Name,
		},
	}
}

func ToPostResponses(posts []domain.Post) []web.PostResponse {
	var postResponses []web.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, ToPostResponse(post))
	}
	return postResponses
}
