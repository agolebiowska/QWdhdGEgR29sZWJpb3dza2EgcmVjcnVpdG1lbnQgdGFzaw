package config

import (
	"os"
	"strconv"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tc := NewConfig()

	p, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if tc.HttpServerPort() != p {
		t.Errorf("Config returned wrong port: got %v want %v", tc.HttpServerPort(), p)
	}

	ak := os.Getenv("API_KEY")
	if tc.ApiKey() != ak {
		t.Errorf("Config returned wrong api key: got %v want %v", tc.ApiKey(), ak)
	}

	ok := os.Getenv("OPEN_WEATHER_API_KEY")
	if tc.OpenWeatherApiKey() != ok {
		t.Errorf("Config returned wrong open weather key: got %v want %v", tc.OpenWeatherApiKey(), ok)
	}

	ou := os.Getenv("OPEN_WEATHER_API_BASE_URL")
	if tc.OpenWeatherApiBaseUrl() != ou {
		t.Errorf("Config returned wrong open weather base url: got %v want %v", tc.OpenWeatherApiBaseUrl(), ou)
	}
}
