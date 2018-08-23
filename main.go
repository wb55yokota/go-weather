package main

import (
	"fmt"
	"log"

	"github.com/wb55yokota/go-weather/client"
	"github.com/wb55yokota/go-weather/libs"
)

func main() {
	// 取得
	client := client.New()
	weather, err := client.SendRequest("GET", "http://weather.livedoor.com/forecast/webservice/json/v1?city=130010")
	if err != nil {
		log.Fatal(err)
	}

	// 表示
	tomorrowWeather := weather.Forecasts[1]
	str := fmt.Sprintf("%sの天気\n%s (%s) 最高：%s度 / 最低：%s度\n",
		weather.Location.City,
		tomorrowWeather.DataLabel,
		tomorrowWeather.Telop,
		tomorrowWeather.Temperature.Max["celsius"],
		tomorrowWeather.Temperature.Min["celsius"])
	fmt.Print(str)

	// save
	_, err = libs.SaveStringToFile("/tmp/output.txt", str)
	if err != nil {
		log.Fatal(err)
	}
}
