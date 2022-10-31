package system

import (
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/global"
	"os"
	"runtime"
)

func Env() *gofoundpb.System {
	return &gofoundpb.System{
		Os:             runtime.GOOS,
		Arch:           runtime.GOARCH,
		Cores:          int32(runtime.NumCPU()),
		Version:        runtime.Version(),
		Goroutines:     int32(runtime.NumGoroutine()),
		DataPath:       global.CONFIG.Databases.Path,
		DictionaryPath: global.CONFIG.Dictionary.Path,
		Shard:          global.CONFIG.Databases.Shard,
		Gomaxprocs:     int32(runtime.NumCPU() * 2),
		// DataSize: GetFloat64MB(utils.DirSizeB(global.CONFIG.Databases.Path)),
		// Dbs: ,
		Executable: os.Args[0],
		Pid:        int32(os.Getegid()),
		EnableAuth: global.CONFIG.Auth.EnableAdmin,
	}
}
