package service

import (
	"context"
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/global"
	"gofound-grpc/internal/searcher"
	"gofound-grpc/internal/searcher/model"
)

var (
	index *Index
)

// Index 索引管理
type Index struct {
	Container *searcher.Container
}

func NewIndex() *Index {
	index = &Index{
		Container: global.Container,
	}
	return index
}

func (s *GofoundService) Index(ctx context.Context, req *gofoundpb.SingleIndexRequest) (*gofoundpb.OperationResponse, error) {
	idxDoc := &model.IndexDoc{
		Id:       uint32(req.IndexDoc.Id),
		Text:     req.IndexDoc.Text,
		Document: req.IndexDoc.Document.AsMap(),
	}
	err := index.Container.GetDataBase(req.GetDatabase()).IndexDocument(idxDoc)
	if err != nil {
		return &gofoundpb.OperationResponse{
			State:   false,
			Message: err.Error(),
		}, err
	}
	return &gofoundpb.OperationResponse{
		State:   true,
		Message: "success",
	}, nil
}

func (s *GofoundService) BatchIndex(ctx context.Context, req *gofoundpb.BatchIndexRequest) (*gofoundpb.OperationResponse, error) {
	for _, v := range req.GetIndexDocs() {
		doc := &model.IndexDoc{
			Id:       uint32(v.Id),
			Text:     v.Text,
			Document: v.Document.AsMap(),
		}
		err := index.Container.GetDataBase(req.GetDatabase()).IndexDocument(doc)
		if err != nil {
			return &gofoundpb.OperationResponse{
				State:   false,
				Message: err.Error(),
			}, err
		}
	}
	return &gofoundpb.OperationResponse{
		State:   true,
		Message: "success",
	}, nil
}

func (s *GofoundService) RemoveIndex(ctx context.Context, req *gofoundpb.RemoveIndexRequest) (*gofoundpb.OperationResponse, error) {
	err := index.Container.GetDataBase(req.Database).RemoveIndex(uint32(req.GetId()))
	if err != nil {
		return &gofoundpb.OperationResponse{
			State:   false,
			Message: err.Error(),
		}, err
	}
	return &gofoundpb.OperationResponse{
		State:   true,
		Message: "success",
	}, nil
}
