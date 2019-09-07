package scheduler

import (
	"github.com/skyeidos/go-spider/src/spider/engine"
	"github.com/skyeidos/go-spider/src/spider/fetcher"
	"os"
)

type DefaultScheduler struct {
	requests  chan engine.Request
	WorkCount int
}

func (s *DefaultScheduler) SetTaskCount(count int) {

	panic("implement me")
}

func (s *DefaultScheduler) Close() {
	close(s.requests)
}

func (s *DefaultScheduler) Submit(request engine.Request) {
	s.requests <- request
}

func (s *DefaultScheduler) Run() chan engine.Result {
	s.requests = make(chan engine.Request, 1000)
	channel := make(chan engine.Result)
	for i := 0; i < s.WorkCount; i++ {
		go func() {
			for request := range s.requests {
				result := Worker(request)
				go func() {
					channel <- result
				}()
			}
		}()
	}
	return channel
}

func Worker(request engine.Request) engine.Result {
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		os.Exit(-1)
	}
	return request.Parser(content, request.Url)
}
