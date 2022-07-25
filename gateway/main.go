package main

import (
	"context"
	"flag"
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/global"
	"gofound-grpc/initialize"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// 配置地址，默认路径:./config/config.yaml
	var config string
	flag.StringVar(&config, "conf", "../config/config.yaml", "choose config file.")
	flag.Parse()

	// 初始化解析器
	initialize.InitViper(config)

	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	if err := gofoundpb.RegisterGofoundServiceHandlerFromEndpoint(c, mux, global.CONFIG.GRPC.Addr, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Fatalf("cannot register service:%v", err)
	}

	log.Printf("grpc gatway started at %s", global.CONFIG.HTTP.Addr)
	log.Fatal(http.ListenAndServe(global.CONFIG.HTTP.Addr, mux))
}
