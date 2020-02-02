package openweather

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/agolebiowska/QWdhdGEgR29sZWJpb3dza2EgcmVjcnVpdG1lbnQgdGFzaw/internal/config"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultLimit = 20
	defaultPage  = 1
)

type Client struct {
	http    *http.Client
	BaseURL *url.URL
	Token   string
	Cache   *cache.Cache
	Page    int
	Limit   int

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
		Cache:   conf.Cache(),
		Page:    defaultPage,
		Limit:   defaultLimit,
	}

	c.Weather = &WeatherService{c}

	return c
}

func (c *Client) NewRequest(urlStr string) (*http.Request, error) {
	u, err := c.BaseURL.Parse(fmt.Sprintf("%s&appid=%s", urlStr, c.Token))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, urlStr string, result interface{}) error {
	req, err := c.NewRequest(urlStr)
	if err != nil {
		return err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	resp, err := c.http.Do(req)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return err
	}

	var body []byte
	done := make(chan struct{})
	go func() {
		body, err = ioutil.ReadAll(resp.Body)
		close(done)
	}()

	select {
	case <-ctx.Done():
		<-done
		err = resp.Body.Close()
		if err == nil {
			err = ctx.Err()
		}
	case <-done:
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) MakeRange(limit, page, len int) (from, to int) {
	from = limit*page - limit
	if from > len {
		from = 0
	}

	return from, len
}
