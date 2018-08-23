package models

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
