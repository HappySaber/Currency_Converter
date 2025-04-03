package utils

import (
	"log"

	"github.com/gocolly/colly"
)

func CheckConnectionToDomain(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Request:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Response with code:", r.StatusCode)
	})
}
