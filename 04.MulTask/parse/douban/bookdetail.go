package douban

import (
	"github.com/gzltommy/crawl-test/04.MulTask/model"
	"github.com/gzltommy/crawl-test/04.MulTask/types"
	"regexp"
)

var (
	authorRegexp    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\s\S]*?<a.*?>([^<]+)</a>`)
	publisherRegexp = regexp.MustCompile(`<span class="pl">出版社:</span>[\s\S]*?<a.*?>([^<]+)</a>`)
	pagesRegexp     = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+)<br`)
	priceRegexp     = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br`)
	scoreRegexp     = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+)</strong>`)
	introRegexp     = regexp.MustCompile(`<div class="intro">[\s\S]*?<p>([^<]+)</p></div>`)
)

func BookDetail(content []byte, bookName string) types.ParseResult {
	book := model.BookDetail{
		BookName:  bookName,
		Author:    ExtraString(content, authorRegexp),
		Publisher: ExtraString(content, publisherRegexp),
		BookPages: ExtraString(content, pagesRegexp),
		Price:     ExtraString(content, priceRegexp),
		Score:     ExtraString(content, scoreRegexp),
		Intro:     ExtraString(content, introRegexp),
	}

	result := types.ParseResult{
		Requests: nil,
		Items:    []interface{}{book.String()},
	}
	return result
}

func ExtraString(content []byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
