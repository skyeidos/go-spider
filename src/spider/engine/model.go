package engine

type ParserFunction func(content []byte, url string) Result

type Request struct {
	Url    string
	Parser ParserFunction
}

type BaseItem interface {
	ToArray() []string
}

type Item struct {
	Id      string
	Payload BaseItem
}

type Result struct {
	Request []Request
	Items   []Item
}

func NilParser(_ []byte) Result {
	return Result{}
}
