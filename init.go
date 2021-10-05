package serve

import (
	"github.com/gorilla/mux"

	"github.com/honestbank/serve/config"
)

func New(cfg interface{}, loggerFactory LoggerFactory, configFileList ...string) (App, error) {
	if err := config.Load(cfg, configFileList...); err != nil {
		return nil, err
	}
	globalCfg, err := globalConfig(configFileList...)
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

func globalConfig(files ...string) (*config.GlobalConfig, error) {
	cfg := &config.GlobalConfig{}
	err := config.Load(cfg, files...)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
