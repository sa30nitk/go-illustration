package statsd

import (
	"github.com/quipo/statsd"
	"go-illustration/config"
)

func Setup(cfg config.StatsD) (*statsd.StatsdClient, error, func()) {
	client := statsd.NewStatsdClient(cfg.Host, cfg.Prefix)
	err := client.CreateSocket()
	if err != nil {
		return nil, err, func() {}
	}
	return client, nil, func() {
		client.Close()
	}
}
