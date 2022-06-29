package statsd

import (
	"time"

	"github.com/quipo/statsd"
	"go-illustration/config"
)

func Setup(cfg config.StatsD) (*statsd.StatsdBuffer, error, func()) {
	client := statsd.NewStatsdClient(cfg.Host, cfg.Prefix)
	err := client.CreateSocket()
	if err != nil {
		return nil, err, func() {}
	}

	interval := time.Second * 2
	stats := statsd.NewStatsdBuffer(interval, client)
	return stats, nil, func() {
		stats.Close()
	}
}
