package server

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/distribute_crawl/work"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/engine"
)

type CrawlService struct {
}

func (CrawlService) Process(req work.Request, result *work.ParseResult) error {
	engineReq, err := work.DeserializeRequest(req)
	name, _ := engineReq.Parser.Serialize()
	if err != nil {
		return err
	}
	_ = name
	engineResult, err := engine.Worker(engineReq)
	*result = work.SerializeResult(engineResult)
	return nil

}
