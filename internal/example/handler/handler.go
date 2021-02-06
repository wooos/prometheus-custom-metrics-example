package handler

import (
	"fmt"
	"net/http"
	"prometheus-custom-metrics-example/internal/counter"
)

var c = counter.Counter{0}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("OnlineUser %d", c.Get())
	w.Write([]byte(msg))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	c.Incr()
	msg := fmt.Sprintf("Hello, you are number %d", c.Get())
	w.Write([]byte(msg))
}

func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("liveness"))
}

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("readiness"))
}
