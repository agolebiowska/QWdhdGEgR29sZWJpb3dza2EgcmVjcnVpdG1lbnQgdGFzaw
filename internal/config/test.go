package config

import (
	gc "github.com/patrickmn/go-cache"
	"time"
)

func GetTestConfig() *Config {
	params := &Params{
		HttpServerPort:        8080,
		ApiKey:                "secret",
		OpenWeatherApiKey:     "xxx",
		OpenWeatherApiBaseUrl: "api.url",
		CacheExpiration:       1,
		CacheInterval:         2,
	}

	return &Config{
		gc.New(
			time.Duration(params.CacheExpiration)*time.Minute,
			time.Duration(params.CacheInterval)*time.Minute),
		params,
	}
}
