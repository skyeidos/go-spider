package engine

import (
	"fmt"
	"log"
)

type Engine struct {
	Persist   BasePersist
	Scheduler Scheduler
	results   chan Result
}

func (e *Engine) Run(seeds []Request) {
	var requests []Request
	requests = append(requests, seeds...)
	resultChannel := e.Scheduler.Run()
	for _, request := range requests {
		e.Scheduler.Submit(request)
	}
	err := e.Persist.Init()
	if err != nil {
		fmt.Printf("error has happend:%v\n", err)
	}
	defer func() {
		if err = e.Persist.Close(); err != nil {
			fmt.Printf("error has happend:%v\n", err)
		}
	}()
	itemChannel := e.Persist.Save()
	if itemChannel == nil {
		log.Fatal("channel in null")
	}
	for result := range resultChannel {
		go func() {
			for _, request := range result.Request {
				e.Scheduler.Submit(request)
			}
		}()
		go func() {
			itemChannel <- result.Items
		}()
	}
}
