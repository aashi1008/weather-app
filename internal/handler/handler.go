package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aashi1008/weather-app/internal/metrics"
	model "github.com/aashi1008/weather-app/internal/models"
	"github.com/aashi1008/weather-app/internal/service"
	valid "github.com/aashi1008/weather-app/internal/validator"
)

type WeatherHandler struct {
	svc service.WeatherService
	m *metrics.Metrics
}

func NewWeatherHandler(svc service.WeatherService, m *metrics.Metrics) *WeatherHandler {
	return &WeatherHandler{
		svc: svc,
		m : m,
	}
}

func (h *WeatherHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *WeatherHandler) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	var req model.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// req.Lat = r.URL.Query().Get("lat")
	// req.Lon = r.URL.Query().Get("lon")
	// if req.Lat == "" || req.Lon == "" {
	// 	http.Error(w, "Missing query params 'lat' or 'lon'", http.StatusBadRequest)
	// 	return
	// }

	err = valid.ValidateCoordinates(req.Lat, req.Lon)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid coordinates")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	apiresponse, err := h.svc.GetCurrentWeatherResponse(ctx, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}

	h.m.WeatherHttpRequests.WithLabelValues().Inc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiresponse)

}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}