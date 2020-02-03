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
	"github.com/patrickmn/go-cache"
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

	if ctx == nil {
		ctx = context.Background()
	}

	resp, err := c.http.Do(req.WithContext(ctx))
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return chooseError(ctx, errs.ErrInvalidRequest)
	}

	if err := errs.CheckResponse(resp); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errs.ErrInvalidResponse
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}

func chooseError(ctx context.Context, err error) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return err
	}
}

func MakeRange(limit, page, len int) (from, to int) {
	from = limit*page - limit
	if from > len {
		from = 0
	}

	return from, len
}
