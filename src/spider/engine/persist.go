package engine

type BasePersist interface {
	Save() chan []Item
}
