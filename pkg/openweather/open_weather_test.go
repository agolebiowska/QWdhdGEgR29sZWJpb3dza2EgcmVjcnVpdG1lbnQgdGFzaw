package openweather

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	client = NewClient(config.GetTestConfig())
	u, _ := url.Parse(server.URL)
	client.BaseURL = u

	return client, mux, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestMakeRange(t *testing.T) {
	tests := []struct {
		limit    int
		page     int
		len      int
		wantFrom int
		wantTo   int
	}{
		{20, 1, 10, 0, 10},
		{5, 2, 10, 5, 10},
		{5, 3, 11, 10, 11},
		{5, 20, 100, 95, 100},
		{63, 54, 10123, 3339, 10123},
	}
	for _, tt := range tests {
		t.Run("Range test", func(t *testing.T) {
			gotFrom, gotTo := MakeRange(tt.limit, tt.page, tt.len)

			if gotFrom != tt.wantFrom {
				t.Errorf("got from = %v, want form %v", gotFrom, tt.wantFrom)
			}

			if gotTo != tt.wantTo {
				t.Errorf("got to = %v, want to %v", gotTo, tt.wantTo)
			}
		})
	}
}
