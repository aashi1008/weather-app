package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aashi1008/weather-app/config"
	"github.com/aashi1008/weather-app/internal/handler"
	"github.com/aashi1008/weather-app/internal/service"
)

func main() {

	httpClient := &http.Client{}

	appConfig := config.NewAppConfig(httpClient)
	weatherService := service.NewWeatherService(appConfig)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Connected to server"))
	})
	http.HandleFunc("/weather", weatherHandler.GetWeatherHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
