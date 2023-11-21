package engine

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/fetcher"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/scheduler"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"log"
)

type Processor func(request types.Request) (types.ParseResult, error)

type ConcurrentEngine struct {
	WorkCount        int
	Scheduler        scheduler.Scheduler
	ItemChan         chan types.Item
	RequestProcessor Processor
}

func (e *ConcurrentEngine) Run(requests ...types.Request) {
	out := make(chan types.ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkCount; i++ {
		e.CreateWork(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for _, r := range requests {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}

	}
}

func (e *ConcurrentEngine) CreateWork(in chan types.Request, out chan types.ParseResult, scheduler scheduler.Scheduler) {
	go func() {
		for {
			scheduler.WorkReady(in)
			request := <-in
			//result, err := worker(request)
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func Worker(r types.Request) (types.ParseResult, error) {
	log.Printf("Fetching url:%s", r.Url)

	body, err := fetcher.FetchWithProxy(r.Url)
	if err != nil {
		log.Printf("Fetch Error:%s", r.Url)
		return types.ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
