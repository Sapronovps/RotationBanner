package grpc

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func (s *BannerGrpcServer) LoggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	s.logger.Info("started",
		zap.String("method", info.FullMethod),
		zap.String("request", fmt.Sprintf("%v", req)),
	)

	start := time.Now()
	resp, err = handler(ctx, req)

	s.logger.Info("completed",
		zap.String("method", info.FullMethod),
		zap.String("request", fmt.Sprintf("%v", req)),
		zap.Duration("duration", time.Since(start)),
		zap.String("error", fmt.Sprintf("%v", err)),
	)

	return resp, err
}
