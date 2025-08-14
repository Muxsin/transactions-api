package main

import (
	"infotecs-transactions-api/internal/app"
	"infotecs-transactions-api/internal/config"
	"infotecs-transactions-api/internal/database"

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
	if err := myApp.Run(); err != nil {
		panic(err)
	}
}
