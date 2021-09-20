package measurements

import (
	"github.com/prometheus/client_golang/prometheus/push"
)

type Prometheus interface {
	GetClient(metricName string) *push.Pusher
}

type PrometheusImpl struct {
	URL string
}

func (p *PrometheusImpl) GetClient(metricName string) *push.Pusher {
	return push.New(p.URL, metricName)
}

func NewPrometheus(url string) Prometheus {
	return &PrometheusImpl{
		URL: url,
	}
}
