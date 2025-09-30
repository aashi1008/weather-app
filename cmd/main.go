package main

import (
	"log"
	"net/http"

	"github.com/aashi1008/weather-app/config"
	"github.com/aashi1008/weather-app/internal/handler"
	"github.com/aashi1008/weather-app/internal/metrics"
	"github.com/aashi1008/weather-app/internal/routes"
	"github.com/aashi1008/weather-app/internal/service"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {

	p := prometheus.NewRegistry()
	m := metrics.NewMetrics(p)

	httpClient := &http.Client{}

	appConfig := config.NewAppConfig(httpClient)
	weatherService := service.NewWeatherService(appConfig)
	weatherHandler := handler.NewWeatherHandler(weatherService, m)

	routes.SetupRoutes(weatherHandler, m)
	log.Print("Starting server at port:", appConfig.Port)
	if err := http.ListenAndServe(":"+appConfig.Port, nil); err != nil {
		log.Fatal(err)
	}
}
