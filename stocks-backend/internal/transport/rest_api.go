package transport

import (
	"log"
	"net/http"

	"fmt"

	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	Port string
}

type RestAPIServer struct {
	config ServerConfig
}

func NewRestAPIServer(config ServerConfig) *RestAPIServer {
	return &RestAPIServer{
		config: config,
	}
}

func (s *RestAPIServer) Run() error{
	router := chi.NewRouter()
	
	log.Printf("Starting server at port: %s", s.config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), router)
}