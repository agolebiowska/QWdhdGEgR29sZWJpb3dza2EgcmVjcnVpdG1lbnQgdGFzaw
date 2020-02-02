package config

import (
	"time"

	gc "github.com/patrickmn/go-cache"
)

type Config struct {
	cache  *gc.Cache
	config *Params
}

func NewConfig() *Config {
	c := &Config{
		config: NewParams(),
	}

	return c
}

func (c *Config) OpenWeatherApiKey() string {
	return c.config.OpenWeatherApiKey
}

func (c *Config) OpenWeatherApiBaseUrl() string {
	return c.config.OpenWeatherApiBaseUrl
}

func (c *Config) HttpServerPort() int {
	if c.config.HttpServerPort < 1 || c.config.HttpServerPort > 65535 {
		c.config.HttpServerPort = 8080
	}

	return c.config.HttpServerPort
}

func (c *Config) ApiKey() string {
	return c.config.ApiKey
}

// Cache returns the in-memory cache.
func (c *Config) Cache() *gc.Cache {
	if c.cache == nil {
		c.cache = gc.New(
			time.Duration(c.config.CacheExpiration)*time.Minute,
			time.Duration(c.config.CacheInterval)*time.Minute)
	}
	return c.cache
}
