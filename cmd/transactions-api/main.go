package main

import (
	"context"
	"infotecs-transactions-api/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	myApp, err := app.New()
	if err != nil {
		panic(err)
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := myApp.Run(); err != nil {
			panic(err)
		}
	}()

	sig := <-stopCh
	log.Printf("Received signal: %v. Starting graceful shutdown...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := myApp.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
	}

	log.Println("Application shutdown gracefully. Exiting.")
}
