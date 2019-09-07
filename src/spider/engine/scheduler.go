package engine

type Scheduler interface {
	Submit(request Request)
	Run() chan Result
	Close()
}
