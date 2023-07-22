package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IsaacDSC/desafio-padawan-go/external/sqlc"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/database"
)

const exchangeRateURL = "https://economia.awesomeapi.com.br/json/daily"

type RateRepository struct{}

type ExchangeRatesModel struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func (*RateRepository) FetchExchangeRate(
	typeMoney string,
) (output []ExchangeRatesModel, err error) {
	url := fmt.Sprintf(
		"%s/%s?start_date=20180901&end_date=20180930",
		exchangeRateURL,
		typeMoney,
	)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return
	}

	return
}

type ExchangeRateBTC_USD struct {
	Data struct {
		Num1 struct {
			ID                int    `json:"id"`
			Name              string `json:"name"`
			Symbol            string `json:"symbol"`
			WebsiteSlug       string `json:"website_slug"`
			Rank              int    `json:"rank"`
			CirculatingSupply int    `json:"circulating_supply"`
			TotalSupply       int    `json:"total_supply"`
			MaxSupply         int    `json:"max_supply"`
			Quotes            struct {
				USD struct {
					Price               float64 `json:"price"`
					Volume24H           int64   `json:"volume_24h"`
					MarketCap           int64   `json:"market_cap"`
					PercentageChange1H  float64 `json:"percentage_change_1h"`
					PercentageChange24H float64 `json:"percentage_change_24h"`
					PercentageChange7D  float64 `json:"percentage_change_7d"`
					PercentChange1H     float64 `json:"percent_change_1h"`
					PercentChange24H    float64 `json:"percent_change_24h"`
					PercentChange7D     float64 `json:"percent_change_7d"`
				} `json:"USD"`
			} `json:"quotes"`
			LastUpdated int `json:"last_updated"`
		} `json:"1"`
	}
}

func (*RateRepository) FetchExchangeRateBTC_USD() (output ExchangeRateBTC_USD, err error) {
	resp, err := http.Get("https://api.alternative.me/v2/ticker/?convert=USD")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return
	}
	return
}

func (*RateRepository) CreateExchangeRateOperationDatabase(input sqlc.CreateInfoExchangeRateParams) (err error) {
	ctx := context.Background()
	conn := database.GetConnectionMysql()
	defer conn.Close()
	db := sqlc.New(conn)
	err = db.CreateInfoExchangeRate(ctx, input)
	return
}
