package main

import (
	"log"
	"os"
	"strconv"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain/services/externalapi"
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain/services/stocks"
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/repository"
	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/transport"
	"github.com/joho/godotenv"
)

const(
	PORT_VAR = "PORT"
	API_URL = "API_URL"
	API_TOKEN = "API_TOKEN"
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_USER = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD" 
	DB_DBNAME = "DB_DBNAME"
	DB_SSL_MODE = "DB_SSL_MODE"
	EXTERNAL_API_URL = "EXTERNAL_API_URL"
	EXTERNAL_API_KEY = "EXTERNAL_API_KEY"
	WEBSOCKET_API_URL = "WEBSOCKET_API_URL"
	ALL_OR_NOTHING = "ALL_OR_NOTHING"
)


func main() {
	/*
		Check if we have env variables loaded, if not, then try to load .env file (dev)
		If neither worked, cant start server
	*/
	if os.Getenv(API_TOKEN) == ""{
		if err := godotenv.Load(".env"); err != nil{
			log.Fatal("NO ENV VARIABLES LOADED, UNABLE TO START SERVER")
		}
	}
	config := transport.ServerConfig{
		Port: os.Getenv(PORT_VAR),
	}
	dbConfig := repository.DBConfig{
		Host: os.Getenv(DB_HOST),
		Port: os.Getenv(DB_PORT),
		User: os.Getenv(DB_USER),
		Password: os.Getenv(DB_PASSWORD),
		DbName: os.Getenv(DB_DBNAME),
		SSLMode: os.Getenv(DB_SSL_MODE),
	}
	allOrNothing, _ := strconv.Atoi(os.Getenv(ALL_OR_NOTHING))
	/*
		Dependency injection
	*/
	db, err := repository.InitDb(dbConfig)
	if err != nil{
		log.Fatalf("Unable to connect to db: %s", err.Error())
	}
	repo := repository.NewStocksPostgresRepo(db)
	externalAPIService := externalapi.NewExternalAPIService(os.Getenv(EXTERNAL_API_URL), os.Getenv(EXTERNAL_API_KEY), os.Getenv(WEBSOCKET_API_URL))
	stocksService := stocks.NewStocksService(repo, os.Getenv(API_URL), os.Getenv(API_TOKEN), externalAPIService)
	if err := stocksService.PopulateDatabase(allOrNothing); err != nil{
		log.Fatalf("Unable to populate db with API data: %s", err.Error())
	}
	log.Print("Succesfully Populated the db")
	server := transport.NewRestAPIServer(config, stocksService)
	/*
		Run server
	*/
	if err := server.Run(); err != nil{
		log.Fatal("Unable to start server: ", err.Error())
	}
}