package test_core

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exchangeRateURL = "https://economia.awesomeapi.com.br/json/daily"

type RateRepository struct{}

type exchangeRatesModel struct {
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

func (*RateRepository) fetchExchangeRate(
	typeMoney string,
) (output []exchangeRatesModel, err error) {
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

type ConvertRateMoneyEntity struct {
	Repository   *RateRepository
	exchangeRate float32
	money        float32
	To           string
	From         string
	typeMoney    string
}

func (this_entity *ConvertRateMoneyEntity) getRate() (list_errors []string) {
	if this_entity.typeMoney == "" {
		list_errors = append(list_errors, "TypeMoneyEqualNil - Set primary amount")
		return
	}
	listRates, err := this_entity.Repository.fetchExchangeRate(this_entity.typeMoney)
	if err != nil {
		list_errors = append(list_errors, err.Error())
	}
	floatValue, err := strconv.ParseFloat(listRates[0].Bid, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string:", err)
		return
	}
	this_entity.exchangeRate = float32(floatValue)

	return
}

func (this_entity *ConvertRateMoneyEntity) setAmount(
	money float32,
	from string,
	to string,
) {
	if to == "BRL" {
		this_entity.money = money
		this_entity.From = from
		this_entity.To = to
		this_entity.typeMoney = this_entity.From
		return
	}
	this_entity.money = money
	this_entity.From = from
	this_entity.To = to
	this_entity.typeMoney = this_entity.To
}

func (this_entity *ConvertRateMoneyEntity) calculateMoneyConvert() (output float32) {
	if this_entity.From != "BRL" {
		output = this_entity.money * this_entity.exchangeRate
	} else {
		output = this_entity.money / this_entity.exchangeRate
	}
	return
}

type ConvertRateMoneyService struct {
	entity *ConvertRateMoneyEntity
}

func (this_service *ConvertRateMoneyService) convertFromTo(
	amountUSD float32, from_money, string, to_money string,
) (converted float32, list_errors []string) {
	this_service.entity.setAmount(amountUSD, from_money, to_money)
	list_errors = this_service.entity.getRate()
	if len(list_errors) > 0 {
		return
	}
	converted = this_service.entity.calculateMoneyConvert()
	return
}

func (this_service *ConvertRateMoneyService) convertBtcToUSD(
	amountUSD float32, from_money, string, to_money string,
) (converted float32, list_errors []string) {
	go this_service.entity.Repository.fetchExchangeRate("BTC")
	go this_service.entity.Repository.fetchExchangeRate("USD")
	return
}

func getServiceInstanceRateMoney() *ConvertRateMoneyService {
	return &ConvertRateMoneyService{
		entity: &ConvertRateMoneyEntity{
			Repository: &RateRepository{},
		},
	}
}

func TestCoreApplication(t *testing.T) {
	t.Run("Expect converter USD to BRL", func(t *testing.T) {
		amountUSD := float32(100.0)
		service := getServiceInstanceRateMoney()
		const (
			FROM = "USD"
			TO   = "BRL"
		)
		convertedAmount, list_errors := service.convertFromTo(amountUSD, FROM, "", TO)
		assert.Equal(t, 0, len(list_errors))
		assert.Equal(t, service.entity.exchangeRate*amountUSD, convertedAmount)
		// fmt.Printf("%.2f USD is equivalent to %.2f BRL\n", amountUSD, convertedAmount)
	})

	t.Run("Expect converter BRL to USD", func(t *testing.T) {
		amountBRL := float32(100.0)
		service := getServiceInstanceRateMoney()
		const (
			FROM = "BRL"
			TO   = "USD"
		)
		convertedAmount, list_errors := service.convertFromTo(amountBRL, FROM, "", TO)
		assert.Equal(t, 0, len(list_errors))
		assert.Equal(t, amountBRL/service.entity.exchangeRate, convertedAmount)
		// fmt.Printf("%.2f BRL is equivalent to %.2f USD\n", amountBRL, convertedAmount)
	})

	t.Run("Expect converter BRL to EUR", func(t *testing.T) {
		amountBRL := float32(100.0)
		service := getServiceInstanceRateMoney()
		const (
			FROM = "BRL"
			TO   = "EUR"
		)
		convertedAmount, list_errors := service.convertFromTo(amountBRL, FROM, "", TO)
		assert.Equal(t, 0, len(list_errors))
		assert.Equal(t, amountBRL/service.entity.exchangeRate, convertedAmount)
		fmt.Printf("%.2f BRL is equivalent to %.2f EUR\n", amountBRL, convertedAmount)
	})

	t.Run("Expect converter EUR to BRL", func(t *testing.T) {
		amountEUR := float32(100.0)
		service := getServiceInstanceRateMoney()
		const (
			FROM = "EUR"
			TO   = "BRL"
		)
		convertedAmount, list_errors := service.convertFromTo(amountEUR, FROM, "", TO)
		assert.Equal(t, 0, len(list_errors))
		assert.Equal(t, service.entity.exchangeRate*amountEUR, convertedAmount)
		fmt.Printf("%.2f EUR is equivalent to %.2f BRL\n", amountEUR, convertedAmount)
	})

	t.Run("Expect converter BTC to BRL", func(t *testing.T) {
		amountBTC := float32(100.0)
		service := getServiceInstanceRateMoney()
		const (
			FROM = "BTC"
			TO   = "BRL"
		)
		convertedAmount, list_errors := service.convertFromTo(amountBTC, FROM, "", TO)
		assert.Equal(t, 0, len(list_errors))
		assert.Equal(t, service.entity.exchangeRate*amountBTC, convertedAmount)
		fmt.Printf("%.2f BTC is equivalent to %.2f BRL\n", amountBTC, convertedAmount)
	})

	t.Run("Expect converter BTC to USD", func(t *testing.T) {
		t.Skip()
		amountBRL := float32(100.0)
		service := getServiceInstanceRateMoney()
		const (
			FROM = "BTC"
			TO   = "USD"
		)
		convertedAmount, list_errors := service.convertBtcToUSD(amountBRL, FROM, "", TO)
		assert.Equal(t, 0, len(list_errors))
		assert.Equal(t, service.entity.exchangeRate*amountBRL, convertedAmount)
		fmt.Printf("%.2f BTC is equivalent to %.2f USD\n", amountBRL, convertedAmount)
	})

}

/*

  100 BR -> USD ???
  <MONEY>/TAX_BID=RESPONSE


  100 BTC -> USD ???
    -> SEARCH BTC TO BRL == VALUE_01 <2681520.00>
    -> SEARCH BRL TO USD == VALUE_02 <4,16>
    -> CALCULATE VALUE_01 / VALUE_02
*/
