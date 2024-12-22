package repository

import (
	"github.com/jmoiron/sqlx"
	reviews "github.com/lavatee/shop_reviews"
)

type Repository struct {
	Reviews
}

type Reviews interface {
	PostReview(text string, userId int, productId int, evaluation int) (int, error)
	DeleteReview(id int) error
	GetProductReviews(productId int) ([]reviews.Review, error)
	GetAverageEvaluation(productId int) (float64, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reviews: NewReviewsPostgres(db),
	}
}
