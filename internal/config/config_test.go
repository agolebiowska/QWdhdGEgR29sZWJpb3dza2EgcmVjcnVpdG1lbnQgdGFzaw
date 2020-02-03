package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	tc := GetTestConfig()

	p := 8080
	if tc.HttpServerPort() != p {
		t.Errorf("Config returned wrong port: got %v want %v", tc.HttpServerPort(), p)
	}

	ak := "secret"
	if tc.ApiKey() != ak {
		t.Errorf("Config returned wrong api key: got %v want %v", tc.ApiKey(), ak)
	}

	ok := "xxx"
	if tc.OpenWeatherApiKey() != ok {
		t.Errorf("Config returned wrong open weather key: got %v want %v", tc.OpenWeatherApiKey(), ok)
	}

	ou := "api.url"
	if tc.OpenWeatherApiBaseUrl() != ou {
		t.Errorf("Config returned wrong open weather base url: got %v want %v", tc.OpenWeatherApiBaseUrl(), ou)
	}
}
