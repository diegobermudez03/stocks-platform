package transport

import (
	"log"
	"net/http"

	"fmt"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	Port string
}

type RestAPIServer struct {
	config ServerConfig
	service domain.StocksService
}

func NewRestAPIServer(config ServerConfig, service domain.StocksService) *RestAPIServer {
	return &RestAPIServer{
		config: config,
		service: service,
	}
}

func (s *RestAPIServer) Run() error{
	router := chi.NewRouter()
	
	log.Printf("Starting server at port: %s", s.config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), router)
}