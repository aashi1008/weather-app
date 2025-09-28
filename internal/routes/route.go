package routes

import (
	"net/http"

	"github.com/aashi1008/weather-app/internal/handler"
	"github.com/aashi1008/weather-app/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(h *handler.WeatherHandler, m *metrics.Metrics) {
	http.Handle("/metrics", promhttp.HandlerFor(m.Registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/health", h.GetHealth)
	http.HandleFunc("/weather", h.GetWeatherHandler)
}
