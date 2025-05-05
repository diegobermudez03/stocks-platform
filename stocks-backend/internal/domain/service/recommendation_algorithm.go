package service

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/diegobermudez03/stocks-platform/stocks-backend/internal/domain"
)


var ratingScoreMap = map[string]int{
	"Positive" : 7,
	"Sector Perform" : 3,
	"Speculative Buy" : 8,
	"Hold" : 5,
	"Negative" : 2,
	"Overweight" : 6,
	"Cautious" : 4,
	"Sector Underperform" : 3,
	"Sell" : 1,
	"Unchanged" : 5,
	"Underperform" : 3,
	"Market Outperform" :7,
	"Market Perform" : 5,
	"Outperform" : 7,
	"In-Line": 6,
	"Sector Outperform" : 7,
	"Reduce" : 2,
	"Buy" : 9,
	"Under Perform" : 3,
	"Sector Weight" : 5,
	"Equal Weight" : 6,
	"Strong-Buy": 10,
	"Underweight" : 3,
	"Peer Perform" : 6,
	"Neutral" : 5,
}

var actionsScores = map[string]int{
	"downgraded by" : -8,
	"initiated by" : 3,
	"target set by" : 1,
	"target raised by" : 2,
	"upgraded by" : 8,
	"reiterated by" : 4,
	"target lowered by" : -3,

}

/*
	Method to get the recommendations to investi today
*/
func (s *StocksServiceImpl) GetRecommendations()([]domain.RecommendationDTO, error){
	topStocks, err := s.repo.GetBasicTopStocks()
	if err != nil{
		return nil, err 
	}
	recommendations := make([]domain.StockWithScore, len(topStocks))
	for i, stock := range topStocks{
		//calculate the recommendation score
		sentimentScore := s.getAvarageSentiment(stock.Ticker)
		variationPercentage := (stock.TargetTo-stock.TargetFrom)/stock.TargetFrom*100
		ratingVariationScore := ratingScoreMap[stock.RatingTo] - ratingScoreMap[stock.RatingFrom]
		score := 0.35 * (sentimentScore/100)
		score += 0.25 * (variationPercentage/100)
		score += 0.20 * (float64(ratingVariationScore)/100)
		score += 0.15 * (float64(actionsScores[stock.Action])/8)
		daysSinceNow := time.Since(stock.Time).Hours()/24
		//this is to punish recommendations based on when they were made, because if they were made a month ago, is not as good as one made 1 day ago
		recencyWeight := 1 / (1+(daysSinceNow / 100))
		score *= recencyWeight
		recommendations[i]= domain.StockWithScore{
			Score: score,
			AvrgSentiment: sentimentScore,
			Stock: stock,
		}
	}
	sort.Slice(recommendations, func(i, j int)bool{
		return recommendations[i].Score > recommendations[j].Score
	})
	recommendationsDTOs := make([]domain.RecommendationDTO, 10)
	for i := 0; i < 10; i++{
		stock:= recommendations[i].Stock
		recommendationsDTOs[i] = domain.RecommendationDTO{
			RecommendationScore: recommendations[i].Score,
			AvrgSentiment: recommendations[i].AvrgSentiment,
			StockDTO: *s.stockModelToDTO(&stock),
		}
	}
	return recommendationsDTOs, nil
}


/*
	Method to get the avarage sentiment of the given symbol
*/
func (s *StocksServiceImpl) getAvarageSentiment(ticker string) float64{
	currentDate := time.Now().Format("2006-01-02")
	response, err := http.Get(s.externalAPIUrl + "/stock/insider-sentiment?symbol=" + ticker + "&from=2024-01-01" + "&to=" +currentDate + "&token=" + s.externalAPIKey)
	if err != nil{
		return 0
	}
	defer response.Body.Close()
	payload, err := io.ReadAll(response.Body)
	if err != nil{
		return 0
	}
	sentiment := InternalSentimentDTO{}
	if err := json.Unmarshal(payload, &sentiment); err != nil{
		return 0
	}
	//read sentiment data
	var avrgSentiment float64 
	for _, sent := range sentiment.Data{
		avrgSentiment += sent.Mspr
	}
	if len(sentiment.Data) > 0 {
		avrgSentiment = avrgSentiment / float64(len(sentiment.Data))
	}
	return avrgSentiment
}