package transport

import (
	"log"
	"net/http"
	"strconv"

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
	
	router.Get("/ratings", s.getRatings)
	router.Get("/actions", s.getActions)
	router.Get("/stocks", s.getStocks)
	log.Printf("Starting server at port: %s", s.config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), router)
}


/*
	Helper endpoint, to retrieve the possible ratings to use as filters
*/
func (s *RestAPIServer) getActions(w http.ResponseWriter, r *http.Request){
	actions, err := s.service.GetActions()
	if err != nil{
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, actions)
}


/*
	Helper endpoint, to retrieve the possible ratings to use as filters
*/
func (s *RestAPIServer) getRatings(w http.ResponseWriter, r *http.Request){
	ratings, err := s.service.GetRatings()
	if err != nil{
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, ratings)
}


/*
	Main endpoint, for retrieving the stocks, with all filters
*/
func (s *RestAPIServer) getStocks(w http.ResponseWriter, r *http.Request){
	filter := domain.GetStocksFilter{}
	page, _ := strconv.Atoi(r.URL.Query().Get("page")) 
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	targetStart, _ := strconv.ParseFloat(r.URL.Query().Get("target_start"), 64)
	targetEnd, _ := strconv.ParseFloat(r.URL.Query().Get("target_end"), 64)
	filter.Page = page
	filter.Size = size
	filter.TextSearch = r.URL.Query().Get("text_search")
	filter.RatingFromList = r.URL.Query()["rating_from"]
	filter.RatingToList = r.URL.Query()["rating_to"]
	filter.TimeStart = r.URL.Query().Get("time_start")
	filter.TimeEnd = r.URL.Query().Get("time_end")
	filter.Sort = r.URL.Query().Get("sort")
	filter.ActionList = r.URL.Query()["action"]
	filter.TargetStart = targetStart
	filter.TargetEnd = targetEnd

	stocks, err := s.service.GetStocks(filter)
	if err != nil{
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, stocks)
}