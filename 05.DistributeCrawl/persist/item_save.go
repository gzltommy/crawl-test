package persist

import (
	"context"
	"errors"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSave() (chan types.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan types.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Got:%d,%s", itemCount, item)
			Save(client, item)
			itemCount++
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item types.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		panic(err)
	}
	return nil
}
