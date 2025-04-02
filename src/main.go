package main

import (
	"CurrencyConverter/model"
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("cbr.ru"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Println("Запрос:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Ответ получен с кодом:", r.StatusCode)
	})

	var currencies []model.Currency

	c.OnHTML("table.data tbody tr", func(e *colly.HTMLElement) {
		currency := model.Currency{}

		currency.ID = e.ChildText("td:nth-child(1)")
		currency.Name = e.ChildText("td:nth-child(4)")
		currency.Value = e.ChildText("td:nth-child(5)")
		currency.Quantity = e.ChildText("td:nth-child(3)")

		currencies = append(currencies, currency)
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("currencies.csv")
		if err != nil {
			log.Fatalf("Failed to create output CSV file: %s", err)
		} else {
			log.Printf("File was created")
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		writer.Comma = ';'
		headers := []string{
			"ID",
			"Name",
			"Value",
			"Quantity",
		}

		if err := writer.Write(headers); err != nil {
			log.Fatalf("Ошибка записи заголовков в CSV: %s", err)
		}

		for _, currency := range currencies {
			record := []string{
				currency.ID,
				currency.Name,
				currency.Value,
				currency.Quantity,
			}

			writer.Write(record)
		}
		defer writer.Flush()
	})

	c.Visit("https://cbr.ru/currency_base/daily/")
}
