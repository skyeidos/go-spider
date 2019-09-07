package javfree

import (
	"github.com/gocolly/colly"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"log"
)


func CompanyParser(_ []byte, url string) engine.Result {
	var request []engine.Request
	c := getCollector()
	c.OnHTML("div.content-block.clear > div > h2 > a", func(element *colly.HTMLElement) {
		url := element.Attr("href")
		name := element.Text
		request = append(request, engine.Request{
			Url: url,
			Parser: func(content []byte, url string) engine.Result {
				return MovieParser(content, name, url)
			},
		})
	})
	c.OnHTML("a.next.page-numbers", func(element *colly.HTMLElement) {
		url := element.Attr("href")
		request = append(request, engine.Request{
			Url:    url,
			Parser: CompanyParser,
		})
	})
	if err := c.Visit(url); err != nil {
		log.Printf("CompanyParser error: %v url:%s,", err, url)
	}
	return engine.Result{
		Request: request,
		Items:   nil,
	}
}
