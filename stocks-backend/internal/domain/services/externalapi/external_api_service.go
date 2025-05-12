package externalapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
	"github.com/gorilla/websocket"
)

type ExternalAPIServiceImpl struct {
	externalAPIUrl string
	externalAPIKey string
	websocketUrl string
	webSocketChannel chan string
	suscribers map[string] []chan domain.PriceUpdate
	suscribersLock sync.RWMutex
}

func NewExternalAPIService(externalAPIUrl,externalAPIKey, websocketUrl string ) domain.ExternalApiService{
	return &ExternalAPIServiceImpl{
		externalAPIUrl: externalAPIUrl,
		externalAPIKey: externalAPIKey,
		websocketUrl: websocketUrl,
		suscribersLock: sync.RWMutex{},
		suscribers: map[string] []chan domain.PriceUpdate{},
	}
}


/*
	Method to retrieve the company info of the stock
*/
func (s *ExternalAPIServiceImpl) GetCompanyProfile(symbol string)(*domain.CompanyProfileDTO, error){
	//get the company profile
	response, err := http.Get(s.externalAPIUrl + "/stock/profile2?symbol=" + symbol + "&token=" + s.externalAPIKey)
	if err != nil{
		return nil, domain.ErrInternalError
	}
	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil{
		return nil, domain.ErrInternalError
	}
	companyProfile := InternalCompanyProfileDTO{}
	if err := json.Unmarshal(payload, &companyProfile); err != nil{
		return nil, domain.ErrInternalError
	}
	companyProfileDTO :=  &domain.CompanyProfileDTO{
		Country: companyProfile.Country,
		Currency: companyProfile.Currency,
		Exchange: companyProfile.Exchange,
		Industry: companyProfile.FinnhubIndustry,
		Ipo: companyProfile.IPO,
		Logo: companyProfile.Logo,
		MarketCapital: companyProfile.MarketCapitalization,
		Name: companyProfile.Name,
		Phone: companyProfile.Phone,
		WebUrl: companyProfile.WebURL,
		ShareOutstanding: companyProfile.ShareOutstanding,
	}
	return companyProfileDTO, nil
}


/*
	Method to retrieve the latest news related with the stock symbol
*/
func (s *ExternalAPIServiceImpl) GetLatestNews(symbol string)([]domain.NewsDTO, error){
	currentDate := time.Now().Format("2006-01-02")
	response2, err := http.Get(s.externalAPIUrl + "/company-news?symbol=" + symbol + "&from=2025-01-01" + "&to=" +currentDate + "&token=" + s.externalAPIKey)
	if err != nil{
		return nil, err
	}
	defer response2.Body.Close()
	payload2, err := io.ReadAll(response2.Body)
	if err != nil{
		return nil, err
	}
	news := []InternalNewsDTO{}
	if err := json.Unmarshal(payload2, &news); err != nil{
		return nil, err
	}
	newsDTO := make([]domain.NewsDTO, len(news))
	for i, news := range news{
		newsDTO[i] = domain.NewsDTO{
			Date: time.Unix(news.Datetime, 0),
			Headline: news.Headline,
			Image: news.Image,
			Source: news.Source,
			Summary: news.Summary,
		}
	}
	return newsDTO, nil
}


/*
	Method to retrieve the latest stock sentiment of the symbol (list of latest months insider sentiment)
*/ 
func (s *ExternalAPIServiceImpl) GetStockSentiment(symbol string)(*domain.InternalSentimentDTO, error){
	currentDate := time.Now().Format("2006-01-02")
	response, err := http.Get(s.externalAPIUrl + "/stock/insider-sentiment?symbol=" + symbol + "&from=2024-01-01" + "&to=" +currentDate + "&token=" + s.externalAPIKey)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil{
		return nil, err
	}
	sentiment := domain.InternalSentimentDTO{}
	if err := json.Unmarshal(payload, &sentiment); err != nil{
		return nil, err
	}
	return &sentiment, nil
}


