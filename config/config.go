/*
This file is designed to provide a database connection string for PostreSQL database,
with configuration values from .env file.
*/
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ No .env file found, using system environment variables.")
	} else {
		fmt.Println("✅ .env file loaded successfully!")
	}

	// 🔍 Debug: Check if JWT_SECRET is being loaded
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("❌ ERROR: JWT_SECRET is not set!")
	} else {
		fmt.Println("🔑 JWT_SECRET Loaded:", secret)
	}
}

func GetDatabaseDSN() string {
	LoadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
}
