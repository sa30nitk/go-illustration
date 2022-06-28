package logger

import (
	"os"

	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"go-illustration/config"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

func Setup(cfg config.App) (error, func()) {
	pathMap := lfshook.PathMap{
		log.DebugLevel: "application.log",
	}

	log.AddHook(lfshook.NewHook(pathMap, &log.JSONFormatter{}))
	if level, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(level)
		log.Info("Logger setup completed")

	}
	return nil, func() {}
}
