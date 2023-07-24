package test

import (
	"path/filepath"
	"strconv"
	"testing"

	"github.com/IsaacDSC/desafio-padawan-go/src/domain"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/environments"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/repositories"
	"github.com/IsaacDSC/desafio-padawan-go/src/services"
	"github.com/stretchr/testify/assert"
)

func init() {
	path_env, _ := filepath.Abs("../.env")
	environments.StartEnv(path_env)
}

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
		assertConverted := int64(service.Entity.GetExchangeRate() * float32(assertAmountUsd))
		assert.Equal(t, assertConverted, int64(convertedAmount.ConvertedMoney))
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
		assertConverted := int64(float32(assertAmountUsd) / service.Entity.GetExchangeRate())
		assert.Equal(t, assertConverted, int64(convertedAmount.ConvertedMoney))
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
		assertConverted := int64(float32(assertAmountBRL) / service.Entity.GetExchangeRate())
		assert.Equal(t, assertConverted, int64(convertedAmount.ConvertedMoney))
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
		assertConverted := int64(float32(assertAmountEUR) * service.Entity.GetExchangeRate())
		assert.Equal(t, assertConverted, int64(convertedAmount.ConvertedMoney))
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
		assertConverted := int64(float32(assertAmountBTC) * service.Entity.GetExchangeRate())
		assert.Equal(t, assertConverted, int64(convertedAmount.ConvertedMoney))
	})

	t.Run("Expect converter BTC to USD", func(t *testing.T) {
		service := getServiceInstanceRateMoney()
		const (
			amountBTC = "100.00"
			FROM      = "BTC"
			TO        = "USD"
		)
		input := services.InputExchangeRateService{
			Amount: amountBTC,
			From:   FROM,
			To:     TO,
		}
		convertedAmount, list_errors := service.ConvertBtcToUSD(input)
		assert.Equal(t, 0, len(list_errors))
		assertAmountBTC, err := strconv.ParseFloat(amountBTC, 32)
		assert.NoError(t, err)
		assertConverted := int64(float32(assertAmountBTC) * service.Entity.GetExchangeRate())
		assert.Equal(t, assertConverted, int64(convertedAmount.ConvertedMoney))
	})

	t.Run("Expect converter BTC to USD", func(t *testing.T) {
		service := getServiceInstanceRateMoney()
		const (
			amountBTC = "100.00"
			FROM      = "BTC"
			TO        = "USD"
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
		assertConverted := int64(float32(assertAmountBTC) * service.Entity.GetExchangeRate())
		assert.Equal(t, assertConverted, int64(convertedAmount.ConvertedMoney))
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
