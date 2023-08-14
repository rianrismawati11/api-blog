package service

import (
	"api-blog/model/web"
	"context"
)

type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest) web.PostResponseWithoutCategoryDetail
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponseWithoutCategoryDetail
	Delete(ctx context.Context, postId int)
	FindById(ctx context.Context, postId int) web.PostResponse
	FindByCategory(ctx context.Context, categoryId int) []web.PostResponse
	FindAll(ctx context.Context) []web.PostResponse
}
