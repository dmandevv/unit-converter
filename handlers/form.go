package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

var meterConversions = map[string]float64{
	"millimeter": 1000.0,
	"centimeter": 100.0,
	"meter":      1.0,
	"kilometer":  0.001,
	"inch":       39.3701,
	"foot":       3.28084,
	"yard":       1.09361,
	"mile":       0.00062137273,
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.FormValue("length") != "" {
		inputString := r.FormValue("length")
		inputLength, err := strconv.ParseFloat(inputString, 64)
		if err != nil {
			http.Error(w, "Invalid length input", http.StatusBadRequest)
			return
		}
		unit_from := r.FormValue("unit_from")
		unit_to := r.FormValue("unit_to")

		fmt.Printf("Converting %f%s to %s", inputLength, unit_from, unit_to)

		// Convert input length to meters
		var lengthInMeters float64
		if val, ok := meterConversions[unit_from]; ok {
			lengthInMeters = inputLength / val
		} else {
			http.Error(w, "Invalid unit from", http.StatusBadRequest)
			return
		}

		// Convert meters to target unit
		var calculatedLength float64
		if val, ok := meterConversions[unit_to]; ok {
			calculatedLength = lengthInMeters * val
		} else {
			http.Error(w, "Invalid unit to", http.StatusBadRequest)
			return
		}

		result := fmt.Sprintf("%.3f %s = %.3f %s", inputLength, unit_from, calculatedLength, unit_to)
		http.Redirect(w, r, "/result?type=Result&value="+result, http.StatusSeeOther)
		return
	}
	http.Error(w, "No valid conversion data provided", http.StatusUnprocessableEntity)

}
