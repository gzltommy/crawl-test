package main

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/persist"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/rpc"
	"gopkg.in/olivere/elastic.v5"
)

func main() {

	serveRpc(":1234")
}

func serveRpc(host string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpc.ServeRpc(host, &persist.ItemService{
		Client: client,
	})
}
