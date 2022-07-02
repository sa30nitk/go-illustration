package placeholder

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

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
	c := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   time.Second * 2,
				KeepAlive: time.Second * 2,
			}).DialContext,
			DisableKeepAlives: false,
			MaxIdleConns:      4,
			MaxConnsPerHost:   4,
			IdleConnTimeout:   time.Second * 5,
		},
		Timeout: time.Second * 10,
	}
	return &Client{c: &c, cfg: cfg}
}

func (c *Client) Placeholder(ctx context.Context) *http.Response {
	url := fmt.Sprintf("%s/ping", c.cfg.Host)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	log.Debug(req)
	res, err := c.c.Do(req)
	if err != nil {
		return nil
	}
	log.Debug(res)
	return res

}
