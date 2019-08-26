package engine

type ParserFunction func(content []byte) Result

type Request struct {
	Url    string
	Parser ParserFunction
}

type Item struct {
	Id      string
	Payload interface{}
}

type Result struct {
	Request []Request
	Items   []Item
}

type BasePersist interface {
	Save() (chan []Item, error)
}

func NilParser(_ []byte) Result {
	return Result{}
}
