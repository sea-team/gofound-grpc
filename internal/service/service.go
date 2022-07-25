package service

import (
	gofoundpb "gofound-grpc/api/gen/v1"
)

type GofoundService struct {
	gofoundpb.UnimplementedGofoundServiceServer
}

// NewGofoundService 初始化服务
func NewGofoundService() *GofoundService {
	return &GofoundService{}
}
