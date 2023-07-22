package test

import (
	"strconv"
	"testing"

	"github.com/IsaacDSC/desafio-padawan-go/src/domain"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/repositories"
	"github.com/IsaacDSC/desafio-padawan-go/src/services"
	"github.com/stretchr/testify/assert"
)

func getServiceInstanceRateMoney() *services.ConvertRateMoneyService {
	return &services.ConvertRateMoneyService{
		Entity: &domain.ConvertRateMoneyEntity{
			Repository: &repositories.RateRepository{},
		},
	}
}

func TestCoreApplication(t *testing.T) {
	t.Run("Expect converter USD to BRL", func(t *testing.T) {
		service := getServiceInstanceRateMoney()
		const (
			amountUSD = "100.00"
			FROM      = "USD"
			TO        = "BRL"
		)
		input := services.InputExchangeRateService{
			Amount: amountUSD,
			From:   FROM,
			To:     TO,
		}
		convertedAmount, list_errors := service.ConvertFromTo(input)
		assert.Equal(t, 0, len(list_errors))
		assertAmountUsd, err := strconv.ParseFloat(amountUSD, 32)
		assert.NoError(t, err)
		assert.Equal(t, service.Entity.ExchangeRate*float32(assertAmountUsd), convertedAmount)
		// fmt.Printf("%.2f USD is equivalent to %.2f BRL\n", amountUSD, convertedAmount)
	})

	t.Run("Expect converter BRL to USD", func(t *testing.T) {
		service := getServiceInstanceRateMoney()
		const (
			amountBRL = "100.00"
			FROM      = "BRL"
			TO        = "USD"
		)
		input := services.InputExchangeRateService{
			Amount: amountBRL,
			From:   FROM,
			To:     TO,
		}
		convertedAmount, list_errors := service.ConvertFromTo(input)
		assert.Equal(t, 0, len(list_errors))
		assertAmountUsd, err := strconv.ParseFloat(amountBRL, 32)
		assert.NoError(t, err)
		assert.Equal(t, float32(assertAmountUsd)/service.Entity.ExchangeRate, convertedAmount)
		// fmt.Printf("%.2f BRL is equivalent to %.2f USD\n", amountBRL, convertedAmount)
	})

	t.Run("Expect converter BRL to EUR", func(t *testing.T) {
		service := getServiceInstanceRateMoney()
		const (
			amountBRL = "100.00"
			FROM      = "BRL"
			TO        = "EUR"
		)
		input := services.InputExchangeRateService{
			Amount: amountBRL,
			From:   FROM,
			To:     TO,
		}
		convertedAmount, list_errors := service.ConvertFromTo(input)
		assert.Equal(t, 0, len(list_errors))
		assertAmountBRL, err := strconv.ParseFloat(amountBRL, 32)
		assert.NoError(t, err)
		assert.Equal(t, float32(assertAmountBRL)/service.Entity.ExchangeRate, convertedAmount)
		// fmt.Printf("%.2f BRL is equivalent to %.2f EUR\n", amountBRL, convertedAmount)
	})

	t.Run("Expect converter EUR to BRL", func(t *testing.T) {
		service := getServiceInstanceRateMoney()
		const (
			amountEUR = "100.00"
			FROM      = "EUR"
			TO        = "BRL"
		)
		input := services.InputExchangeRateService{
			Amount: amountEUR,
			From:   FROM,
			To:     TO,
		}
		convertedAmount, list_errors := service.ConvertFromTo(input)
		assert.Equal(t, 0, len(list_errors))
		assertAmountEUR, err := strconv.ParseFloat(amountEUR, 32)
		assert.NoError(t, err)
		assert.Equal(t, service.Entity.ExchangeRate*float32(assertAmountEUR), convertedAmount)
		// fmt.Printf("%.2f EUR is equivalent to %.2f BRL\n", amountEUR, convertedAmount)
	})

	t.Run("Expect converter BTC to BRL", func(t *testing.T) {
		service := getServiceInstanceRateMoney()
		const (
			amountBTC = "100.00"
			FROM      = "BTC"
			TO        = "BRL"
		)
		input := services.InputExchangeRateService{
			Amount: amountBTC,
			From:   FROM,
			To:     TO,
		}
		convertedAmount, list_errors := service.ConvertFromTo(input)
		assert.Equal(t, 0, len(list_errors))
		assertAmountBTC, err := strconv.ParseFloat(amountBTC, 32)
		assert.NoError(t, err)
		assert.Equal(t, service.Entity.ExchangeRate*float32(assertAmountBTC), convertedAmount)
		// fmt.Printf("%.2f BTC is equivalent to %.2f BRL\n", amountBTC, convertedAmount)
	})

	t.Run("Expect converter BTC to USD", func(t *testing.T) {
		amountBTC := float32(100.0)
		service := getServiceInstanceRateMoney()
		const (
			FROM = "BTC"
			TO   = "USD"
		)
		value, list_errors := service.ConvertBtcToUSD(amountBTC, FROM, "", TO)
		assert.Equal(t, 0, len(list_errors))
		assert.True(t, value > 0)
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
