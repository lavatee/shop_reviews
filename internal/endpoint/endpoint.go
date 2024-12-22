package endpoint

import (
	pb "github.com/lavatee/shop_protos/gen"
	"github.com/lavatee/shop_reviews/internal/service"
)

type Endpoint struct {
	services *service.Service
	pb.UnimplementedReviewsServer
}

func NewEndpoint(services *service.Service) *Endpoint {
	return &Endpoint{
		services: services,
	}
}
