package stocks

import (
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
func (s *StocksServiceImpl) getRecommendationsInternal()([]domain.RecommendationDTO, error){
	topStocks, err := s.repo.GetBasicTopStocks()
	if err != nil{
		return nil, err 
	}
	recommendations := make([]domain.StockWithScore, len(topStocks))
	for i, stock := range topStocks{
		sentimentScore := s.getAvarageSentiment(stock.Ticker)
		score := s.getRecommendationScore(stock, sentimentScore)
		recommendations[i]= domain.StockWithScore{
			Score: score,
			AvrgSentiment: sentimentScore,
			Stock: stock,
		}
	}
	sort.Slice(recommendations, func(i, j int)bool{
		return recommendations[i].Score > recommendations[j].Score
	})
	lenght := 10 
	if lenght > len(recommendations){
		lenght = len(recommendations)
	}
	recommendationsDTOs := make([]domain.RecommendationDTO, lenght)
	for i := 0; i < lenght; i++{
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
	Method to calculate the recommendation score for a given stock
*/
func (s *StocksServiceImpl) getRecommendationScore(stock domain.StockModel, sentimentScore float64) float64{
	//calculate the recommendation score
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
	return score
}

/*
	Method to get the avarage sentiment of the given symbol
*/
func (s *StocksServiceImpl) getAvarageSentiment(ticker string) float64{
	sentiment, err := s.externalAPI.GetStockSentiment(ticker)
	if err != nil{
		return 0
	}
	//read sentiment data and calculate the avarage sentiment 
	var avrgSentiment float64 
	for _, sent := range sentiment.Data{
		avrgSentiment += sent.Mspr
	}
	if len(sentiment.Data) > 0 {
		avrgSentiment = avrgSentiment / float64(len(sentiment.Data))
	}
	return avrgSentiment
}