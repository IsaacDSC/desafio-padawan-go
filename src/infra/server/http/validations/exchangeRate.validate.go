package validations

import "fmt"

func ValidateOperationsExchangeRate(list_params []string) (list_errors []string) {
	condition_error_01 := list_params[1] == "BRL" && list_params[2] == "BTC"
	condition_error_02 := list_params[1] == "USD" && list_params[2] == "BTC"
	condition_error_03 := list_params[1] == "EUR" && list_params[2] == "USD"
	condition_error_04 := list_params[1] == "EUR" && list_params[2] == "BTC"
	condition_error_05 := list_params[1] == "USD" && list_params[2] == "EUR"
	if condition_error_01 || condition_error_02 || condition_error_03 || condition_error_04 || condition_error_05 {
		list_errors = append(list_errors, fmt.Sprintf("Operation-Not-Permitted: %s-%s", list_params[1], list_params[2]))
		list_errors = append(list_errors, fmt.Sprintf("Operation-Permitted: %s", "BRL-USD"))
		list_errors = append(list_errors, fmt.Sprintf("Operation-Permitted: %s", "USD-BRL"))
		list_errors = append(list_errors, fmt.Sprintf("Operation-Permitted: %s", "BRL-EUR"))
		list_errors = append(list_errors, fmt.Sprintf("Operation-Permitted: %s", "EUR-BRL"))
		list_errors = append(list_errors, fmt.Sprintf("Operation-Permitted: %s", "BTC-USD"))
		list_errors = append(list_errors, fmt.Sprintf("Operation-Permitted: %s", "BTC-BRL"))
	}
	return
}

func ValidateTypeMoneyFromAndTo(list_params []string) (messageError []string) {
	acceptsTypeMoney := []string{"USD", "BRL", "EUR", "BTC"}
	counterErrorTO := 0
	counterErrorFROM := 0
	for index := range acceptsTypeMoney {
		if list_params[1] != acceptsTypeMoney[index] {
			if counterErrorFROM == 3 {
				messageError = append(messageError, fmt.Sprintf("Type-Money-Not-Found: %s", list_params[1]))
				messageError = append(messageError, fmt.Sprintf("Expected 'USD', 'BRL', 'EUR', 'BTC', Receiver: %s", list_params[1]))
				break
			}
			counterErrorFROM++
		}
		if list_params[2] != acceptsTypeMoney[index] {
			if counterErrorTO == 3 {
				messageError = append(messageError, fmt.Sprintf("Type-Money-Not-Found: %s", list_params[2]))
				messageError = append(messageError, fmt.Sprintf("Expected 'USD', 'BRL', 'EUR', 'BTC', Receiver: %s", list_params[2]))
				break
			}
			counterErrorTO++
		}
	}
	return
}
