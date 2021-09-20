package serve_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/honestbank/serve"
)

type config struct {
	Test string `env:"TEST_ENV" default:"my_value"`
}
type configInvalid struct {
	Test string `env:"TEST_ENV" required:"true"`
}

func TestMustNew(t *testing.T) {
	t.Run("works when config can be initialized", func(t *testing.T) {
		cfg := &config{}
		_ = os.Setenv("CONFIG__GRAPHQL_GATEWAY_URL", "GW_URL")
		_ = os.Setenv("CONFIG__PROMETHEUS_PUSH_GATEWAY_URL", "PROM_URL")
		defer func() {
			_ = os.Unsetenv("CONFIG__GRAPHQL_GATEWAY_URL")
			_ = os.Unsetenv("CONFIG__PROMETHEUS_PUSH_GATEWAY_URL")
		}()
		_ = serve.MustNew(cfg, func(r *http.Request) *zap.Logger {
			return nil
		})
		assert.Equal(t, "my_value", cfg.Test)
	})
	t.Run("panics if global config can't be initialized", func(t *testing.T) {
		cfg := &config{}
		assert.Panics(t, func() {
			_ = serve.MustNew(cfg, func(r *http.Request) *zap.Logger {
				return nil
			})
		})
	})
	t.Run("panics when config can't be initialized", func(t *testing.T) {
		cfg := &configInvalid{}
		assert.Panics(t, func() {
			_ = serve.MustNew(cfg, func(r *http.Request) *zap.Logger {
				return nil
			})
		})
	})
}
