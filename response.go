package serve

import (
	"encoding/json"
	"net/http"
)

func (a app) JSON(status int, val interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(val)
}
