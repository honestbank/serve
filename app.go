package serve

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/honestbank/serve/config"
	"github.com/honestbank/serve/measurements"
)

type LoggerFactory func(r *http.Request) *zap.Logger

type app struct {
	router        *mux.Router
	config        config.GlobalConfig
	loggerFactory LoggerFactory
}

type App interface {
	AddHandler(url string, handler http.HandlerFunc) *mux.Route
	JSON(status int, val interface{}, w http.ResponseWriter)
	// MeasurementClient - once golang finally introduces generics we will no longer hardcode prometheus
	MeasurementClient() measurements.Prometheus
	Logger(r *http.Request) *zap.Logger
	Start()
}
