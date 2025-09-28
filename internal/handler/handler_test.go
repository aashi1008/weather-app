package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aashi1008/weather-app/config"
	model "github.com/aashi1008/weather-app/internal/models"
	"github.com/aashi1008/weather-app/internal/service"
	"github.com/stretchr/testify/assert"
)

func setupHandler(mockResponse string) *WeatherHandler {
	// Spin up a fake API server that WeatherService will call
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockResponse)
	}))

	cfg := &config.AppConfig{
		HttpClient: &http.Client{},
		BaseURL:    ts.URL, // point to fake server
		Port:       "8080",
	}

	ws := service.NewWeatherService(cfg)
	return NewWeatherHandler(ws, nil)
}

func TestGetWeatherHandler_Success(t *testing.T) {
	mockResponse := `{
		"latitude":50.5,
		"longitude":109.9,
		"timezone":"GMT",
		"current_units":{"time":"iso","temperature_2m":"C","wind_speed_10m":"km/h"},
		"current":{"time":"now","temperature_2m":20.0,"wind_speed_10m":5.0}
	}`

	h := setupHandler(mockResponse)

	reqBody := `{"lat":"50.5","lon":"109.9"}`
	req := httptest.NewRequest(http.MethodPost, "/weather", bytes.NewBufferString(reqBody))
	w := httptest.NewRecorder()

	h.GetWeatherHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got model.CurrentWeatherResponse
	err := json.NewDecoder(w.Body).Decode(&got)
	assert.NoError(t, err)
	assert.Equal(t, 50.5, got.Latitude)
	assert.Equal(t, 109.9, got.Longitude)
	assert.Equal(t, "GMT", got.Timezone)
	assert.Equal(t, 20.0, got.Current.Temperature_2m)
}

func TestGetWeatherHandler_InvalidBody(t *testing.T) {
	h := setupHandler(`{}`)

	req := httptest.NewRequest(http.MethodPost, "/weather", bytes.NewBufferString("invalid-json"))
	w := httptest.NewRecorder()

	h.GetWeatherHandler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid request body")
}

func TestGetWeatherHandler_InvalidCoordinates(t *testing.T) {
	h := setupHandler(`{}`)

	reqBody := `{"lat":"abc","lon":"109.9"}`
	req := httptest.NewRequest(http.MethodPost, "/weather", bytes.NewBufferString(reqBody))
	w := httptest.NewRecorder()

	h.GetWeatherHandler(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "invalid coordinates")
}
