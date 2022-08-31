package initialize

import (
	"gofound-grpc/global"
	"gofound-grpc/internal/searcher"
	"gofound-grpc/internal/searcher/words"
)

func InitContainer() *searcher.Container {
	tokenizer := NewTokenizer(global.CONFIG.Databases.Path)
	return NewContainer(tokenizer)
}

func NewTokenizer(dictionaryPath string) *words.Tokenizer {
	return words.NewTokenizer(dictionaryPath)
}

func NewContainer(tokenizer *words.Tokenizer) *searcher.Container {
	container := &searcher.Container{
		Dir:       global.CONFIG.Databases.Path,
		Tokenizer: tokenizer,
		Shard:     int(global.CONFIG.Databases.Shard),
	}
	if err := container.Init(); err != nil {
		panic(err)
	}
	return container
}
