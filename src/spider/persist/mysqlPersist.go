package persist

import (
	"errors"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"log"
)

type SQLPersist struct {
	UserName string
	Password string
	Host     string
	Port     string
}

func (persist *SQLPersist) Save() chan []engine.Item {
	out := make(chan []engine.Item)
	go func() {
		itemCount := 0
		for {
			items := <-out
			for _, item := range items {
				itemCount++
				log.Printf("Item Saver: Got item #%d", itemCount)
				err := dbSave(persist, item)
				if err != nil {
					//log.Printf("Item Saver:error saving item %v : %v", item, err)
				}
			}
		}
	}()
	return out
}

func dbSave(persist *SQLPersist, item engine.Item) error {
	return errors.New("NOT implements")
}
