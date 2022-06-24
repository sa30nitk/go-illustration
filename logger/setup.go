package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	"go-illustration/config"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

}

func Setup(cfg config.App) (*os.File, error) {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)

	var level log.Level
	if err := level.UnmarshalText([]byte(cfg.LogLevel)); err == nil {
		log.SetLevel(level)
		log.Info("Logger setup completed")

	}
	return file, nil
}
