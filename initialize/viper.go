package initialize

import (
	"log"

	"gofound-grpc/global"

	"github.com/spf13/viper"
)

// InitViper 初始化解析器
func InitViper(path string) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error read config file: %s \n", err)
	}

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		log.Fatalf("Fatal error unmarshal config file: %s \n", err)
	}
}
