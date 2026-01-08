package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dmandevv/unit-converter/handlers"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", handlers.LengthHandler)
	mux.HandleFunc("/length", handlers.LengthHandler)
	mux.HandleFunc("/weight", handlers.WeightHandler)
	mux.HandleFunc("/temperature", handlers.TemperatureHandler)
	mux.HandleFunc("/form/length", handlers.FormLengthHandler)
	mux.HandleFunc("/form/weight", handlers.FormWeightHandler)
	mux.HandleFunc("/form/temperature", handlers.FormTemperatureHandler)
	mux.HandleFunc("/result", handlers.ResultHandler)

	srv := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	// start server
	go func() {
		log.Printf("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan
	log.Printf("received signal %v, shutting down", sig)
}
