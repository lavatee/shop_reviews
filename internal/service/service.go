package service

import (
	reviews "github.com/lavatee/shop_reviews"
	"github.com/lavatee/shop_reviews/internal/repository"
)

type Service struct {
	Reviews
}

type Reviews interface {
	PostReview(text string, userId int, productId int, evaluation int) (int, error)
	DeleteReview(id int) error
	GetProductReviews(productId int) ([]reviews.Review, error)
	GetAverageEvaluation(productId int) (float64, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Reviews: NewReviewsService(repo),
	}
}
