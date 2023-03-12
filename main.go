package main

import (
	"beiwanglu/conf"
	"beiwanglu/routes"
)

func main() {
	// 从配置文件读入配置
	conf.Init()
	// 转载路由 swag init -g common.go
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
