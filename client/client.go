package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wb55yokota/go-weather/models"
)

type Client struct {
	client http.Client
}

func New() Client {
	c := Client{http.Client{}}
	return c
}

func getRequest(method string, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}

func (c Client) SendRequest(method string, url string) (*models.Weather, error) {
	// request生成
	req, err := getRequest(method, url)
	if err != nil {
		return nil, err
	}

	// request送信
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// responseBodyを読み込む
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// parseした内容を返す
	return parseJson(bytes)
}

func parseJson(bytes []byte) (*models.Weather, error) {
	// json parse
	var weather models.Weather
	if err := json.Unmarshal(bytes, &weather); err != nil {
		return nil, err
	}
	return &weather, nil
}
