package openweather

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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

func NewClient(token string, baseUrl string) *Client {
	baseURL, _ := url.Parse(baseUrl)
	c := &Client{
		http: &http.Client{},
		BaseURL: baseURL,
		Token: token,
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
		fmt.Print(err.Error())
		os.Exit(1)
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return err
	}

	return nil
}
