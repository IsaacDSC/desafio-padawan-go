package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	server "github.com/IsaacDSC/desafio-padawan-go/src/infra/server/http"
	"github.com/IsaacDSC/desafio-padawan-go/src/services"
	"github.com/stretchr/testify/assert"
)

func init() {
	go func() {
		http_server := server.HttpServer{}
		server_http := http_server.StartServerHttp()
		http_server.SetMiddleware()
		http_server.SetRouters()
		println("[ * ] Start server http://localhost:3000/")
		if err := http.ListenAndServe(":3000", server_http); err != nil {
			panic(err)
		}
	}()
}

const BaseUrl = "http://localhost:3000"

func TestEnd2EndApplicationEndpoints(t *testing.T) {
	t.Run("Expect alive response health-check", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/", BaseUrl))
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		expect := map[string]string{}
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, "alive!", expect["status"])
	})

	t.Run("Expect converter exchange rate USD To BRL (AUTOMATIC)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/USD/BRL/rate", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, 403.6, expect.ConvertedMoney)
		assert.Equal(t, "R$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate USD To BRL (SEND RATE)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/USD/BRL/5", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, float64(500), expect.ConvertedMoney)
		assert.Equal(t, "R$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate BRL To USD (AUTOMATIC)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/BRL/USD/rate", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, 24.77, expect.ConvertedMoney)
		assert.Equal(t, "US$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate BRL To USD (SEND RATE)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/BRL/USD/5", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, float64(20), expect.ConvertedMoney)
		assert.Equal(t, "US$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate BRL To EUR (AUTOMATIC)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/BRL/EUR/rate", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, 21.25, expect.ConvertedMoney)
		assert.Equal(t, "€", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate BRL To EUR (SEND RATE)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/BRL/EUR/5", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, float64(20), expect.ConvertedMoney)
		assert.Equal(t, "€", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate EUR To BRL (AUTOMATIC)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/EUR/BRL/rate", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, 470.52, expect.ConvertedMoney)
		assert.Equal(t, "R$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate EUR To BRL (SEND RATE)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/EUR/BRL/5", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, float64(500), expect.ConvertedMoney)
		assert.Equal(t, "R$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate BTC To BRL (AUTOMATIC)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/BTC/BRL/rate", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, 2.68152e+06, expect.ConvertedMoney)
		assert.Equal(t, "R$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate BTC To BRL (SEND RATE)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/BTC/BRL/5", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, float64(500), expect.ConvertedMoney)
		assert.Equal(t, "R$", expect.SymbolMoney)
	})

	t.Run("Expect converter exchange rate BTC To USD (SEND RATE)", func(t *testing.T) {
		url := fmt.Sprintf("%s/exchange/100.00/BTC/USD/5", BaseUrl)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		var expect services.OutputExchangeRateService
		err = json.Unmarshal(body, &expect)
		assert.NoError(t, err)
		assert.Equal(t, float64(500), expect.ConvertedMoney)
		assert.Equal(t, "US$", expect.SymbolMoney)
	})
}
