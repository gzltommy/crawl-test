package main

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/rpc"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/work/client"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/engine"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/parse"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/parse/zhengai"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/scheduler"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"log"
)

func main() {
	itemsave, err := ItemSave(":1234")
	if err != nil {
		panic(err)
	}
	process, err := client.CreateProcess()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		WorkCount:        100,
		Scheduler:        &scheduler.QueueScheduler{},
		ItemChan:         itemsave,
		RequestProcessor: process,
	}
	e.Run(types.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parse.NewParser(zhengai.CityList, "ParseCityList"),
	})
}

func ItemSave(host string) (chan types.Item, error) {
	rpcClient, err := rpc.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan types.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver:Got$%d,%v", itemCount, item)
			result := ""
			err = rpcClient.Call("ItemService.Save", item, &result)

			if err != nil {
				log.Printf("item saver:error saving item %v:%v", item, err)
			}
			itemCount++
		}
	}()
	return out, nil
}
