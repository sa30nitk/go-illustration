package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	log "github.com/sirupsen/logrus"
	"go-illustration/config"
	"go-illustration/httpapi/server"
	"go-illustration/logger"
	translator "go-illustration/pkg/i18n"
	"go-illustration/pkg/placeholder"
	"go-illustration/statsd"
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

	err, closeFn := logger.Setup(cfg.App)
	if err != nil {
		panic("Failed to set up logger")
	}
	defer closeFn()

	log.Info("logger set up completed")

	statsdClient, err, statsdCloseFun := statsd.Setup(cfg.StatsD)
	if err != nil {
		panic("Failed to set up statsdClient")
	}
	defer statsdCloseFun()
	log.Info("statsd set up completed")

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	trn := translator.NewTranslator()
	if trn == nil {
		log.Info("failed to instantiate translator")
		return
	}

	cmd := args[0]
	switch cmd {
	case startServer:
		log.Debug("Starting server")
		server.StartServer(cfg, server.Dependencies{
			StatsD:      statsdClient,
			PlaceHolder: placeholder.NewClient(cfg.PlaceHolder),
			Translator:  trn,
		})
	}
}
