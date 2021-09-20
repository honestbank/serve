package serve

import (
	"github.com/gorilla/mux"

	"github.com/honestbank/serve/config"
)

func New(cfg interface{}, loggerFactory LoggerFactory) (App, error) {
	if err := config.Load(cfg); err != nil {
		return nil, err
	}
	globalCfg, err := globalConfig()
	if err != nil {
		return nil, err
	}

	return &app{
		router:        mux.NewRouter(),
		loggerFactory: loggerFactory,
		config:        *globalCfg,
	}, nil
}

func MustNew(cfg interface{}, loggerFactory LoggerFactory) App {
	application, err := New(cfg, loggerFactory)
	if err != nil {
		panic(err)
	}

	return application
}

func globalConfig() (*config.GlobalConfig, error) {
	cfg := &config.GlobalConfig{}
	err := config.Load(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
