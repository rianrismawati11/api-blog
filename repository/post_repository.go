package repository

import (
	"api-blog/model/domain"
	"context"
	"database/sql"
)

type PostRepository interface {
	Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, post domain.Post)
	FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	FindByCategory(ctx context.Context, tx *sql.Tx, postId int) []domain.Post
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Post, error)
}
