package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wb55yokota/go-weather/client"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

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
	// 取得
	client := client.New()
	weather, err := client.SendRequest("GET", "http://weather.livedoor.com/forecast/webservice/json/v1?city=130010")
	if err != nil {
		log.Fatal(err)
	}

	// 表示
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
