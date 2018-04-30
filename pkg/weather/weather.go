package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const urlTemplate = `http://api.openweathermap.org/data/2.5/weather?id=%d&appid=%s&units=metric`

type WeatherEntry struct {
	IterationTime time.Time `json:"iterationTime"`
	EntryTime     time.Time `json:"time"`
	Lat           float64   `json:"lat"`
	Lon           float64   `json:"lon"`
	ID            int       `json:"id"`
	Desc          string    `json:"description"`
	Name          string    `json:"name"`
	Temp          float64   `json:"temp"`
	TempMin       float64   `json:"tempMin"`
	TempMax       float64   `json:"tempMax"`
	Pressure      float64   `json:"pressure"`
	Humidity      float64   `json:"humidity"`
	Visibility    float64   `json:"visibility"`
	WindSpeed     float64   `json:"windSpeed"`
	WindDeg       float64   `json:"windDeg"`
	Clouds        float64   `json:"clouds"`
}

func GetWeather(apikey string, t time.Time, id int) (*WeatherEntry, error) {
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
		IterationTime: t,
		EntryTime:     time.Now(),
		Lat:           entry.Coord.Lat,
		Lon:           entry.Coord.Lon,
		ID:            entry.ID,
		Name:          entry.Name,
		Temp:          entry.Main.Temp,
		TempMin:       entry.Main.TempMin,
		TempMax:       entry.Main.TempMax,
		Pressure:      entry.Main.Pressure,
		Humidity:      entry.Main.Humidity,
		Visibility:    entry.Visibility,
		WindSpeed:     entry.Wind.Speed,
		WindDeg:       entry.Wind.Deg,
		Clouds:        entry.Clouds.All,
	}
	if len(entry.Weather) > 0 {
		we.Desc = entry.Weather[0].Description
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
