package persist

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/skyeidos/go-spider/src/spider/engine"
)

type CSVPersist struct {
}

var fileHandle map[string]*os.File

func (persist *CSVPersist) Save() chan []engine.Item {
	out := make(chan []engine.Item)
	fileHandle = make(map[string]*os.File)
	go func() {
		itemCount := 0
		for {
			items := <-out
			for _, item := range items {
				itemCount++
				log.Printf("Item Saver: Got item #%d", itemCount)
				fileName := strings.ToLower(reflect.TypeOf(item.Payload).Name())
				err := csvSave(fileName, &item)
				if err != nil {
					log.Printf("Item Saver:error saving item %v : %v", item, err)
				}
			}
		}
	}()
	return out

}

func csvSave(fileName string, item *engine.Item) error {
	var f *os.File
	if value, ok := fileHandle[fileName]; ok {
		f = value
	} else {
		f, err := os.OpenFile(fileName+".csv", os.O_CREATE|os.O_APPEND, 0755)
		fileHandle[fileName] = f
		if err != nil {
			return err
		}
	}
	writer := csv.NewWriter(f)
	var rows []string
	rows = append(rows, item.Id)
	rows = append(rows, item.Payload.ToArray()...)
	var err = writer.Write(rows)
	writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
