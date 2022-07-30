package service

import (
	"context"
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/internal/searcher/system"
	"runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Welcome 欢迎语
func (s *GofoundService) Welcome(ctx context.Context, req *gofoundpb.EmptyRequest) (resp *gofoundpb.WelcomeResponse, err error) {
	return &gofoundpb.WelcomeResponse{
		Msg: "Welcome to GoFound",
	}, nil
}

// GC 释放GC
func (s *GofoundService) GC(ctx context.Context, req *gofoundpb.EmptyRequest) (*gofoundpb.EmptyResponse, error) {
	go runtime.GC()
	return &gofoundpb.EmptyResponse{}, nil
}

// Status 服务器状态
func (s *GofoundService) Status(ctx context.Context, req *gofoundpb.EmptyRequest) (*gofoundpb.StatusResponse, error) {
	// TODO：获取每秒cpu使用频率比较慢
	cup, err := system.GetCPUInfo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	disk, err := system.GetDiskInfo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	memory, err := system.GetMemInfo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gofoundpb.StatusResponse{
		Cup:    cup,
		Disk:   disk,
		Memory: memory,
		System: system.Env(),
	}, nil
}

// Query 查询
func (s *GofoundService) Query(ctx context.Context, req *gofoundpb.QueryRequest) (*gofoundpb.QueryResponse, error) {
	return &gofoundpb.QueryResponse{}, nil
}
