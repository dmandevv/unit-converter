package handlers

import (
	"html/template"
	"net/http"
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/length.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/weight.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/temperature.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
