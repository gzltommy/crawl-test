package scheduler

import "github.com/gzltommy/crawl-test/05.DistributeCrawl/types"

type Scheduler interface {
	Submit(request types.Request)
	Run()
	WorkReady(r chan types.Request)
	WorkChan() chan types.Request
}
