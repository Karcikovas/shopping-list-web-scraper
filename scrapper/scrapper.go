package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
)

type product struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Website *struct {
	Url string `json:"url"`
}

func Scrap(websites []Website) []product {
	var products []product

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML(".ingredients tbody", func(e *colly.HTMLElement) {
		newProduct := product{}

		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {

			fmt.Println(e)

			newProduct.Quantity = e.ChildAttr("span", "amount")
			newProduct.Quantity = e.ChildAttr("span", "ingredient")

			products = append(products, newProduct)
		})
	})

	for _, website := range websites {
		c.Visit(website.Url)
	}

	return products
}
