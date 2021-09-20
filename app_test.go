package serve_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/honestbank/serve"
)

type val struct {
	Hello string `json:"hello"`
}

func ExampleNew() {
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
	app.AddHandler("/hello", func(w http.ResponseWriter, r *http.Request) {
		app.JSON(200, &val{Hello: "world"}, w)
	}).Methods("GET")
	go app.Start()
	time.Sleep(time.Millisecond * 100)
	resp, _ := http.DefaultClient.Get("http://localhost:8080/hello")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	// Output: {"hello":"world"}
}
