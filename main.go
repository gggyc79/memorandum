package main

import (
	"beiwanglu/conf"
	"beiwanglu/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
