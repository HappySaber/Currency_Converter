package main

import (
	"CurrencyConverter/model"
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("https://www.sravni.ru"),
	)

	var currencies []model.Currency

	c.OnHTML("Card_card__TMpxE", func(e *colly.HTMLElement) {
		currency := model.Currency{}
		currency.Name = e.ChildText("div._44dlit _1lp9i7s")
		currency.Value = e.ChildText("div.styles_value-text__JklK4 _1lp9i7s")
	})
	fmt.Println("Hello, World")
}
