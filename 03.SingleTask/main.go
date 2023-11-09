package main

import (
	"github.com/gzltommy/crawl-test/03.SingleTask/engine"
	"github.com/gzltommy/crawl-test/03.SingleTask/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:      "https://book.douban.com/",
		ParseFun: parse.ParseContent,
	})
}
