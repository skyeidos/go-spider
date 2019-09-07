package scheduler

import "github.com/skyeidos/go-spider/src/spider/engine"

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
				result := engine.Worker(request)
				go func() {
					channel <- result
				}()
			}
		}()
	}
	return channel
}
