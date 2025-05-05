package transport

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"fmt"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
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
	/*
		Suscribe routes handlers
	*/
	router.Get("/ratings", s.getRatings)
	router.Get("/actions", s.getActions)
	router.Get("/stocks", s.getStocks)
	router.Get("/stocks/{id}", s.getStockData)
	router.Get("/recommendations", s.GetRecommendations)

	log.Printf("Starting server at port: %s", s.config.Port)
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, 
    }))
	r.Mount("/api/v1", router)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), r)
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

/*
	Endpoint which retrieves the full information for a stock, including company profile and news
*/
func (s *RestAPIServer) getStockData(w http.ResponseWriter, r *http.Request){
	stockId := chi.URLParam(r, "id")
	if stockId == ""{
		WriteError(w, http.StatusBadRequest, errors.New("invalid stock id"))
		return
	}
	uuidId, err := uuid.Parse(stockId)
	if err != nil{
		WriteError(w, http.StatusBadRequest, errors.New("invalid stock id"))
		return
	}
	info, err := s.service.GetStockFullData(uuidId)
	if err != nil{
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, info)
}


/*
	Endpoint to retrieve the recommendations
*/
func (s *RestAPIServer) GetRecommendations(w http.ResponseWriter, r *http.Request){
	recommendations, err := s.service.GetRecommendations()
	if err != nil{
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, recommendations)
}