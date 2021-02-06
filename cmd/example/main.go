package main

import (
	"log"
	"net/http"
	"prometheus-custom-metrics-example/internal/example/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/liveness", handler.LivenessHandler)
	mux.HandleFunc("/readiness", handler.ReadinessHandler)
	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/metrics", handler.MetricsHandler)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
