package models

type Request struct {
	Lat string `json:"lat,omitempty"`
	Lon string `json:"lon,omitempty"`
}

type OpenMeteoResponse struct {
	Latitude      float64      `json:"latitude,omitempty"`
	Longitude     float64      `json:"longitude,omitempty"`
	Timezone      string       `json:"timezone,omitempty"`
	Current_units CurrentUnits `json:"current_units,omitempty"`
	Current       Current      `json:"current,omitempty"`
}

type CurrentUnits struct {
	Time           string `json:"time,omitempty"`
	Temperature_2m string `json:"temperature_2m,omitempty"`
	Wind_speed_10m string `json:"wind_speed_10m,omitempty"`
}

type Current struct {
	Time           string  `json:"time,omitempty"`
	Temperature_2m float64 `json:"temperature_2m,omitempty"`
	Wind_speed_10m float64 `json:"wind_speed_10m,omitempty"`
}

type CurrentWeatherResponse struct {
	Latitude      float64             `json:"latitude,omitempty"`
	Longitude     float64             `json:"longitude,omitempty"`
	Timezone      string              `json:"timezone,omitempty"`
	Current_units WeatherCurrentUnits `json:"currentWeatherUnits,omitempty"`
	Current       WeatherCurrent      `json:"currentWeather,omitempty"`
}

type WeatherCurrentUnits struct {
	Time           string `json:"time,omitempty"`
	Temperature_2m string `json:"temperature,omitempty"`
	Wind_speed_10m string `json:"wind_speed,omitempty"`
}

type WeatherCurrent struct {
	Time           string  `json:"time,omitempty"`
	Temperature_2m float64 `json:"temperature,omitempty"`
	Wind_speed_10m float64 `json:"wind_speed,omitempty"`
}
