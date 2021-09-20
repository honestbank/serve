package serve

import "github.com/honestbank/serve/measurements"

func (a app) MeasurementClient() measurements.Prometheus {
	return measurements.NewPrometheus(a.config.PrometheusPushGatewayURL)
}
