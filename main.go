package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dmandevv/unit-converter/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.LengthHandler)
	mux.HandleFunc("/length", handlers.LengthHandler)
	mux.HandleFunc("/weight", handlers.WeightHandler)
	mux.HandleFunc("/temperature", handlers.TemperatureHandler)
	mux.HandleFunc("/form", handlers.FormHandler)
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
