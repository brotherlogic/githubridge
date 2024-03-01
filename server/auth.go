package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthCall(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	key, err := getContextKey(ctx)
	if err != nil {
		if key != s.authKey {
			return nil, status.Errorf(codes.Unauthenticated, "You need to specify the auth-token (you sent %v)", key)
		}
	}

	return handler(ctx, req)
}

func getContextKey(ctx context.Context) (string, error) {
	md, found := metadata.FromIncomingContext(ctx)
	if found {
		if _, ok := md["auth-token"]; ok {
			idt := md["auth-token"][0]

			if idt != "" {
				return idt, nil
			}
		}
	}

	md, found = metadata.FromOutgoingContext(ctx)
	if found {
		if _, ok := md["auth-token"]; ok {
			idt := md["auth-token"][0]

			if idt != "" {
				return idt, nil
			}
		}
	}

	return "", status.Errorf(codes.NotFound, "Could not extract token from incoming or outgoing")
}
