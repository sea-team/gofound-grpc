package service

import (
	"context"
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/global"
	"gofound-grpc/internal/searcher"
)

var (
	database *Database
)

// Database 数据源管理
type Database struct {
	Container *searcher.Container
}

func NewDatabase() *Database {
	database = &Database{
		Container: global.Container,
	}
	return database
}
func (s *GofoundService) ShowDatabase(ctx context.Context, req *gofoundpb.EmptyRequest) (*gofoundpb.DatabaseResponse, error) {
	r := database.Container.GetDataBases()
	res := make(map[string]*gofoundpb.Engine)
	for i, v := range r {
		res[i] = &gofoundpb.Engine{
			IndexPath: v.IndexPath,
			Option: &gofoundpb.Options{
				InvertedIndexName: v.Option.InvertedIndexName,
				PositiveIndexName: v.Option.PositiveIndexName,
				DocIndexName:      v.Option.DocIndexName,
			},
			IsDebug:      v.IsDebug,
			DatabaseName: v.DatabaseName,
			Shard:        int32(v.Shard),
			Timeout:      v.Timeout,
			BuffNum:      int64(v.BufferNum),
		}
	}
	return &gofoundpb.DatabaseResponse{
		Len:       int32(len(res)),
		Databases: res,
	}, nil
}

func (s *GofoundService) DropDatabase(ctx context.Context, req *gofoundpb.DatabaseRequest) (*gofoundpb.OperationResponse, error) {
	err := database.Container.DropDataBase(req.GetDatabaseName())
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

func (s *GofoundService) CreateDatabase(ctx context.Context, req *gofoundpb.DatabaseRequest) (*gofoundpb.OperationResponse, error) {
	database.Container.GetDataBase(req.GetDatabaseName())
	return &gofoundpb.OperationResponse{
		State:   true,
		Message: "success",
	}, nil
}
