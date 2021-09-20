package serve

import (
	"net/http"

	"go.uber.org/zap"
)

func (a app) Logger(r *http.Request) *zap.Logger {
	return a.loggerFactory(r)
}
