package persist

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/skyeidos/go-spider/src/spider/engine"
	"log"
	"reflect"
	"strings"
)

type ElasticPersist struct {
	Client *elastic.Client
}

func (persist *ElasticPersist) Close() error {
	persist.Client.Stop()
}

func (persist *ElasticPersist) Save() chan []engine.Item {
	out := make(chan []engine.Item)
	go func() {
		itemCount := 0
		for {
			items := <-out
			for _, item := range items {
				itemCount++
				log.Printf("Item Saver: Got item #%d", itemCount)
				index := strings.ToLower(reflect.TypeOf(item.Payload).Name())
				err := save(persist.Client, index, item)
				if err != nil {
					//log.Printf("Item Saver:error saving item %v : %v", item, err)
				}
			}
		}
	}()
	return out
}

func save(client *elastic.Client, index string, item engine.Item) (err error) {
	indexService := client.Index().Index(index).
		Type("doc").Id(item.Id).BodyJson(item)
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
