package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Params struct {
	HttpServerPort        int    `env:"HTTP_PORT" envDefault:"8080"`
	OpenWeatherApiKey     string `env:"OPEN_WEATHER_API_KEY"`
	OpenWeatherApiBaseUrl string `env:"OPEN_WEATHER_API_BASE_URL" envDefault:"https://api.openweathermap.org/data/2.5/"`
	CacheExpiration       int    `env:"CACHE_EXPIRATION" envDefault:"30"`
	CacheInterval         int    `env:"CACHE_INTERVAL" envDefault:"60"`
}

func NewParams() *Params {
	p := Params{}
	if err := env.Parse(&p); err != nil {
		log.Fatal(err)
	}

	return &p
}
