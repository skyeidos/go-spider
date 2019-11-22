package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"github.com/skyeidos/go-spider/src/spider/parser/colly/javfree"
	"github.com/skyeidos/go-spider/src/spider/persist"
	"github.com/skyeidos/go-spider/src/spider/scheduler"
)

func main() {
	simpleEngine := engine.Engine{
		Persist: &persist.CSVPersist{},
		Scheduler: &scheduler.CollyScheduler{
			WorkCount: 10,
		},
	}
	seeds := []engine.Request{
		{
			Url:    "https://javfree.me/",
			Parser: javfree.CompanyListParser,
		},
	}
	simpleEngine.Run(seeds)
}
