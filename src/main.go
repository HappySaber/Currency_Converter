package main

import (
	"CurrencyConverter/parser"
	"CurrencyConverter/utils"
)

func main() {
	parser.Parser()

	utils.ShowCurrenciesID()
}
