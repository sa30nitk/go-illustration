package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	"go-illustration/config"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func Setup(cfg config.App) (error, func()) {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err, func() {}
	}
	log.SetOutput(file)

	if level, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(level)
		log.Info("Logger setup completed")

	}
	return nil, func() {
		file.Close()
	}
}
