package openweather

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/pkg/errs"
)

type Client struct {
	http    *http.Client
	BaseURL *url.URL
	Token   string

	Weather *WeatherService
}

type service struct {
	client *Client
}

func NewClient(conf *config.Config) *Client {
	baseURL, _ := url.Parse(conf.OpenWeatherApiBaseUrl())
	c := &Client{
		http:    &http.Client{},
		BaseURL: baseURL,
		Token:   conf.OpenWeatherApiKey(),
	}

	c.Weather = &WeatherService{c}

	return c
}

func (c *Client) Do(ctx context.Context, urlStr string, result interface{}) error {
	u, err := c.BaseURL.Parse(fmt.Sprintf("%s&appid=%s", urlStr, c.Token))
	if err != nil {
		return err
	}

	response, err := c.http.Get(u.String())
	if err != nil {
		return errs.ErrInvalidRequest
	}

	if err := errs.FindError(response); err != nil {
		return err
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errs.ErrInvalidResponse
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return err
	}

	return nil
}
