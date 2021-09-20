package serve_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/honestbank/serve"
)

func TestApp_MeasurementClient(t *testing.T) {
	t.Run("a measurement client is available", func(t *testing.T) {
		cfg := &config{}
		_ = os.Setenv("CONFIG__GRAPHQL_GATEWAY_URL", "GW_URL")
		_ = os.Setenv("CONFIG__PROMETHEUS_PUSH_GATEWAY_URL", "PROM_URL")
		defer func() {
			_ = os.Unsetenv("CONFIG__GRAPHQL_GATEWAY_URL")
			_ = os.Unsetenv("CONFIG__PROMETHEUS_PUSH_GATEWAY_URL")
		}()
		app := serve.MustNew(cfg, func(r *http.Request) *zap.Logger {
			return zap.NewNop()
		})
		request := httptest.NewRequest("GET", "/", nil)
		assert.NotNil(t, app.Logger(request))
		assert.NotNil(t, app.MeasurementClient().GetClient("something"))
	})
}
