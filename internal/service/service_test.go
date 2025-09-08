package service

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aashi1008/weather-app/config"
	model "github.com/aashi1008/weather-app/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentWeatherResponse_Success(t *testing.T) {
	// Mock API server
	mockResponse := `{
		"latitude":50.5,
		"longitude":109.9,
		"timezone":"GMT",
		"current_units":{"time":"iso","temperature_2m":"C","wind_speed_10m":"km/h"},
		"current":{"time":"now","temperature_2m":20.0,"wind_speed_10m":5.0}
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockResponse)
	}))
	defer ts.Close()

	cfg := &config.AppConfig{
		HttpClient: &http.Client{},
		BaseURL:    ts.URL,
		Port:       "8080",
	}

	ws := NewWeatherService(cfg)

	req := model.Request{Lat: "50.5", Lon: "109.9"}
	resp, err := ws.GetCurrentWeatherResponse(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, 50.5, resp.Latitude)
	assert.Equal(t, 109.9, resp.Longitude)
	assert.Equal(t, "GMT", resp.Timezone)
	assert.Equal(t, 20.0, resp.Current.Temperature_2m)
}

func TestGetCurrentWeatherResponse_InvalidCoordinates(t *testing.T) {
	cfg := &config.AppConfig{}
	ws := NewWeatherService(cfg)

	req := model.Request{Lat: "abc", Lon: "109.9"}
	resp, err := ws.GetCurrentWeatherResponse(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid coordinates")
}
