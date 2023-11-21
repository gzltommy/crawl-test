package zhengai

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/parse"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"regexp"
)

var (
	cityRe    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func City(contents []byte, _ string) types.ParseResult {

	matches := cityRe.FindAllSubmatch(contents, -1)

	result := types.ParseResult{}
	for _, m := range matches {

		//url:=string(m[1])
		name := string(m[2])
		//println(string(m[1]))
		//不用用户名了
		//result.Items = append(result.Items,"User:"+string(m[2]))
		result.Requests = append(result.Requests, types.Request{
			Url:    string(m[1]),
			Parser: NewParseUserProfile(name),
		})
	}

	//查找城市页面下的城市链接
	matches = cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, types.Request{
			Url:    string(m[1]),
			Parser: parse.NewParser(City, "ParseCity"),
		})
	}

	return result

}
