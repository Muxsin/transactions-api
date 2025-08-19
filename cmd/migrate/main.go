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
)

var migrations = map[string]any{
	"wallet":      &models.Wallet{},
	"transaction": &models.Transaction{},
}

func generateRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
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

	for name, model := range migrations {
		if db.Migrator().HasTable(model) {
			log.Printf("Table '%s' already exists, skipping migration.", name)
			continue
		}

		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate table '%s': %v", name, err)
		}

		// --- SEEDING LOGIC EXPLANATION ---
		// If you manually delete all data from the 'wallets' table
		// but the table itself still exists, this seeding code will NOT run again
		if name == "wallet" {
			log.Println("Seeding initial data...")

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
		}

		log.Printf("Table '%s' migrated successfully.", name)
	}

	log.Println("Migration process finished.")
}
