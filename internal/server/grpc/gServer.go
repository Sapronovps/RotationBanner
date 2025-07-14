package grpc

import (
	"context"
	"github.com/Sapronovps/RotationBanner/internal/app"
	internalgrpcprotobuf "github.com/Sapronovps/RotationBanner/internal/server/grpc/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
)

type BannerGrpcServer struct {
	address string
	logger  *zap.Logger
	app     *app.App
	server  *grpc.Server
	internalgrpcprotobuf.UnimplementedBannerServiceServer
}

func NewBannerGrpcServer(address string, logger *zap.Logger, app *app.App) *BannerGrpcServer {
	return &BannerGrpcServer{
		address: address,
		logger:  logger,
		app:     app,
	}
}

func (s *BannerGrpcServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.address)
	_ = lis
	if err != nil {
		s.logger.Fatal("failed to listen", zap.Error(err))
	}

	s.server = grpc.NewServer(grpc.UnaryInterceptor(s.LoggerInterceptor))
	internalgrpcprotobuf.RegisterBannerServiceServer(s.server, s)

	s.logger.Info("gRPC server started", zap.String("address", s.address))
	if err := s.server.Serve(lis); err != nil {
		s.logger.Fatal("failed to serve", zap.Error(err))
	}

	<-ctx.Done()
	return nil
}

func (s *BannerGrpcServer) Stop(_ context.Context) error {
	s.server.GracefulStop()

	os.Exit(1)
	return nil
}
