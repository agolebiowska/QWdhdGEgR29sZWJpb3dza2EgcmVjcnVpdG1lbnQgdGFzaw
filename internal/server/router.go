package server

import (
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/api"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(conf *config.Config) http.Handler {
	r := mux.NewRouter()
	a := r.PathPrefix("/api/v1").Subrouter()

	a.HandleFunc("/weather/current", api.CurrentWeatherByCityNames(conf)).Methods("GET")

	return r
}