/*
	Method to suscribe to live connections of a stock
*/
func (s *ExternalAPIServiceImpl) LiveSymbolPrice(symbol string)(chan domain.PriceUpdate, error){
	//if the channel is nil, means that we dont have a websocket connection yet, so we need to create it
	if s.webSocketChannel == nil{
		channel, err := s.connectWebSocket()
		if err != nil{
			return nil, err
		}
		s.webSocketChannel = channel
	}
	s.webSocketChannel <- symbol
	suscriberChannel := make(chan domain.PriceUpdate)

	//ADD THE SUSCRIBER, CRITIC ZONE
	s.suscribersLock.Lock()
	slice, ok := s.suscribers[symbol]
	if !ok{
		s.suscribers[symbol] = []chan domain.PriceUpdate{}
		slice = s.suscribers[symbol]
	}
	s.suscribers[symbol] = append(slice, suscriberChannel)
	s.suscribersLock.Unlock()
	return suscriberChannel, nil
}


/*
	Internal  method to stablish webSocket connection and return the websocket from where it will listen to
*/
func (s *ExternalAPIServiceImpl) connectWebSocket() (chan string, error){
	channel := make(chan string)
	var externalError error
	wg := sync.WaitGroup{}
	wg.Add(1)

	/*
		Here we create the go routine that will handle the connection with the websocket channel
	*/
	go func(){
		url := s.websocketUrl + "?token=" + s.externalAPIKey
		connection, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil{
			externalError = err
		}
		wg.Done()
		defer connection.Close()
		defer close(channel)
		//run go routine for the reader
		go s.webSocketReader(connection)

		//continue in this go routine as the writer (suscriber to new symbols)
		for message := range channel{
			suscribeMsg := InternalSuscribeDTO{
				Type: "suscribe",
				Symbol: message,
			}
			jsonBytes, _ := json.Marshal(suscribeMsg)
			//if error when trying to send, probably was already closed, so we end
			if err := connection.WriteMessage(websocket.TextMessage, jsonBytes); err!= nil{
				s.webSocketChannel = nil
				break
			}
		}
	}()

	wg.Wait()
	return channel, externalError
}

/*
	Websocket reader go routine, this go routine only focus on reading from the websocket and broadcasting
*/
func (s *ExternalAPIServiceImpl) webSocketReader(connection *websocket.Conn){
	for{
		//always at the beginning check if we are empty of suscribers, if thats the case, we close the connection
		s.suscribersLock.RLock()
		if len(s.suscribers) == 0{
			//close connection
		}
		s.suscribersLock.RUnlock()

		_, msgBytes, err := connection.ReadMessage()
		log.Print(string(msgBytes))
		if err != nil{
			continue
		}
		priceUpdate := InternalStockPriceUpdate{}
		if err := json.Unmarshal(msgBytes, &priceUpdate); err != nil{
			continue 
		}

		//deliver update to the suscribers to the stock
		s.suscribersLock.RLock()
		for _, update := range priceUpdate.Data{
			suscribers, ok := s.suscribers[update.S]
			priceUpdate := domain.PriceUpdate{
				Price: update.P,
			}
			//if no suscribers, ommit
			if !ok{
				continue 
			}
			//send update to all channels
			for _, suscriber := range suscribers{
				suscriber <-priceUpdate
			}
		}
		s.suscribersLock.RUnlock()
	}
}


/*
	To unsuscribe client
*/
func(s *ExternalAPIServiceImpl) UnsuscribePriceClient(symbol string, channel chan domain.PriceUpdate){
	s.suscribersLock.Lock()
	defer s.suscribersLock.Unlock()
	slice, ok := s.suscribers[symbol]
	if !ok{
		return 
	}
	var index int  = -1
	for i, suscriber := range slice{
		if suscriber == channel{
			index = i 
			break
		}
	}
	if index == -1{
		return 
	}
	slice = append(slice[:index], slice[index+1:]...)
	//if there are no suscribers, then we remove the whole key
	if len(slice) == 0{
		delete(s.suscribers, symbol)
	}else{
		s.suscribers[symbol] = slice
	}
}