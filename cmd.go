package main

import (
	"flag"
	"fmt"
)

/*
Possible commands:
		start-server
		start-worker
		show-configs
*/

const (
	server = "start-server"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}

	cmd := args[0]
	switch cmd {
	case server:
		fmt.Println("starting server")
	default:
		return
	}
}
