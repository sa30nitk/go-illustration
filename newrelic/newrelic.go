package newrelic

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"go-illustration/config"
)

func NRApp(cfg config.NewRelic) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.App),
		newrelic.ConfigLicense(cfg.License),
		newrelic.ConfigEnabled(cfg.Enabled))
}
