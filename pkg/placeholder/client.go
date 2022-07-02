package placeholder

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	log "github.com/sirupsen/logrus"
)

type Cfg struct {
	Host string
}

type Client struct {
	c   *http.Client
	cfg Cfg
}

func NewClient(cfg Cfg) *Client {
	hystrix.ConfigureCommand("placeholder_ping", hystrix.CommandConfig{
		Timeout:                int(time.Millisecond * 500),
		MaxConcurrentRequests:  3,
		RequestVolumeThreshold: 2,
		SleepWindow:            5000,
		ErrorPercentThreshold:  10,
	})

	c := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   time.Second,
				KeepAlive: time.Second,
			}).DialContext,
			DisableKeepAlives: false,
			MaxIdleConns:      4,
			MaxConnsPerHost:   4,
			IdleConnTimeout:   time.Millisecond * 500,
		},
		Timeout: time.Millisecond * 500,
	}
	return &Client{c: &c, cfg: cfg}
}

func (c *Client) Placeholder(ctx context.Context) *http.Response {
	var res *http.Response
	hystrixErr := hystrix.Do("placeholder_ping", func() error {
		url := fmt.Sprintf("%s/ping", c.cfg.Host)
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		log.Debug(req)
		var err error
		res, err = c.c.Do(req)
		return err
	}, func(err error) error {
		log.Debug("placeholder ping error: ", err)
		return err
	})
	if hystrixErr != nil {
		log.Debug("placeholder ping return error: ", hystrixErr)
		return nil
	}
	log.Debug("placeholder ping return response: ", res)
	return res
}
