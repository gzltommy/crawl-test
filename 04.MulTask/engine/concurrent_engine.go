package engine

import (
	"github.com/gzltommy/crawl-test/03.SingleTask/fetcher"
	"github.com/gzltommy/crawl-test/04.MulTask/scheduler"
	"github.com/gzltommy/crawl-test/04.MulTask/types"
	"log"
)

type ConcurrentEngine struct {
	WorkCount int
	Scheduler scheduler.Scheduler
}

func (e *ConcurrentEngine) Run(requests ...types.Request) {
	in := make(chan types.Request)
	out := make(chan types.ParseResult)

	e.Scheduler.ConfigureWorkChan(in)
	for i := 0; i < e.WorkCount; i++ {
		CreateWork(in, out)
	}
	for _, r := range requests {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item:%d,%s", itemCount, item)
			itemCount++
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}

	}
}

func CreateWork(in chan types.Request, out chan types.ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r types.Request) (types.ParseResult, error) {
	log.Printf("Fetching url:%s", r.Url)

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Error:%s", r.Url)
		return types.ParseResult{}, err
	}
	return r.ParseFun(body), nil
}
