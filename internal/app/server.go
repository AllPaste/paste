package app

import (
	"net"

	"github.com/AllPaste/paste/config"
	"github.com/AllPaste/paste/internal/service"
	"github.com/AllPaste/paste/pkg/interceptor"
	"github.com/AllPaste/paste/pkg/log"

	cpb "github.com/AllPaste/sdk/core/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
)

type Server struct {
	c      *config.Config
	logger *log.Logger

	gs *grpc.Server
}

func NewServer(
	c *config.Config,
	logger *log.Logger,

	ps *service.PasteService,
) *Server {
	gs := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger.ZapLogger()),
			interceptor.RecoveryStreamUnaryServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger.ZapLogger()),
			interceptor.RecoveryUnaryServerInterceptor,
		)),
	)
	cpb.RegisterPasteServiceServer(gs, ps)

	return &Server{
		c:      c,
		logger: logger,
		gs:     gs,
	}
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.c.Server.Listen)
	if err != nil {
		s.logger.Fatalf("failed to listen: %v", err)
	}

	s.logger.Infof("server listen on %s", s.c.Server.Listen)

	if err := s.gs.Serve(lis); err != nil {
		s.logger.Fatalf("failed to serve: %v", err)
	}
}
