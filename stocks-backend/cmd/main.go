package main

import (
	"log"
	"os"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain/service"
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
	/*
		Dependency injection
	*/
	db, err := repository.InitDb(dbConfig)
	if err != nil{
		log.Fatalf("Unable to connect to db: %s", err.Error())
	}
	repo := repository.NewStocksPostgresRepo(db)
	service := service.NewStocksService(repo, os.Getenv(API_URL), os.Getenv(API_TOKEN), os.Getenv(EXTERNAL_API_URL), os.Getenv(EXTERNAL_API_KEY))
	if err := service.PopulateDatabase(); err != nil{
		log.Fatalf("Unable to populate db with API data: %s", err.Error())
	}
	log.Print("Succesfully Populated the db")
	server := transport.NewRestAPIServer(config, service)
	/*
		Run server
	*/
	if err := server.Run(); err != nil{
		log.Fatal("Unable to start server: ", err.Error())
	}
}