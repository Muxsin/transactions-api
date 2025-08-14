package main

import (
	"context"
	"infotecs-transactions-api/internal/app"
	"infotecs-transactions-api/internal/config"
	"infotecs-transactions-api/internal/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	myApp := app.New(config.New(), db)

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
