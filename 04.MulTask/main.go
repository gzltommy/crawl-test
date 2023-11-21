package main

import (
	"github.com/gzltommy/crawl-test/04.MulTask/engine"
	"github.com/gzltommy/crawl-test/04.MulTask/parse"
	"github.com/gzltommy/crawl-test/04.MulTask/scheduler"
	"github.com/gzltommy/crawl-test/04.MulTask/types"
)

func main() {
	e := engine.ConcurrentEngine{
		WorkCount: 100,
		//Scheduler: scheduler.NewSimpleScheduler(),
		Scheduler: scheduler.NewQueueScheduler(),
	}

	e.Run(types.Request{
		Url:      "https://book.douban.com/",
		ParseFun: parse.Tag,
	})
}
