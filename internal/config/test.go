package config

import (
	gc "github.com/patrickmn/go-cache"
	"time"
)

func GetTestConfig() *Config {
	params := &Params{
		HttpServerPort:        8080,
		OpenWeatherApiKey:     "xxx",
		OpenWeatherApiBaseUrl: "",
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
