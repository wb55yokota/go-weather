package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Weather struct {
	PinpointLocations []PinpointLocation `json:"pinpointLocations"`
	Forecasts         []Forecast         `json:"forecasts"`
	Location          Location           `json:"location"`
	Desctiption       Description        `json:"description"`
}

type PinpointLocation struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

type Forecast struct {
	DataLabel   string      `json:"dateLabel"`
	Telop       string      `json:"telop"`
	Date        string      `json:"date"`
	Temperature Temperature `json:"temperature"`
	Image       Image       `json:"image"`
}

type Temperature struct {
	Min map[string]string `json:"min"`
	Max map[string]string `json:"max"`
}

type Image struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
	Title  string `json:"title"`
}

type Location struct {
	City       string `json:"city"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
}

type Description struct {
	Text       string `json:"text"`
	PublicTime string `json:"publicTime"`
}

func cp932(str string) string {
	r, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	return r
}

func file_save(str string) {
	file, err := os.Create("/tmp/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes := []byte(cp932(str))
	_, err = file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// request作る
	req, err := http.NewRequest("GET", "http://weather.livedoor.com/forecast/webservice/json/v1?city=130010", nil)
	if err != nil {
		log.Fatal(err)
	}

	// client作る
	client := &http.Client{}

	// request投げる
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// responseBodyを読み込む
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// json parse
	var weather Weather
	if err := json.Unmarshal(bytes, &weather); err != nil {
		log.Fatal(err)
	}

	// 表示処理
	tomorrow := weather.Forecasts[1]
	str := fmt.Sprintf("%sの天気\n%s (%s) 最高：%s度 / 最低：%s度\n",
		weather.Location.City,
		tomorrow.DataLabel,
		tomorrow.Telop,
		tomorrow.Temperature.Max["celsius"],
		tomorrow.Temperature.Min["celsius"])
	fmt.Print(str)

	// save
	file_save(str)
}
