package main

import (
	"context"
	"golumbus/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "book-api", log.LstdFlags)

	// create API handlers
	ph := handlers.NewProducts(logger)

	// create new servemux and register the handlers
	servemux := http.NewServeMux()
	servemux.Handle("/", ph)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      servemux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Graceful Shutdown
	signalChann := make(chan os.Signal, 1)
	signal.Notify(signalChann, os.Interrupt)
	signal.Notify(signalChann, syscall.SIGTERM)
	sig := <-signalChann
	logger.Println("Received termination signal, graceful shutdown", sig)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
