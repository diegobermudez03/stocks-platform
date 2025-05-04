package main

import (
	"log"
	"os"

	"github.com/diegobermudez03/stocks-backend/internal/transport"
	"github.com/joho/godotenv"
)

const(
	PORT_VAR = "PORT"
)


func main() {
	/*
		Check if we have env variables loaded, if not, then try to load .env file (dev)
		If neither worked, cant start server
	*/
	if os.Getenv(PORT_VAR) == ""{
		if err := godotenv.Load(".env"); err != nil{
			log.Fatal("NO ENV VARIABLES LOADED, UNABLE TO START SERVER")
		}
	}
	config := transport.ServerConfig{
		Port: os.Getenv(PORT_VAR),
	}

	/*
		Create and run server
	*/
	server := transport.NewRestAPIServer(config)
	if err := server.Run(); err != nil{
		log.Fatal("Unable to start server: ", err.Error())
	}
}