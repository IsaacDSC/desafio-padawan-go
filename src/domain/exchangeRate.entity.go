package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/IsaacDSC/desafio-padawan-go/external/sqlc"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/repositories"
)

type ConvertRateMoneyEntity struct {
	Repository   *repositories.RateRepository
	ExchangeRate float32
	Money        float32
	To           string
	From         string
	TypeMoney    string
	MoneyOut     float32
}

func (this_entity *ConvertRateMoneyEntity) GetRate() (list_errors []string) {
	if this_entity.TypeMoney == "" {
		list_errors = append(list_errors, "TypeMoneyEqualNil - Set primary amount")
		return
	}
	listRates, err := this_entity.Repository.FetchExchangeRate(this_entity.TypeMoney)
	if err != nil {
		list_errors = append(list_errors, err.Error())
	}
	floatValue, err := strconv.ParseFloat(listRates[0].Bid, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string:", err)
		return
	}
	this_entity.ExchangeRate = float32(floatValue)

	return
}

func (this_entity *ConvertRateMoneyEntity) SetAmount(
	money float32,
	from string,
	to string,
) {
	if to == "BRL" {
		this_entity.Money = money
		this_entity.From = from
		this_entity.To = to
		this_entity.TypeMoney = this_entity.From
		return
	}
	this_entity.Money = money
	this_entity.From = from
	this_entity.To = to
	this_entity.TypeMoney = this_entity.To
}

func (this_entity *ConvertRateMoneyEntity) CalculateMoneyConvert() (output float32) {
	if this_entity.From != "BRL" {
		output = this_entity.Money * this_entity.ExchangeRate
	} else {
		output = this_entity.Money / this_entity.ExchangeRate
	}
	this_entity.MoneyOut = output
	return
}

func (this_entity *ConvertRateMoneyEntity) GetSymbolMoney() (output string) {
	symbols := map[string]string{"BRL": "R$", "USD": "US$", "EUR": "€", "BTC": "₿"}
	output = symbols[this_entity.To]
	return
}

func (this_entity *ConvertRateMoneyEntity) SaveInfo() (list_errors []string) {
	err := this_entity.Repository.CreateExchangeRateOperationDatabase(sqlc.CreateInfoExchangeRateParams{
		Input:     this_entity.From,
		Output:    this_entity.To,
		AmountIn:  sql.NullFloat64{Float64: float64(this_entity.Money), Valid: true},
		AmountOut: sql.NullFloat64{Float64: float64(this_entity.MoneyOut), Valid: true},
		Rate:      sql.NullFloat64{Float64: float64(this_entity.ExchangeRate), Valid: true},
	})
	if err != nil {
		list_errors = append(list_errors, err.Error())
	}
	return
}
