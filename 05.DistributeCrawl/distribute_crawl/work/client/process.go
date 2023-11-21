package client

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/rpc"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/work"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/engine"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
)

func CreateProcess() (engine.Processor, error) {
	client, err := rpc.NewClient(":1235")
	if err != nil {
		return nil, err
	}

	return func(req types.Request) (types.ParseResult, error) {
		sReq := work.SerializeRequest(req)
		var sResult work.ParseResult
		err := client.Call("CrawlService.Process", sReq, &sResult)
		if err != nil {
			return types.ParseResult{}, nil
		}
		return work.DeserializeResult(sResult), nil
	}, nil

}
