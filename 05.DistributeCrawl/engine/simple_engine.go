package engine

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/fetcher"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"log"
)

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(requests ...types.Request) {
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching url:%s", r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch Error:%s", r.Url)
			continue
		}
		parseResult := r.Parser.Parse(body, r.Url)

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item:%s \n", item)
		}
	}
}
