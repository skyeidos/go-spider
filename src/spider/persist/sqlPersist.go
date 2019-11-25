package persist

import (
	"github.com/go-xorm/xorm"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"log"
)

type SQLPersist struct {
	Driver     string
	DataSource string
	ormEngine  *xorm.Engine
}

func (persist *SQLPersist) Init() error {
	var err error
	persist.ormEngine, err = xorm.NewEngine(persist.Driver, persist.Driver)
	if err != nil {
		return err
	}
	return nil
}

func (persist *SQLPersist) Close() error {
	err := persist.ormEngine.Close()
	if err != nil {
		return err
	}
	return nil
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
					log.Printf("Item Saver:error saving item %v : %v", item, err)
				}
			}
		}
	}()
	return out
}

func dbSave(persist *SQLPersist, item engine.Item) error {
	effectRows, err := persist.ormEngine.InsertOne(&item)
	if err != nil {
		log.Printf("Item Saver: error saving item %v : %d", item, err)
		return err
	} else {
		log.Printf("%d rows has effected", effectRows)
		return nil
	}
}
