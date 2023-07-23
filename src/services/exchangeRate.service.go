package services

import (
	"math"
	"strconv"

	"github.com/IsaacDSC/desafio-padawan-go/src/domain"
)

type InputExchangeRateService struct {
	Amount string  `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Rate   float32 `json:"rate"`
}

type OutputExchangeRateService struct {
	ConvertedMoney float64 `json:"valorConvertido"`
	SymbolMoney    string  `json:"simboloMoeda"`
}

type ConvertRateMoneyService struct {
	// Entity *domain.ConvertRateMoneyEntity
	Entity domain.ConvertRateMoneyEntityInterface
}

func (this_service *ConvertRateMoneyService) ConvertFromTo(
	input InputExchangeRateService,
) (output OutputExchangeRateService, list_errors []string) {
	if input.From == "BTC" && input.To == "USD" {
		return this_service.ConvertBtcToUSD(input)
	}
	amount, err := strconv.ParseFloat(input.Amount, 32)
	if err != nil {
		list_errors = append(list_errors, err.Error())
		return
	}
	this_service.Entity.SetAmount(float32(amount), input.From, input.To)
	if input.Rate > 0 {
		this_service.Entity.SetRate(input.Rate)
	} else {
		list_errors = this_service.Entity.GetAndSetRate()
		if len(list_errors) > 0 {
			return
		}
	}
	converted := this_service.Entity.CalculateMoneyConvert()
	symbol := this_service.Entity.GetSymbolMoney()
	list_errors = this_service.Entity.SaveInfo()
	if len(list_errors) > 0 {
		return
	}
	output = OutputExchangeRateService{
		ConvertedMoney: math.Floor(float64(converted)*100) / 100,
		SymbolMoney:    symbol,
	}
	return
}

func (this_service *ConvertRateMoneyService) ConvertBtcToUSD(
	input InputExchangeRateService,
) (output OutputExchangeRateService, list_errors []string) {
	amount, err := strconv.ParseFloat(input.Amount, 32)
	if err != nil {
		list_errors = append(list_errors, err.Error())
		return
	}
	this_service.Entity.SetAmount(float32(amount), input.From, input.To)
	if input.Rate > 0 {
		this_service.Entity.SetRate(input.Rate)
	} else {
		list_errors = this_service.Entity.GetAndSetRateBTC_USD()
		if len(list_errors) > 0 {
			return
		}
	}
	converted := this_service.Entity.CalculateMoneyConvert()
	symbol := this_service.Entity.GetSymbolMoney()
	list_errors = this_service.Entity.SaveInfo()
	if len(list_errors) > 0 {
		return
	}
	output = OutputExchangeRateService{
		ConvertedMoney: math.Floor(float64(converted)*100) / 100,
		SymbolMoney:    symbol,
	}
	// data, err := this_service.Entity.Repository.FetchExchangeRateBTC_USD()
	// if err != nil {
	// 	list_errors = append(list_errors, err.Error())
	// }
	// converted = float32(data.Data.Num1.Quotes.USD.Price) * amountUSD
	return
}
