package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
)

func TestAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "weather/current?q=london", nil)
	if err != nil {
		t.Fatal(err)
	}

	conf := config.NewConfig()
	req.Header.Set("api-key", "secret")

	if conf.ApiKey() == req.Header.Get("api-key") {
		t.Skip("Authenticated.")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCurrentWeather(conf))
	handler.ServeHTTP(rr, req)

	if got := rr.Code; got != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", got, http.StatusUnauthorized)
	}

	got := strings.TrimSpace(rr.Body.String())
	want := `{"code":401,"message":"Authentication failed: check for valid API key."}`
	if got != want {
		t.Errorf("Handler returned unexpected body: got %v want %v", got, want)
	}
}
