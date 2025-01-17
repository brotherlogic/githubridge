package server

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	creates = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "githubridge_creates",
		Help: "The number of repos being tracked",
	}, []string{"repo", "code"})

	serverRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "githubridge_requests",
		Help: "The number of server requests",
	}, []string{"method", "status"})
)

func (s *Server) ServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	h, err := handler(ctx, req)
	serverRequests.With(prometheus.Labels{"status": status.Convert(err).Code().String(), "method": info.FullMethod}).Inc()
	return h, err
}
