# Unit Converter 🔁

A small Go web app to convert units for:

Length: millimeter, centimeter, meter, kilometer, inch, foot, yard, mile.

Weight: milligram, gram, kilogram, ounce, pound.

Temperature: Celsius, Fahrenheit, Kelvin.

This project is an exercise from the roadmap.sh project: https://roadmap.sh/projects/unit-converter 📚

## Available pages
- `/length` — Length converter UI
- `/weight` — Weight converter UI
- `/temperature` — Temperature converter UI
- `/result` — Shows conversion result

## Project layout

- `main.go` — Server bootstrap and routes
- `handlers/` — Request handlers (`form.go`, `result.go`, `units.go`)
- `static/` — HTML templates and CSS (`length.html`, `weight.html`, `temperature.html`, `result.html`, `css/style.css`)

Note
- The server listens on port `:9090` by default.
