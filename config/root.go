package config

import "github.com/jinzhu/configor"

type GlobalConfig struct {
	GraphQLGatewayURL        string `env:"CONFIG__GRAPHQL_GATEWAY_URL" required:"true"`
	PrometheusPushGatewayURL string `env:"CONFIG__PROMETHEUS_PUSH_GATEWAY_URL" required:"true"`
}

func Load(cfg interface{}) error {
	return configor.
		New(&configor.Config{AutoReload: false, Silent: true}).
		Load(cfg)
}
