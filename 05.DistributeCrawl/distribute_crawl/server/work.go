package main

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/rpc"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/work/server"
	"log"
)

func main() {
	log.Fatal(rpc.ServeRpc(":1235", &server.CrawlService{}))
}
