package openweather

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	u "github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/utils"
)

func TestWeatherService_ListCurrentByNames(t *testing.T) {
	c, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write(weatherJSON)
	})

	filters := &Filters{[]string{"warsaw"}, 1, 20}
	got, err := c.Weather.ListCurrentByNames(context.Background(), filters)

	if err != nil {
		t.Errorf("Weather.ListCurrentByNames returned error: %v", err)
	}

	if want := wantResponse; !reflect.DeepEqual(got, want) {
		t.Errorf("Weather.ListCurrentByNames = %+v, want %+v", got, want)
	}
}

var weatherJSON = []byte(`{
 "coord": {
   "lon": 21.01,
   "lat": 52.23
 },
 "weather": [
   {
     "id": 801,
     "main": "Clouds",
     "description": "few clouds",
     "icon": "02n"
   }
 ],
 "base": "stations",
 "main": {
   "temp": 279.3,
   "feels_like": 274.44,
   "temp_min": 278.15,
   "temp_max": 280.93,
   "pressure": 1001,
   "humidity": 87
 },
 "visibility": 10000,
 "wind": {
   "speed": 5.1,
   "deg": 280
 },
 "clouds": {
   "all": 20
 },
 "dt": 1580659814,
 "sys": {
   "type": 1,
   "id": 1713,
   "country": "PL",
   "sunrise": 1580624175,
   "sunset": 1580656969
 },
 "timezone": 3600,
 "id": 756135,
 "name": "Warsaw",
 "cod": 200
}`)

var wantResponse = &CurrentWeatherListResponse{
	Items: []CurrentWeather{
		{
			Coord: &Coord{
				Lon: u.Float32(21.01),
				Lat: u.Float32(52.23),
			},
			Weather: []Weather{
				{
					ID:          u.Int(801),
					Main:        u.String("Clouds"),
					Description: u.String("few clouds"),
					Icon:        u.String("02n"),
				},
			},
			Base: u.String("stations"),
			Main: &Main{
				Temp:      u.Float32(279.3),
				FeelsLike: u.Float32(274.44),
				TempMin:   u.Float32(278.15),
				TempMax:   u.Float32(280.93),
				Pressure:  u.Float32(1001),
				Humidity:  u.Float32(87),
			},
			Visibility: u.Int(10000),
			Wind: &Wind{
				Speed: u.Float32(5.1),
				Deg:   u.Float32(280),
			},
			Clouds: &Clouds{
				All: u.Int(20),
			},
			Dt: u.Int64(1580659814),
			Sys: &Sys{
				Type:    u.Int(1),
				ID:      u.Int(1713),
				Country: u.String("PL"),
				Sunrise: u.Int(1580624175),
				Sunset:  u.Int(1580656969),
			},
			Timezone: u.Int(3600),
			ID:       u.Int(756135),
			Name:     u.String("Warsaw"),
			Cod:      u.Int(200),
		},
	},
	Count: 1,
}

func TestWeatherService_MakeFilters(t *testing.T) {
	c, _, teardown := setup()
	defer teardown()

	tests := []struct {
		name  string
		query string
		want  *Filters
	}{
		{
			"Empty query",
			"",
			&Filters{nil, 1, 20},
		},
		{
			"One name",
			"q=warsaw",
			&Filters{[]string{"warsaw"}, 1, 20},
		},
		{
			"Multiple names",
			"q=warsaw,london,new york,budapest",
			&Filters{[]string{"warsaw", "london", "new york", "budapest"}, 1, 20},
		},
		{
			"All filters",
			"q=warsaw,london&page=15&limit=5",
			&Filters{[]string{"warsaw", "london"}, 15, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, _ := url.ParseQuery(tt.query)
			got := c.Weather.MakeFilters(q)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
