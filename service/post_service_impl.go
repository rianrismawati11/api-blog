package service

import (
	"api-blog/exception"
	"api-blog/helper"
	"api-blog/model/domain"
	"context"
	"database/sql"

	"api-blog/model/web"
	"api-blog/repository"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostService(postRepository repository.PostRepository, DB *sql.DB, validate *validator.Validate) *PostServiceImpl {
	return &PostServiceImpl{
		PostRepository: postRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest) web.PostResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	post := domain.Post{
		FeaturedImage: request.FeaturedImage,
		Title:         request.Title,
		Description:   request.Description,
		CategoryId:    request.CategoryId,
	}

	post = service.PostRepository.Save(ctx, tx, post)

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, request.Id)
	// helper.PanicIfError(err)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	post.FeaturedImage = request.FeaturedImage
	post.Title = request.Title
	post.Description = request.Description
	post.CategoryId = request.CategoryId

	post = service.PostRepository.Update(ctx, tx, post)

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PostRepository.Delete(ctx, tx, post)
}

func (service *PostServiceImpl) FindById(ctx context.Context, postId int) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPostResponse(post)
}

func (service *PostServiceImpl) FindByCategory(ctx context.Context, categoryId int) []web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	post := service.PostRepository.FindByCategory(ctx, tx, categoryId)

	return helper.ToPostResponses(post)
}

func (service *PostServiceImpl) FindAll(ctx context.Context) []web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitorRollback(tx)

	posts, err := service.PostRepository.FindAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToPostResponses(posts) // Return the post responses and nil error
}
