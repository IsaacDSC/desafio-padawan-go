package repositories

import "github.com/IsaacDSC/desafio-padawan-go/external/sqlc"

type RateRepositoryInterface interface {
	FetchExchangeRate(typeMoney string) (output []ExchangeRatesModel, err error)
	FetchExchangeRateBTC_USD() (output ExchangeRateBTC_USD, err error)
	CreateExchangeRateOperationDatabase(input sqlc.CreateInfoExchangeRateParams) (err error)
}
