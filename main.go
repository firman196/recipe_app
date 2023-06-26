package main

import (
	database "Recipe_App/database/mysql"
	"Recipe_App/models"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal().Err(envErr).Msg("cannot load environment")
	}

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080"
	}

	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_DATABASE")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	db, dbErr := database.ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("cannot connect to database")
	}

	//auto migrate database
	db.AutoMigrate(&models.Bahan{})
	db.AutoMigrate(&models.Kategori{})
	db.AutoMigrate(&models.Resep{})
	db.AutoMigrate(&models.Komposisi{})
}
