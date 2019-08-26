package main

import (
	"github.com/olivere/elastic"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"github.com/skyeidos/go-spider/src/spider/parser/javfree"
	"github.com/skyeidos/go-spider/src/spider/persist"
)

func main() {
	client, _ := elastic.NewClient(elastic.SetSniff(false))
	simpleEngine := engine.Engine{
		Persist: &persist.ElasticPersist{Client: client},
	}
	seeds := []engine.Request{
		{
			Url:    "https://javfree.me/",
			Parser: javfree.CompanyListParser,
		},
	}
	simpleEngine.Run(seeds)
}
