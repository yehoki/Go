package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

const expectedArgs int = 4
const expectedCurrencyCodeLength int = 3

func main() {
	/*
		USAGE:
		converter [from] [to] [amount]
		[from]: {string} - ISO 4217 currency code converting from
		[to]: {string} - ISO 4217 currency code converting to
		[amount]: {float64} - amount we are trying to convert
	*/
	fmt.Println("Welcome to the currency converter")
	var p = fmt.Println
	to, from, amount := getArguments()
	var currencyNameMap map[string]string = getCurrencyMap()
	fromName := currencyNameMap[s.ToUpper(from)]
	toName := currencyNameMap[s.ToUpper(to)]
	p(amount)
	// TODO: ADD ERROR HANDLING
	if fromName == "" {
		p("Could not find valid 'from' currency code corresponding to code: ", from)
	}
	if toName == "" {
		p("Could not find valid 'to' currency code corresponding to code: ", to)
	}
	// var exchangeRateMap map[string]map[string]float64 = getCurrencyExchangeRateMap()
	// fromExchangeRateMap := exchangeRateMap
}
func getArguments() (string, string, float64) {
	args := os.Args
	argCount := len(args)
	// Check correct no. of args
	if argCount != expectedArgs {
		panic("Unexpected number of arguments")
	}
	var from string = args[1]
	if len(from) != expectedCurrencyCodeLength {
		panic("The 'from' argument is not a valid ISO 4217 currency code.")
	}
	var to string = args[2]
	if len(to) != expectedCurrencyCodeLength {
		panic("The 'to' argument is not a valid ISO 4217 currency code.")
	}
	amount, floatParseErr := strconv.ParseFloat(args[3], 64)
	if floatParseErr != nil {
		panic(floatParseErr)
	}
	if amount < 0 {
		panic("Cannot convert negative currency")
	}
	return to, from, amount
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

// func getCurrencyNames(from string, to string) (string, string) {

// }

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
