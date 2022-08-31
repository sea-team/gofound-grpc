package service

import (
	gofoundpb "gofound-grpc/api/gen/v1"
)

var srv *Services

type GofoundService struct {
	gofoundpb.UnimplementedGofoundServiceServer
}

// NewGofoundService 初始化服务
func NewGofoundService() *GofoundService {
	return &GofoundService{}
}

type Services struct {
	Base *Base
}

func NewServices() {
	srv = &Services{
		Base: NewBase(),
	}
}
