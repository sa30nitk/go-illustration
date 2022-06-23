package main

import (
	"flag"
	"fmt"

	"go-illustration/config"
	"go-illustration/httpapi/server"
)

/*
Possible commands:
		start-server
		start-worker
		show-configs
*/

const (
	startServer = "start-server"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	config.Load()
	cfg := config.NewConfig()
	cmd := args[0]
	switch cmd {
	case startServer:
		fmt.Println("starting server")
		server.StartServer(cfg)
	default:
		return
	}
}
