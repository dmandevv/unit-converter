package handlers

import (
	"html/template"
	"net/http"
)

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/result.html")
	if err != nil {
		http.Error(w, "template parse error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Type  string
		Value string
	}{
		Type:  r.URL.Query().Get("type"),
		Value: r.URL.Query().Get("value"),
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "template execute error", http.StatusInternalServerError)
		return
	}
}
