package config_test

import (
	"testing"

	"github.com/honestbank/serve/config"

	"github.com/stretchr/testify/assert"
)

type myCfg struct {
	MyCfgField string `env:"CONFIG__GRAPHQL_GATEWAY_URL" default:"something"`
}

func TestLoad(t *testing.T) {
	cfg := &myCfg{}
	err := config.Load(cfg)

	assert.NoError(t, err)
	assert.Equal(t, "something", cfg.MyCfgField)
}
