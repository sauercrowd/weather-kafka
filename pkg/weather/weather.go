package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const urlTemplate = `http://api.openweathermap.org/data/2.5/weather?id=%d&appid=%s`

type WeatherEntry struct {
	Lat, Lon   float64
	ID         int
	Name       string
	Temp       float64
	TempMin    float64
	TempMax    float64
	Pressure   float64
	Humidity   float64
	Visibility float64
	WindSpeed  float64
	WindDeg    float64
	Clouds     float64
}

func GetWeather(apikey string, id int) (*WeatherEntry, error) {
	url := fmt.Sprintf(urlTemplate, id, apikey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Could not http-get station JSON: %v", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Could not read station JSON: %v", err)
	}
	var entry openWeatherEntry
	if err := json.Unmarshal(bytes, &entry); err != nil {
		return nil, fmt.Errorf("Could not parse JSON: %v", err)
	}
	we := WeatherEntry{
		Lat:        entry.Coord.Lat,
		Lon:        entry.Coord.Lon,
		ID:         entry.ID,
		Name:       entry.Name,
		Temp:       entry.Main.Temp,
		TempMin:    entry.Main.TempMin,
		TempMax:    entry.Main.TempMax,
		Pressure:   entry.Main.Pressure,
		Humidity:   entry.Main.Humidity,
		Visibility: entry.Visibility,
		WindSpeed:  entry.Wind.Speed,
		WindDeg:    entry.Wind.Deg,
		Clouds:     entry.Clouds.All,
	}
	return &we, nil
}

type openWeatherEntry struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility float64 `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}
