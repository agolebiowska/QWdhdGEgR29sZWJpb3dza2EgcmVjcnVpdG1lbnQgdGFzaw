package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/errs"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/openweather"
)

// GET /api/v1/weather/current?q=Warsaw,London
func CurrentWeatherByCityNames(conf *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		names := strings.Split(v.Get("q"), ",")

		c := openweather.NewClient(conf.OpenWeatherApiKey(), conf.OpenWeatherApiBaseUrl())

		res, err := c.Weather.ListCurrentByCityNames(r.Context(), names)
		if err != nil {
			errs.WriteError(w, err)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			errs.WriteError(w, err)
		}
	}
}