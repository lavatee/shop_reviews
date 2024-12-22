package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	reviews "github.com/lavatee/shop_reviews"
)

type ReviewsPostgres struct {
	db *sqlx.DB
}

func NewReviewsPostgres(db *sqlx.DB) *ReviewsPostgres {
	return &ReviewsPostgres{
		db: db,
	}
}

func (r *ReviewsPostgres) PostReview(text string, userId int, productId int, evaluation int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (text, user_id, product_id, evaluation) VALUES ($1, $2, $3, $4) RETURNING id", reviewsTable)
	row := r.db.QueryRow(query, text, userId, productId, evaluation)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ReviewsPostgres) DeleteReview(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", reviewsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ReviewsPostgres) GetProductReviews(productId int) ([]reviews.Review, error) {
	var productReviews []reviews.Review
	query := fmt.Sprintf("SELECT * FROM %s WHERE product_id = $1", reviewsTable)
	if err := r.db.Select(&productReviews, query, productId); err != nil {
		return nil, err
	}
	return productReviews, nil
}

type AverageEvaluationInput struct {
	EvaluationAmount int `db:"count"`
	SumOfEvaluations int `db:"sum"`
}

func (r *ReviewsPostgres) GetAverageEvaluation(productId int) (float64, error) {
	query := fmt.Sprintf("SELECT count(evaluation), sum(evaluation) FROM %s WHERE product_id = $1", reviewsTable)
	var input AverageEvaluationInput
	if err := r.db.Get(&input, query, productId); err != nil {
		return 0, err
	}
	return float64(input.SumOfEvaluations) / float64(input.EvaluationAmount), nil
}
