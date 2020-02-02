package config

import (
	"testing"

	"github.com/caarlos0/env"
)

func TestNewParams(t *testing.T) {
	p := Params{}
	if err := env.Parse(&p); err != nil {
		t.Fatal("Error while parsing .env file: ", err)
	}
}
