package global

import (
	"gofound-grpc/config"
	"gofound-grpc/internal/searcher"
)

var (
	CONFIG    config.Server
	Container *searcher.Container
)
