package engine

type BasePersist interface {
	Init() error
	Save() chan []Item
	Close() error
}
