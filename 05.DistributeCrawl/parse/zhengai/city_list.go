package zhengai

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/parse"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"regexp"
)

var (
	cityListRe = regexp.MustCompile(`(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
)

func CityList(contents []byte, _ string) types.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests, types.Request{
			Url:    string(m[1]),
			Parser: parse.NewParser(City, "ParseCity"),
		})
	}
	return result
}
