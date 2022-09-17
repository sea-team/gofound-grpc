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

type Services struct {
	Base  *Base
	Index *Index
}

func NewServices() {
	NewBase()
	NewIndex()
	NewDatabase()
}
