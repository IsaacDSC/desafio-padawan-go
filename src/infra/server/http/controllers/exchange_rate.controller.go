package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/IsaacDSC/desafio-padawan-go/src/domain"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/repositories"
	"github.com/IsaacDSC/desafio-padawan-go/src/infra/server/http/validations"
	"github.com/IsaacDSC/desafio-padawan-go/src/services"
)

func getServiceInstanceRateMoney() *services.ConvertRateMoneyService {
	return &services.ConvertRateMoneyService{
		Entity: &domain.ConvertRateMoneyEntity{
			Repository: &repositories.RateRepository{},
		},
	}
}

func Get_ExchangeRateController(res http.ResponseWriter, req *http.Request) {
	list_params := strings.Split(req.RequestURI, "/")[2:]
	messageError := validations.ValidateTypeMoneyFromAndTo(list_params)
	if len(messageError) > 0 {
		output_error, err := json.Marshal(messageError)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusNotAcceptable)
		res.Write(output_error)
		return
	}
	list_errors := validations.ValidateOperationsExchangeRate(list_params)
	if len(list_errors) > 0 {
		output_error, err := json.Marshal(list_errors)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusNotAcceptable)
		res.Write(output_error)
		return
	}
	// if list_params[0]
	input := services.InputExchangeRateService{
		Amount: list_params[0],
		From:   list_params[1],
		To:     list_params[2],
		Rate:   list_params[3],
	}
	service := getServiceInstanceRateMoney()
	converted, list_errors := service.ConvertFromTo(input)
	if len(list_errors) > 0 {
		output, err := json.Marshal([]string(list_errors))
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(output)
		return
	}
	fmt.Println("converted", converted)
	output, err := json.Marshal(converted)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(output)
}
