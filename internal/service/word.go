package service

import (
	"context"
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/global"
	"gofound-grpc/internal/searcher"
)

var (
	word *Word
)

// Word 分词管理
type Word struct {
	Container *searcher.Container
}

func NewWord() *Word {
	word = &Word{
		Container: global.Container,
	}
	return word
}

func (s *GofoundService) WordCut(ctx context.Context, req *gofoundpb.WordCutRequest) (*gofoundpb.WordCutResponse, error) {
	return &gofoundpb.WordCutResponse{
		Word: word.Container.Tokenizer.Cut(req.GetKeyWord()),
	}, nil
}
