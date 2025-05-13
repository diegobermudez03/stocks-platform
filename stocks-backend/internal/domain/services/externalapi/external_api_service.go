package externalapi

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
)

type ExternalAPIServiceImpl struct {
	externalAPIUrl string
	externalAPIKey string
	websocketUrl string
	suscribedStocks map[string]float64
	readChannel chan string
	writeChannel chan domain.StockPriceUpdate
	closeChannel chan bool
}

func NewExternalAPIService(externalAPIUrl,externalAPIKey, websocketUrl string ) domain.ExternalApiService{
	return &ExternalAPIServiceImpl{
		externalAPIUrl: externalAPIUrl,
		externalAPIKey: externalAPIKey,
		websocketUrl: websocketUrl,
		suscribedStocks: map[string]float64{},
		closeChannel:  make(chan bool),
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
	Method to stablish websocket connection (is the one we can mock)
*/
/*func (s *ExternalAPIServiceImpl) StartLiveConnection() (chan string, chan domain.StockPriceUpdate, error){
	reader := make(chan string)
	writer := make(chan domain.StockPriceUpdate)
	var externalError error 
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(){
		url := s.websocketUrl + "?token=" + s.externalAPIKey
		connection, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil{
			externalError = err
		}
		wg.Done()
		defer connection.Close()
		defer close(reader)
		//run go routine for the reader
		go s.webSocketReader(connection, writer)

		//continue in this go routine as the writer (suscriber to new symbols)
		for message := range reader{
			suscribeMsg := InternalSuscribeDTO{
				Type: "suscribe",
				Symbol: message,
			}
			jsonBytes, _ := json.Marshal(suscribeMsg)
			//if error when trying to send, probably was already closed, so we end
			if err := connection.WriteMessage(websocket.TextMessage, jsonBytes); err!= nil{
				break
			}
		}
	}()
	wg.Wait()
	return reader, writer, externalError
}
*/

/*
	To close live websocket connection
*/
func (s *ExternalAPIServiceImpl) CloseLiveConnection(){
	//we send 2 times so that both goroutines get it
	s.closeChannel <- true 
	s.closeChannel <- true 
	s.suscribedStocks = map[string]float64{}
}


/*
	Websocket reader go routine, this go routine only focus on reading from the websocket and broadcasting
*/
/*func (s *ExternalAPIServiceImpl) webSocketReader(connection *websocket.Conn, channel chan domain.StockPriceUpdate){
	for{
		_, msgBytes, err := connection.ReadMessage()
		if err != nil{
			continue
		}
		priceUpdate := InternalStockPriceUpdate{}
		if err := json.Unmarshal(msgBytes, &priceUpdate); err != nil{
			continue 
		}
		for _, update := range priceUpdate.Data{
			channel <- domain.StockPriceUpdate{
				Symbol: update.S,
				Price: update.P,
			}
		}
	}
}
*/

////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////
//////////				MOCK WEBSOCKET LIVE IMPLEMENTATION 				/			///////////////

func (s *ExternalAPIServiceImpl) StartLiveConnection() (chan string, chan domain.StockPriceUpdate, error){
	reader := make(chan string)
	writer := make(chan domain.StockPriceUpdate)

	//start reader, in this case, it looks for the current real price of the symbol
	go func(){
		for{
			select{
			case stock :=<-reader:{
				if _, ok := s.suscribedStocks[stock]; !ok{
					url := s.externalAPIUrl + "/quote?symbol=" + stock  + "&token=" + s.externalAPIKey
					response, err := http.Get(url)
					if err != nil{
						continue
					}
					payload, err := io.ReadAll(response.Body)
					if err != nil{
						continue
					}
					price := InternalStockPrice{}
					if err := json.Unmarshal(payload, &price); err != nil{
						continue
					}
					s.suscribedStocks[stock] = price.C 
					response.Body.Close()
				}
			}
			case <- s.closeChannel:
				return 
			}
		}
	}()
	
	//start the writer, each 4 seconds will send a price uypdate to all 
	go func(){
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for{
			select {
			case <- ticker.C:{
				for stock, currentPrice := range s.suscribedStocks{
					randomChange := (rand.Float64()*20) - 10
					currentPrice = currentPrice + randomChange
					s.suscribedStocks[stock] = currentPrice
					writer <- domain.StockPriceUpdate{
						Symbol: stock,
						Price:  currentPrice,
					}
				}
			}
			case <- s.closeChannel:
				return 
			}
		}
	}()

	return reader, writer, nil
}