package serve

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a app) AddHandler(url string, handler http.HandlerFunc) *mux.Route {
	return a.router.Handle(url, handler)
}
