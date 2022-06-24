package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"go-illustration/config"
	"go-illustration/httpapi/server"
	"go-illustration/logger"
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

	if err := config.Load(); err != nil {
		panic("Failed to load configs")
	}

	log.Info("config loaded")
	cfg := config.NewConfig()

	if err := logger.Setup(cfg.App); err != nil {
		panic("Failed to set up logger")
	}

	log.Info("logger set up completed")

	cmd := args[0]
	switch cmd {
	case startServer:
		log.Debug("Starting server")
		if err := server.StartServer(cfg); err != nil {
			log.Errorf("Failed to launch server with error: %s\n", err)
		}
	}
}
