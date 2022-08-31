package service

import (
	"context"
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/global"
	"gofound-grpc/internal/searcher"
	"gofound-grpc/internal/searcher/model"
	"gofound-grpc/internal/searcher/system"
	"runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	base *Base
)

// Base 基础管理
type Base struct {
	Container *searcher.Container
}

func NewBase() *Base {
	base = &Base{
		Container: global.Container,
	}
	return base
}

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
	request := &model.SearchRequest{
		Query:    req.Query,
		Order:    req.Order,
		ScoreExp: req.ScoreExp,
		Page:     req.Page,
		Limit:    req.Limit,
		Highlight: &model.Highlight{
			PreTag:  req.Highlight.PreTag,
			PostTag: req.Highlight.PostTag,
		},
		Database: req.Database,
	}
	res, err := base.Container.GetDataBase(req.Database).MultiSearch(request)
	if err != nil {
		return nil, err
	}
	docs := make([]gofoundpb.ResponseDoc, len(res.Documents))
	for _, v := range res.Documents {
		docs = append(docs, gofoundpb.ResponseDoc{
			Id:           v.Id,
			Text:         v.Text,
			Document:     v.Document,
			OriginalText: v.OriginalText,
			Score:        v.Score,
			Keys:         v.Keys,
		})
	}
	if len(docs) == 0 {
		docs = append(docs, gofoundpb.ResponseDoc{})
	}
	return &gofoundpb.QueryResponse{
		Time:      res.Time,
		Total:     res.Total,
		PageCount: res.PageCount,
		Page:      res.Page,
		Limit:     res.Limit,
		Documents: &docs[0],
		Words:     res.Words,
	}, nil
}
