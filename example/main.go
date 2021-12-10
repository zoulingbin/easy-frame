package main

import (
	"github.com/zoulingbin/easy-frame/framework"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	core := framework.NewCore()
	RegisterRoute(core)
	server := http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
