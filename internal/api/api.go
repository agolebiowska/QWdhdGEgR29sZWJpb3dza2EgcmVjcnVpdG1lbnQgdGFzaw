package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/errs"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/openweather"
)

// GET /api/v1/weather/current?q=warsaw,london
func CurrentWeather(conf *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		names := strings.Split(strings.ToLower(v.Get("q")), ",")

		c := openweather.NewClient(conf)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res, err := c.Weather.ListCurrentByCityNames(r.Context(), names)
		if err != nil {
			errs.WriteError(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			errs.WriteError(w, err)
		}
	}
}