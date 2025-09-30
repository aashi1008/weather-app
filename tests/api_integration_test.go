//go:build integration
// +build integration

package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/aashi1008/weather-app/config"
	"github.com/aashi1008/weather-app/internal/handler"
	"github.com/aashi1008/weather-app/internal/metrics"
	model "github.com/aashi1008/weather-app/internal/models"
	"github.com/aashi1008/weather-app/internal/service"
)

func TestWeatherApp_Integration(t *testing.T) {
	// Fake external API (OpenMeteo)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"latitude":50.5,
			"longitude":109.9,
			"timezone":"GMT",
			"current_units":{"time":"iso","temperature_2m":"C","wind_speed_10m":"km/h"},
			"current":{"time":"now","temperature_2m":22.0,"wind_speed_10m":4.0}
		}`)
	}))
	defer ts.Close()

	// Override env var for BaseURL
	os.Setenv("WEATHER_API_URL", ts.URL)
	defer os.Unsetenv("WEATHER_API_URL")

	// Wire up app (real config, service, handler)
	appCfg := config.NewAppConfig(&http.Client{})
	ws := service.NewWeatherService(appCfg)
	wh := handler.NewWeatherHandler(ws, nil)

	// Register handlers globally (like in main.go)
	http.HandleFunc("/weather", wh.GetWeatherHandler)

	// Run app server with DefaultServeMux
	appSrv := httptest.NewServer(http.DefaultServeMux)
	defer appSrv.Close()

	// Send real request
	reqBody := `{"lat":"50.5","lon":"109.9"}`
	resp, err := http.Post(appSrv.URL+"/weather", "application/json", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", resp.StatusCode)
	}

	var got model.CurrentWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Validate response content
	if got.Latitude != 50.5 {
		t.Errorf("expected latitude 50.5, got %v", got.Latitude)
	}
	if got.Longitude != 109.9 {
		t.Errorf("expected longitude 109.9, got %v", got.Longitude)
	}
	if got.Timezone != "GMT" {
		t.Errorf("expected timezone GMT, got %s", got.Timezone)
	}
	if got.Current.Temperature_2m != 22.0 {
		t.Errorf("expected temperature 22.0, got %v", got.Current.Temperature_2m)
	}
	if got.Current.Wind_speed_10m != 4.0 {
		t.Errorf("expected wind_speed 4.0, got %v", got.Current.Wind_speed_10m)
	}
}
