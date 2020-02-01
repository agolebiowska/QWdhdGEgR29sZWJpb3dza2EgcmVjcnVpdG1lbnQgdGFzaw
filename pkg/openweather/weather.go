package openweather

import (
	"context"
	"fmt"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/errs"
	"github.com/patrickmn/go-cache"
)

// OpenWeather API docs: https://openweathermap.org/current
type WeatherService service

type CurrentWeatherListResponse struct {
	Items []CurrentWeather `json:"items"`
	Count int              `json:"count"`
}

type CurrentWeather struct {
	Coord      *Coord    `json:"coord,omitempty"`
	Weather    []Weather `json:"weather,omitempty"`
	Base       *string   `json:"base,omitempty"`
	Main       *Main     `json:"main,omitempty"`
	Visibility *int      `json:"visibility,omitempty"`
	Wind       *Wind     `json:"wind,omitempty"`
	Clouds     *Clouds   `json:"clouds,omitempty"`
	Dt         *int64    `json:"dt,omitempty"`
	Sys        *Sys      `json:"sys,omitempty"`
	ID         *int      `json:"id,omitempty"`
	Name       *string   `json:"name,omitempty"`
	Cod        *int      `json:"cod,omitempty"`
}

type Coord struct {
	Lon *float32 `json:"lon,omitempty"`
	Lat *float32 `json:"lat,omitempty"`
}

type Weather struct {
	ID          *int    `json:"id,omitempty"`
	Main        *string `json:"main,omitempty"`
	Description *string `json:"description,omitempty"`
	Icon        *string `json:"icon,omitempty"`
}

type Main struct {
	Temp     *float32 `json:"temp,omitempty"`
	Pressure *float32 `json:"pressure,omitempty"`
	Humidity *float32 `json:"humidity,omitempty"`
	TempMin  *float32 `json:"temp_min,omitempty"`
	TempMax  *float32 `json:"temp_max,omitempty"`
}

type Wind struct {
	Speed *float32 `json:"speed,omitempty"`
	Deg   *float32 `json:"deg,omitempty"`
}

type Clouds struct {
	All *int `json:"all,omitempty"`
}

type Sys struct {
	Type    *int     `json:"type,omitempty"`
	ID      *int     `json:"id,omitempty"`
	Message *float32 `json:"message,omitempty"`
	Country *string  `json:"country,omitempty"`
	Sunrise *int     `json:"sunrise,omitempty"`
	Sunset  *int     `json:"sunset,omitempty"`
}

func (s *WeatherService) ListCurrentByCityNames(ctx context.Context, names []string) (*CurrentWeatherListResponse, error) {
	if len(names) <= 0 || names[0] == "" {
		return nil, errs.ErrBadRequest
	}

	var weathers []CurrentWeather
	for _, n := range names {
		weather := new(CurrentWeather)

		w, found := s.client.Cache.Get(n)
		if !found {
			err := s.client.Do(ctx, fmt.Sprintf("weather?q=%s", n), weather)
			if err != nil {
				return nil, err
			}

			s.client.Cache.Set(n, weather, cache.DefaultExpiration)
		}

		weather = w.(*CurrentWeather)
		weathers = append(weathers, *weather)
	}

	return &CurrentWeatherListResponse{weathers, len(weathers)}, nil
}
