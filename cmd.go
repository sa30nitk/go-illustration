package main

import (
	"fmt"
	"os"
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
	// Removes program name
	args := os.Args[1:]
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
