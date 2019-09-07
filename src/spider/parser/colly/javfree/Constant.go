package javfree

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"log"
)

func getCollector() *colly.Collector {
	collector := colly.NewCollector()
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	pr, err := proxy.RoundRobinProxySwitcher("socks5:localhost:1080")
	if err != nil {
		log.Fatal(err)
	}
	collector.SetProxyFunc(pr)
	extensions.RandomUserAgent(collector)
	return collector.Clone()
}
