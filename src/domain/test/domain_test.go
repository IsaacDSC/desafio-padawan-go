package test

import (
	"testing"

	"github.com/IsaacDSC/desafio-padawan-go/src/domain"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/repositories"
	"github.com/stretchr/testify/assert"
)

func TestDomainExchangeRate(t *testing.T) {
	t.Run("Expect set value Rate Exchange", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.SetRate(float32(10))
		assert.Equal(t, float32(10), entity.ExchangeRate)
	})

	t.Run("Expect return list_errors get and set rate exchange", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		list_errors := entity.GetAndSetRate()
		assert.Equal(t, 1, len(list_errors))
		assert.Equal(t, "TypeMoneyEqualNil - Set primary amount", list_errors[0])
	})

	t.Run("Expect get and set rate exchange", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{Repository: &repositories.RateRepository{}}
		entity.TypeMoney = "USD"
		list_errors := entity.GetAndSetRate()
		assert.Equal(t, 0, len(list_errors))
		assert.Equal(t, float32(4.0361), entity.ExchangeRate)
	})

	t.Run("Expect return list_errors not found type-money", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{Repository: &repositories.RateRepository{}}
		entity.TypeMoney = "ANY"
		list_errors := entity.GetAndSetRate()
		assert.Equal(t, 1, len(list_errors))
		assert.Equal(t, "Not-Found-Exchange-Rate", list_errors[0])
	})

	t.Run("Expect return list_errors", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		list_errors := entity.GetAndSetRateBTC_USD()
		assert.Equal(t, 1, len(list_errors))
		assert.Equal(t, "TypeMoneyEqualNil - Set primary amount", list_errors[0])
	})

	t.Run("Expect get and set btc to usd", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{Repository: &repositories.RateRepository{}}
		entity.TypeMoney = "ANY"
		list_errors := entity.GetAndSetRateBTC_USD()
		assert.Equal(t, 0, len(list_errors))
		assert.True(t, entity.ExchangeRate > 0)
	})

	t.Run("Expect set amount BRL TO USD", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.SetAmount(float32(100.00), "BRL", "USD")
		assert.Equal(t, float32(100.00), entity.Money)
		assert.Equal(t, "BRL", entity.From)
		assert.Equal(t, "USD", entity.To)
		assert.Equal(t, "USD", entity.TypeMoney)
	})

	t.Run("Expect set amount USD TO BRL", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.SetAmount(float32(100.00), "USD", "BRL")
		assert.Equal(t, float32(100.00), entity.Money)
		assert.Equal(t, "USD", entity.From)
		assert.Equal(t, "BRL", entity.To)
		assert.Equal(t, "USD", entity.TypeMoney)
	})

	t.Run("Expect not calculate conversion USD TO BRL", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		converted := entity.CalculateMoneyConvert()
		assert.Equal(t, float32(0), converted)
		assert.Equal(t, float32(0), entity.MoneyOut)
	})

	t.Run("Expect converter USD TO BRL", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.From = "USD"
		entity.To = "BRL"
		entity.Money = float32(100.00)
		entity.ExchangeRate = float32(5)
		converted := entity.CalculateMoneyConvert()
		assert.Equal(t, float32(500), converted)
		assert.Equal(t, float32(500), entity.MoneyOut)
	})

	t.Run("Expect converter BRL TO USD", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.From = "BRL"
		entity.To = "USD"
		entity.Money = float32(100.00)
		entity.ExchangeRate = float32(5)
		converted := entity.CalculateMoneyConvert()
		assert.Equal(t, float32(20), converted)
		assert.Equal(t, float32(20), entity.MoneyOut)
	})

	t.Run("Expect not exist symbol ANY", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.To = "ANY"
		symbol := entity.GetSymbolMoney()
		assert.Equal(t, "", symbol)
	})

	t.Run("Expect get symbol US$", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.To = "USD"
		symbol := entity.GetSymbolMoney()
		assert.Equal(t, "US$", symbol)
	})

	t.Run("Expect get symbol BRL", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.To = "BRL"
		symbol := entity.GetSymbolMoney()
		assert.Equal(t, "R$", symbol)
	})

	t.Run("Expect get exchange rate", func(t *testing.T) {
		entity := domain.ConvertRateMoneyEntity{}
		entity.ExchangeRate = float32(500.00)
		rate := entity.GetExchangeRate()
		assert.Equal(t, float32(500.00), rate)
	})

}
