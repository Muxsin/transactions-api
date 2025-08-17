package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"infotecs-transactions-api/internal/database"
	"infotecs-transactions-api/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func generateRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func truncateWallets(db *gorm.DB) error {
	log.Println("Truncating wallets table...")

	return db.Exec("TRUNCATE TABLE wallets RESTART IDENTITY CASCADE").Error
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	db, err := database.Connect(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Wallet{}, &models.Transaction{}); err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get underlying sql.DB: %v", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}()

	if err := truncateWallets(db); err != nil {
		log.Fatalf("Failed to truncate wallets table: %v", err)
	}

	const numWallets = 10
	for i := 0; i < numWallets; i++ {
		address, err := generateRandomAddress()
		if err != nil {
			log.Fatalf("failed to generate address: %v", err)
		}

		newWallet := models.Wallet{
			Address: address,
			Balance: 10000,
		}

		if err := db.Create(&newWallet).Error; err != nil {
			log.Println("can't create a wallet: ", newWallet)
		}

		fmt.Println(newWallet.Address)
	}

	log.Println("migrate the schemas finished")
}
