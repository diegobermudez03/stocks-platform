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
	router.Get("/stocks/{id}/live", s.LivePriceUpdates)

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
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", s.config.Port), r)
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


/*
 	Server sent events for price live updates
*/
func (s *RestAPIServer) LivePriceUpdates(w http.ResponseWriter, r *http.Request){
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
	channel, err := s.service.SuscribeStockPrice(uuidId)

	// SETTING SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("data: CONNECTED\n\n"))
	flusher.Flush()

	//channel to detect connection closing
	clientGone := r.Context().Done()
	select{
	case <- clientGone:
		s.service.UnsuscribeFromStock(uuidId, channel)
		return 
	case update, ok :=<-channel:{
		//if for any reason the channel was closed, then we close the SSE connection
		if !ok{
			w.Write([]byte("event:close\ndata: Connection closing\n\n"))
			flusher.Flush()
			return 
		}
		//if it was an update, we send the update message
		message := fmt.Sprintf("data: %v\n\n", update.Price)
		w.Write([]byte(message))
		flusher.Flush()
	}
	}
}