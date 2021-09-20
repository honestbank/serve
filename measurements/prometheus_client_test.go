package measurements_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/honestbank/serve/measurements"
)

func TestPrometheusImpl_GetClient(t *testing.T) {
	prometheus := measurements.NewPrometheus("url")
	client := prometheus.GetClient("metric")

	assert.NotNil(t, client)
}
