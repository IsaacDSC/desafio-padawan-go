package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	timeInit := time.Now().Unix()
	money := 0.2
	valueUTC := make(chan float32)
	valueUSD := make(chan float32)

	go func() { //GET  BTC
		time.Sleep(time.Second * 1)
		responseValue := "2681520.00"
		floatValue, err := strconv.ParseFloat(responseValue, 32)
		if err != nil {
			fmt.Println("Erro ao converter a string:", err)
			return
		}
		valueUTC <- float32(floatValue)
	}()

	go func() { //GET USD
		time.Sleep(time.Second * 2)
		responseValue := "4.0361"
		floatValue, err := strconv.ParseFloat(responseValue, 32)
		if err != nil {
			fmt.Println("Erro ao converter a string:", err)
			return
		}
		valueUSD <- float32(floatValue)
	}()

	println("RESULT", (<-valueUTC/<-valueUSD)*float32(money))
	timeEnd := time.Now().Unix()
	fmt.Println(timeEnd - timeInit)

}

/*

  100 BR -> USD ???
  <MONEY>/TAX_BID=RESPONSE


  100 BTC -> USD ???
    -> SEARCH BRL TO BTC == VALUE_01 <2681520.00>
    -> SEARCH BRL TO USD == VALUE_02 <4.0361>
    -> CALCULATE result=VALUE_01 / VALUE_02
    -> CALCULATE 100BTC * result
*/
