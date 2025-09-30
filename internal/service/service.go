package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aashi1008/weather-app/config"
	model "github.com/aashi1008/weather-app/internal/models"
	valid "github.com/aashi1008/weather-app/internal/validator"
)

type WeatherService interface {
	GetCurrentWeatherResponse(ctx context.Context, req model.Request) (*model.CurrentWeatherResponse, error)
}

type weatherService struct {
	cfg *config.AppConfig
}

// Constructor
func NewWeatherService(appcfg *config.AppConfig) WeatherService {
	return &weatherService{
		cfg: appcfg,
	}
}

func (ws *weatherService) GetCurrentWeatherResponse(ctx context.Context, req model.Request) (*model.CurrentWeatherResponse, error) {
	lat, lon := valid.GetCoordinates(req.Lat, req.Lon)

	url := fmt.Sprintf("%s?latitude=%.6f&longitude=%.6f&current=temperature_2m,wind_speed_10m", ws.cfg.BaseURL, lat, lon)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request:%s", err)
	}
	resp, err := ws.cfg.HttpClient.Do(request)
	if err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get response:%s", err)
	}

	defer resp.Body.Close()

	// res, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	// var apiresponse model.OpenMeteoResponse
	// err = json.Unmarshal(res, &apiresponse)
	// if err != nil {
	// 	return nil, err
	// }

	var apiresponse model.OpenMeteoResponse
	err = json.NewDecoder(resp.Body).Decode(&apiresponse)
	if err != nil {
		return nil, err
	}

	currentWeather := &model.CurrentWeatherResponse{
		Latitude:      apiresponse.Latitude,
		Longitude:     apiresponse.Longitude,
		Timezone:      apiresponse.Timezone,
		Current_units: model.WeatherCurrentUnits(apiresponse.Current_units),
		Current:       model.WeatherCurrent(apiresponse.Current),
	}

	return currentWeather, nil
}
