package reviews

import (
	"net"

	pb "github.com/lavatee/shop_protos/gen"
	"google.golang.org/grpc"
)

type Server struct {
	GRPCServer *grpc.Server
}

func (s *Server) Run(port string, handler pb.ReviewsServer) error {
	pb.RegisterReviewsServer(s.GRPCServer, handler)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.GRPCServer.Serve(listener)
}

func (s *Server) Shutdown() {
	s.GRPCServer.GracefulStop()
}
