package engine

import (
	"github.com/skyeidos/go-spider/src/spider/fetcher"
	"log"
	"os"
)

type Engine struct {
	Persist BasePersist
}

func (e *Engine) Run(seeds []Request) {
	var requests []Request
	requests = append(requests, seeds...)
	channel, _ := e.Persist.Save()
	if channel == nil {
		log.Fatal("channel is null")
	}
	for len(requests) > 0 {
		request := requests[0]
		var result Result
		result = worker(request)
		channel <- result.Items
		requests = append(requests, result.Request...)
		requests = requests[1:]
	}
}

func worker(request Request) Result {
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		os.Exit(-1)
	}
	return request.Parser(content)
}
