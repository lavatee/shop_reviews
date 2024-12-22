package service

import (
	"errors"

	reviews "github.com/lavatee/shop_reviews"
	"github.com/lavatee/shop_reviews/internal/repository"
)

type ReviewsService struct {
	repo *repository.Repository
}

func NewReviewsService(repo *repository.Repository) *ReviewsService {
	return &ReviewsService{
		repo: repo,
	}
}

func (s *ReviewsService) PostReview(text string, userId int, productId int, evaluation int) (int, error) {
	producer := PostReviewProducer{
		repo:      s.repo,
		observers: []PostReviewObserver{},
	}
	event := producer.PostReview(text, userId, productId, evaluation)
	if !event.IsOk {
		return 0, errors.New(event.ErrorText)
	}
	return event.ReviewId, nil
}

type PostReviewEvent struct {
	ReviewId         int
	ReviewText       string
	ReviewEvaluation int
	ReviewProductId  int
	ReviewUserId     int
	IsOk             bool
	ErrorText        string
}

type PostReviewProducer struct {
	repo      *repository.Repository
	observers []PostReviewObserver
}

func (p PostReviewProducer) PostReview(text string, userId int, productId int, evaluation int) PostReviewEvent {
	reviewId, err := p.repo.PostReview(text, userId, productId, evaluation)
	if err != nil {
		return PostReviewEvent{IsOk: false, ErrorText: err.Error()}
	}
	event := PostReviewEvent{
		IsOk:             true,
		ReviewId:         reviewId,
		ReviewText:       text,
		ReviewEvaluation: evaluation,
		ReviewProductId:  productId,
		ReviewUserId:     userId,
	}
	for _, observer := range p.observers {
		observer.Update(&event)
		if !event.IsOk {
			return event
		}
	}
	return event
}

type PostReviewObserver interface {
	Update(event *PostReviewEvent)
}

func (s *ReviewsService) DeleteReview(id int) error {
	producer := DeleteReviewProducer{
		Repo:      s.repo,
		Observers: []DeleteReviewObserver{},
	}
	event := producer.DeleteReview(id)
	if !event.IsOk {
		return errors.New(event.ErrorText)
	}
	return nil
}

type DeleteReviewEvent struct {
	ReviewId  int
	IsOk      bool
	ErrorText string
}

type DeleteReviewProducer struct {
	Observers []DeleteReviewObserver
	Repo      *repository.Repository
}

func (p DeleteReviewProducer) DeleteReview(id int) DeleteReviewEvent {
	if err := p.Repo.Reviews.DeleteReview(id); err != nil {
		return DeleteReviewEvent{
			IsOk:      false,
			ErrorText: err.Error(),
		}
	}
	event := DeleteReviewEvent{
		ReviewId: id,
		IsOk:     true,
	}
	for _, observer := range p.Observers {
		observer.Update(&event)
		if !event.IsOk {
			return event
		}
	}
	return event
}

type DeleteReviewObserver interface {
	Update(event *DeleteReviewEvent)
}

func (s *ReviewsService) GetProductReviews(productId int) ([]reviews.Review, error) {
	return s.repo.Reviews.GetProductReviews(productId)
}

func (s *ReviewsService) GetAverageEvaluation(productId int) (float64, error) {
	return s.repo.Reviews.GetAverageEvaluation(productId)
}
