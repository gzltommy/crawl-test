package parse

import "github.com/gzltommy/crawl-test/05.DistributeCrawl/types"

type Parser interface {
	Parse(contents []byte, url string) types.ParseResult
	Serialize() (name string, args interface{})
}

type ParseFunc func(contents []byte, url string) types.ParseResult

// parserWrap 解析器包装
type parserWrap struct {
	parseFunc ParseFunc
	name      string
}

func NewParser(p ParseFunc, name string) Parser {
	return &parserWrap{
		parseFunc: p,
		name:      name,
	}
}

func (p parserWrap) Parse(contents []byte, url string) types.ParseResult {
	return p.parseFunc(contents, url)
}

func (p parserWrap) Serialize() (name string, args interface{}) {
	return p.name, nil
}
