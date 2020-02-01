package api

import (
	"encoding/json"
	"net/http"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/errs"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/openweather"
)

/*
	GET /api/v1/weather/current

	required:
		q, example=warsaw,london,tokyo
	optional:
		limit, default=20
		page, default=1
*/
func CurrentWeather(conf *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := openweather.NewClient(conf)
		f := c.Weather.MakeFilters(r.URL.Query())

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		res, err := c.Weather.ListCurrentByCityNames(r.Context(), f)
		if err != nil {
			errs.WriteError(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			errs.WriteError(w, err)
		}
	}
}
