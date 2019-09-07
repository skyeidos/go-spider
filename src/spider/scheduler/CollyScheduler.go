package scheduler

import "github.com/skyeidos/go-spider/src/spider/engine"

type CollyScheduler struct {
	requests  chan engine.Request
	WorkCount int
}

func (c *CollyScheduler) Submit(request engine.Request) {
	c.requests <- request
}

func (c *CollyScheduler) Run() chan engine.Result {
	c.requests = make(chan engine.Request, 1000)
	channel := make(chan engine.Result)
	for i := 0; i < c.WorkCount; i++ {
		go func() {
			for request := range c.requests {
				result := request.Parser(nil, request.Url)
				go func() {
					channel <- result
				}()
			}
		}()
	}
	return channel
}

func (c *CollyScheduler) Close() {
	close(c.requests)
}
