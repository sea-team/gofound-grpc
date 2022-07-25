package service

import (
	"context"
	gofoundpb "gofound-grpc/api/gen/v1"
)

// Welcome
func (s *GofoundService) Welcome(ctx context.Context, req *gofoundpb.EmptyRequest) (resp *gofoundpb.WelcomeResponse, err error) {
	return &gofoundpb.WelcomeResponse{
		Msg: "Welcome to GoFound",
	}, nil
}
