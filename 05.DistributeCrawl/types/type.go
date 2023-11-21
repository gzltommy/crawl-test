package types

import "github.com/gzltommy/crawl-test/05.DistributeCrawl/parse"

type Request struct {
	Url    string
	Parser parse.Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
