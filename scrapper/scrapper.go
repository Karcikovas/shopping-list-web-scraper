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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Scrap() []product {
	var products []product
	var pagesToScrape []string

	pageToScrape := "https://scrapeme.live/shop/page/1/"
	pagesDiscovered := []string{pageToScrape}
	i := 1
	limit := 5

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("a.page-numbers", func(e *colly.HTMLElement) {
		newPaginationLink := e.Attr("href")

		if !contains(pagesToScrape, newPaginationLink) {
			if !contains(pagesDiscovered, newPaginationLink) {
				pagesToScrape = append(pagesToScrape, newPaginationLink)
			}
			pagesDiscovered = append(pagesDiscovered, newPaginationLink)
		}
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		newProduct := product{}

		newProduct.Url = e.ChildAttr("a", "href")
		newProduct.Image = e.ChildAttr("img", "src")
		newProduct.Name = e.ChildText("h2")
		newProduct.Price = e.ChildText(".price")

		products = append(products, newProduct)
	})

	c.OnScraped(func(response *colly.Response) {
		if len(pagesToScrape) != 0 && i < limit {
			pageToScrape = pagesToScrape[0]
			pagesToScrape = pagesToScrape[1:]

			i++
			c.Visit(pageToScrape)
		}
	})

	c.Visit(pageToScrape)

	return products
}
