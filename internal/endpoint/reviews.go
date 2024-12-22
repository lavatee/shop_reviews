package endpoint

import (
	"context"

	pb "github.com/lavatee/shop_protos/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *Endpoint) PostReview(c context.Context, req *pb.PostReviewRequest) (*pb.PostReviewResponse, error) {
	if req.Text == "" {
		return nil, status.Error(codes.Internal, "request must have text")
	}
	if req.Evaluation == 0 {
		return nil, status.Error(codes.Internal, "request must have evaluation")
	}
	if req.ProductId == 0 {
		return nil, status.Error(codes.Internal, "request must have product's id")
	}
	if req.UserId == 0 {
		return nil, status.Error(codes.Internal, "request must have user's id")
	}
	reviewId, err := e.services.Reviews.PostReview(req.Text, int(req.UserId), int(req.ProductId), int(req.Evaluation))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.PostReviewResponse{
		Id: int64(reviewId),
	}, nil
}

func (e *Endpoint) DeleteReview(c context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "request must have review's id")
	}
	err := e.services.Reviews.DeleteReview(int(req.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DeleteReviewResponse{
		Status: "ok",
	}, nil
}

func (e *Endpoint) GetProductReviews(c context.Context, req *pb.GetProductReviewsRequest) (*pb.GetProductsReviewResponse, error) {
	if req.ProductId == 0 {
		return nil, status.Error(codes.InvalidArgument, "request must have product's id")
	}
	reviews, err := e.services.Reviews.GetProductReviews(int(req.ProductId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pbReviews := make([]*pb.Review, len(reviews))
	for i, review := range reviews {
		pbReviews[i] = &pb.Review{
			Id:         int64(review.Id),
			Text:       review.Text,
			Evaluation: int64(review.Evaluation),
			ProductId:  int64(review.ProductId),
			UserId:     int64(review.UserId),
		}
	}
	return &pb.GetProductsReviewResponse{
		Reviews: pbReviews,
	}, nil
}

func (e *Endpoint) GetAverageEvaluation(c context.Context, req *pb.GetAverageEvaluationRequest) (*pb.GetAverageEvaluationResponse, error) {
	if req.ProductId == 0 {
		return nil, status.Error(codes.InvalidArgument, "request must have product's id")
	}
	averageEvaluation, err := e.services.Reviews.GetAverageEvaluation(int(req.ProductId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetAverageEvaluationResponse{
		AverageEvaluation: float32(averageEvaluation),
	}, nil
}
