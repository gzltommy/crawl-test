package parse

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
)

type NilParse struct {
}

func (n NilParse) Parse(contents []byte, url string) types.ParseResult {
	return types.ParseResult{}
}

func (n NilParse) Serialize() (name string, args interface{}) {
	return "NilParse", nil
}
