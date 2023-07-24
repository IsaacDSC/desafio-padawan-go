package test

import (
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/IsaacDSC/desafio-padawan-go/external/sqlc"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/environments"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/repositories"
	"github.com/stretchr/testify/assert"
)

func init() {
	path_env, _ := filepath.Abs("../../../../.env")
	environments.StartEnv(path_env)
}

func TestSaveInfoRepository(t *testing.T) {
	repository := repositories.RateRepository{}
	err := repository.CreateExchangeRateOperationDatabase(
		sqlc.CreateInfoExchangeRateParams{
			Input:     "USD",
			Output:    "BRL",
			AmountIn:  sql.NullFloat64{Float64: float64(100.00), Valid: true},
			AmountOut: sql.NullFloat64{Float64: float64(24.98), Valid: true},
			Rate:      sql.NullFloat64{Float64: 4.99, Valid: true},
		},
	)
	assert.NoError(t, err)
}
