package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	WeatherHttpRequests *prometheus.CounterVec
	Registry                   *prometheus.Registry
}

func NewMetrics(p *prometheus.Registry) *Metrics {
	m := &Metrics{
		WeatherHttpRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "TotalHttpRequestsForWeather",
			Help: "Total no of http requests for Weather",
		}, []string{}),
		Registry: p,
	}
	p.MustRegister(m.WeatherHttpRequests)
	return m
}
