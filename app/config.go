package app

import "github.com/cosmos/cosmos-sdk/container"

func ProvideAppConfig(config *Config) container.Option {
	panic("TODO")
}

func ProvideAppConfigJSON(json []byte) container.Option {
	cfg, err := ReadJSONConfig(json)
	if err != nil {
		return container.Error(err)
	}
	return ProvideAppConfig(cfg)
}

func ProvideAppConfigYAML(yaml []byte) container.Option {
	cfg, err := ReadYAMLConfig(yaml)
	if err != nil {
		return container.Error(err)
	}
	return ProvideAppConfig(cfg)
}
