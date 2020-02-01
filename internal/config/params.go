package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Params struct {
	HttpServerPort        int    `env:"HTTP_PORT" envDefault:"8080"`
	OpenWeatherApiKey     string `env:"OPEN_WEATHER_API_KEY"`
	OpenWeatherApiBaseUrl string `env:"OPEN_WEATHER_API_BASE_URL" envDefault:"https://api.openweathermap.org/data/2.5/"`
	CacheExpiration       int    `env:"CACHE_EXPIRATION" envDefault:"336"`
	CacheInterval         int    `env:"CACHE_INTERVAL" envDefault:"30"`
}

func NewParams() *Params {
	p := Params{}
	if err := env.Parse(&p); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &p
}
