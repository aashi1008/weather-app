package main

import (
	"log"
	"net/http"

	"github.com/aashi1008/weather-app/config"
	"github.com/aashi1008/weather-app/internal/handler"
	"github.com/aashi1008/weather-app/internal/service"
)

func main() {

	httpClient := &http.Client{}

	appConfig := config.NewAppConfig(httpClient)
	weatherService := service.NewWeatherService(appConfig)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	http.HandleFunc("/weather", weatherHandler.GetWeatherHandler)

	if err := http.ListenAndServe(":"+appConfig.Port, nil); err != nil {
		log.Fatal(err)
	}
}
