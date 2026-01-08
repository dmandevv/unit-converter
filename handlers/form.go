package handlers

import (
	"fmt"
	"math"
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

var kilogramConversions = map[string]float64{
	"milligram": 1000000.0,
	"gram":      1000.0,
	"kilogram":  1.0,
	"ounce":     35.274,
	"pound":     2.20462,
}

var absoluteZero_Celsius = -273.15
var absoluteZero_Fahrenheit = -459.67

func FormLengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	inputString := r.FormValue("length")
	inputLength, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		inputLength = 1
	}

	if inputLength < 0 {
		http.Redirect(w, r, "/result?type=Result&value="+"INVALID INPUT LENGTH", http.StatusSeeOther)
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

	result := fmt.Sprintf("%g %s = %g %s", inputLength, unit_from, calculatedLength, unit_to)
	http.Redirect(w, r, "/result?type=Result&value="+result, http.StatusSeeOther)
}

func FormWeightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	inputString := r.FormValue("weight")
	inputWeight, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		inputWeight = 1
	}

	if inputWeight < 0 {
		http.Redirect(w, r, "/result?type=Result&value="+"INVALID INPUT WEIGHT", http.StatusSeeOther)
		return
	}

	unit_from := r.FormValue("unit_from")
	unit_to := r.FormValue("unit_to")

	fmt.Printf("Converting %f%s to %s", inputWeight, unit_from, unit_to)

	// Convert input weight to kilograms
	var weightInKilograms float64
	if val, ok := kilogramConversions[unit_from]; ok {
		weightInKilograms = inputWeight / val
	} else {
		http.Error(w, "Invalid unit from", http.StatusBadRequest)
		return
	}

	// Convert kilograms to target unit
	var calculatedWeight float64
	if val, ok := kilogramConversions[unit_to]; ok {
		calculatedWeight = weightInKilograms * val
	} else {
		http.Error(w, "Invalid unit to", http.StatusBadRequest)
		return
	}

	result := fmt.Sprintf("%g %s = %g %s", inputWeight, unit_from, calculatedWeight, unit_to)
	http.Redirect(w, r, "/result?type=Result&value="+result, http.StatusSeeOther)
}

func FormTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	inputString := r.FormValue("temperature")
	inputTemp, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		inputTemp = 0
	}
	unit_from := r.FormValue("unit_from")
	unit_to := r.FormValue("unit_to")

	fmt.Printf("Converting %f%s to %s", inputTemp, unit_from, unit_to)

	//Check if input temperature is above absolute zero
	switch unit_from {
	case "celsius":
		if inputTemp < absoluteZero_Celsius {
			http.Redirect(w, r, "/result?type=Result&value="+"INVALID INPUT TEMPERATURE", http.StatusSeeOther)
			return
		}
	case "fahrenheit":
		if inputTemp < absoluteZero_Fahrenheit {
			http.Redirect(w, r, "/result?type=Result&value="+"INVALID INPUT TEMPERATURE", http.StatusSeeOther)
			return
		}
	case "kelvin":
		if inputTemp < 0 {
			http.Redirect(w, r, "/result?type=Result&value="+"INVALID INPUT TEMPERATURE", http.StatusSeeOther)
			return
		}
	}

	var calculatedTemp float64 = math.NaN()
	if unit_from == unit_to {
		calculatedTemp = inputTemp
	} else {
		switch unit_from {
		case "celsius":
			switch unit_to {
			case "fahrenheit":
				calculatedTemp = (inputTemp * 9 / 5) + 32
			case "kelvin":
				calculatedTemp = inputTemp - absoluteZero_Celsius
			}
		case "fahrenheit":
			switch unit_to {
			case "celsius":
				calculatedTemp = (inputTemp - 32) * 5 / 9
			case "kelvin":
				calculatedTemp = (inputTemp - absoluteZero_Fahrenheit) * 5 / 9
			}
		case "kelvin":
			switch unit_to {
			case "celsius":
				calculatedTemp = inputTemp + absoluteZero_Celsius
			case "fahrenheit":
				calculatedTemp = (inputTemp+absoluteZero_Celsius)*9/5 + 32
			}
		}
	}
	result := fmt.Sprintf("%g %s = %g %s", inputTemp, unit_from, calculatedTemp, unit_to)
	http.Redirect(w, r, "/result?type=Result&value="+result, http.StatusSeeOther)
}
