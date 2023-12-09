package scrapper

import (
	"github.com/gocolly/colly"
)

type product struct {
	Url   string `json:"url"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Website *struct {
	Url string `json:"url"`
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Scrap(websites []Website) []product {
	var products []product

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		newProduct := product{}

		newProduct.Url = e.ChildAttr("a", "href")
		newProduct.Image = e.ChildAttr("img", "src")
		newProduct.Name = e.ChildText("h2")
		newProduct.Price = e.ChildText(".price")

		products = append(products, newProduct)
	})

	for _, website := range websites {
		c.Visit(website.Url)
	}

	return products
}
