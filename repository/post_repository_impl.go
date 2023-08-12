package repository

import (
	helper "api-blog/helper"
	"api-blog/model/domain"
	"api-blog/model/web"
	"context"
	"database/sql"
	"errors"
)

type PostRepositoryImpl struct {
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

//semua contract dr PostRepositoryImpl struct

func (repository *PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "insert into posts(featured_image, title, description, category_id) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, post.FeaturedImage, post.Title, post.Description, post.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)
	return post
}

func (repository *PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "update posts set featured_image=?, title=?, description=?, category_id=? where id= ?"
	_, err := tx.ExecContext(ctx, SQL, post.FeaturedImage, post.Title, post.Description, post.CategoryId, post.Id)
	helper.PanicIfError(err)

	return post
}

func (repository *PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, post domain.Post) {
	SQL := "delete from posts where id= ?"
	_, err := tx.ExecContext(ctx, SQL, post.Id)
	helper.PanicIfError(err)
}

func (repository *PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	SQL := "SELECT posts.featured_image, posts.title, posts.description, posts.category_id, categories.name, categories.id FROM posts JOIN categories ON posts.category_id = categories.id WHERE posts.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, postId)
	if err != nil {
		return domain.Post{}, err
	}
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {
		categoryDetail := web.CategoryDetail{}

		err := rows.Scan(&post.FeaturedImage, &post.Title, &post.Description, &post.CategoryId, &categoryDetail.Name, &categoryDetail.Id)
		if err != nil {
			return domain.Post{}, err
		}

		post.CategoryDetail = categoryDetail
		return post, nil
	} else {
		return domain.Post{}, errors.New("post is not found")
	}
}

func (repository *PostRepositoryImpl) FindByCategory(ctx context.Context, tx *sql.Tx, categoryId int) []domain.Post {
	SQL := "select id, featured_image, title, description, category_id from posts where category_id=?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		post := domain.Post{}
		err := rows.Scan(&post.Id, &post.FeaturedImage, &post.Title, &post.Description, &post.CategoryId)
		helper.PanicIfError(err)
		posts = append(posts, post)
	}
	return posts
}

// func (repository *PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Post, error) {
// 	SQL := "select id as id_posts, featured_image, title, description, category_id from posts JOIN categories ON posts.category_id = categories.id"
// 	rows, err := tx.QueryContext(ctx, SQL)
// 	helper.PanicIfError(err)
// 	defer rows.Close()

// 	// slice
// 	var posts []domain.Post
// 	for rows.Next() {
// 		post := domain.Post{}
// 		categoryDetail := web.CategoryDetail{}
// 		err := rows.Scan(&post.FeaturedImage, &post.Title, &post.Description, &post.CategoryId, &categoryDetail.Name, &categoryDetail.Id)
// 		helper.PanicIfError(err)
// 		posts = append(posts, post)
// 	}
// 	return posts, nil
// }

func (repository *PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Post, error) {
	SQL := `SELECT posts.id AS id_posts, posts.featured_image, posts.title, posts.description, posts.category_id, categories.name AS category_name, categories.id AS category_id FROM posts JOIN categories ON posts.category_id = categories.id`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		post := domain.Post{}
		categoryDetail := web.CategoryDetail{}
		err := rows.Scan(
			&post.Id,
			&post.FeaturedImage,
			&post.Title,
			&post.Description,
			&post.CategoryId,
			&categoryDetail.Name,
			&categoryDetail.Id,
		)
		helper.PanicIfError(err)
		post.CategoryDetail = categoryDetail
		posts = append(posts, post)
	}
	return posts, nil
}
