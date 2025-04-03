package main

import (
	"CurrencyConverter/parser"
	"CurrencyConverter/utils"
	"log"
)

func main() {
	parser.Parser()

	utils.ShowCurrenciesID()
	res, err := utils.GetCurrencyByID("156")
	if err != nil {
		log.Fatalf("Could not get currency: %s", err)
	}
	log.Println(res)
}
