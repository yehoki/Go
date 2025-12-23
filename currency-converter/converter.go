package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the currency converter")
	/*
		USAGE:
		convert [from] [to] [amount]
		[from]: {string} - ISO 4217 currency code converting from
		[to]: {string} - ISO 4217 currency code converting to
		[amount]: {float64} - amount we are trying to convert
	*/
	var p = fmt.Println
	// out:
	// 	for {
	// 		var input string
	// 		fmt.Scanln(&input)
	// 		switch s.ToLower(input) {
	// 		case "q":
	// 			p("Exiting program")
	// 			break out
	// 		default:
	// 			p("Please enter a valid command")
	// 		}
	// 	}
	args := os.Args
	var from string = string(args[1])
	var to string = string(args[2])
	amount, floatParseErr := strconv.ParseFloat(args[3], 64)
	if floatParseErr != nil {
		p(floatParseErr)
	}
	expectedArgs := 4
	p(len(args))
	p(args[0])
	// args, err := fmt.Scanf(from, to, amount)
	if len(args) != 4 {
		p("Expected", expectedArgs, "arguments, got", len(args))
	} else {
		p(from, to, amount)
	}
}

func chooseCurrency(currencyCode string) string {
	return ""
}

func getCurrencyMap() map[string]string {
	currencyMap := map[string]string{
		"GBP": "Pound Sterling",
		"USD": "United States Dollar",
		"EUR": "Euro",
		"CHF": "Swiss Franc",
		"NOK": "Norwegian Krone",
	}
	return currencyMap
}

func getCurrencyExchangeRateMap() map[string]map[string]float64 {
	exchangeRateMap := make(map[string]map[string]float64)
	exchangeRateMap["GBP"] = map[string]float64{
		"GBP": 1.00,
		"USD": 1.35,
		"EUR": 1.15,
		"CHF": 1.06,
		"NOK": 13.59,
	}
	exchangeRateMap["USD"] = map[string]float64{
		"GBP": 0.74,
		"USD": 1.00,
		"EUR": 0.85,
		"CHF": 0.79,
		"NOK": 10.06,
	}
	return exchangeRateMap
}
