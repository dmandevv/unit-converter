package handlers

import (
	"net/http"
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/length.html")
}

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/weight.html")
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/temperature.html")
}
