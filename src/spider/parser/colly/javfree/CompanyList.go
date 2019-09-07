package javfree

import (
	"github.com/gocolly/colly"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"log"
)



func CompanyListParser(_ []byte, url string) engine.Result {
	var request []engine.Request
	c := getCollector()
	c.OnXML(`//li[@id="menu-item-114217"]/ul//a`, func(element *colly.XMLElement) {
		url := element.Attr("href")
		request = append(request, engine.Request{
			Url:    url,
			Parser: CompanyParser,
		})
	})
	if err := c.Visit(url); err != nil {
		log.Printf("CompanyListParser error: %v url:%s,", err, url)
	}
	return engine.Result{
		Request: request,
		Items:   nil,
	}
}
