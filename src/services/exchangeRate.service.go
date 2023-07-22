package services

import (
	"strconv"

	"github.com/IsaacDSC/desafio-padawan-go/src/domain"
)

type InputExchangeRateService struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
	Rate   string `json:"rate"`
}

type OutputExchangeRateService struct {
	ConvertedMoney float32 `json:"valorConvertido"`
	SymbolMoney    string  `json:"simboloMoeda"`
}

type ConvertRateMoneyService struct {
	Entity *domain.ConvertRateMoneyEntity
}

func (this_service *ConvertRateMoneyService) ConvertFromTo(
	input InputExchangeRateService,
) (output OutputExchangeRateService, list_errors []string) {
	amount, err := strconv.ParseFloat(input.Amount, 32)
	if err != nil {
		list_errors = append(list_errors, err.Error())
		return
	}
	this_service.Entity.SetAmount(float32(amount), input.From, input.To)
	list_errors = this_service.Entity.GetRate()
	if len(list_errors) > 0 {
		return
	}
	converted := this_service.Entity.CalculateMoneyConvert()
	symbol := this_service.Entity.GetSymbolMoney()
	list_errors = this_service.Entity.SaveInfo()
	if len(list_errors) > 0 {
		return
	}
	output = OutputExchangeRateService{
		ConvertedMoney: converted,
		SymbolMoney:    symbol,
	}
	return
}

func (this_service *ConvertRateMoneyService) ConvertBtcToUSD(
	amountUSD float32, from_money, string, to_money string,
) (converted float32, list_errors []string) {
	data, err := this_service.Entity.Repository.FetchExchangeRateBTC_USD()
	if err != nil {
		list_errors = append(list_errors, err.Error())
	}
	converted = float32(data.Data.Num1.Quotes.USD.Price) * amountUSD
	return
}
