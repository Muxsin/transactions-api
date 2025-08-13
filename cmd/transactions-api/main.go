package main

import (
	"infotecs-transactions-api/internal/app"
	"infotecs-transactions-api/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	myApp := app.New(config.New())
	if err := myApp.Run(); err != nil {
		panic(err)
	}
}
