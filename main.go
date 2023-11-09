package main

import (
	"fmt"
	"regexp"
)

const str = `<div class="intro">
    <p>本书首先从简单的思路着手，详细介绍了理解神经网络如何工作所必须的基础知识。第一部分介绍基本的思路，包括神经网络底层的数学知识，第2部分是实践，介绍了学习 Python 编程的流行和轻松的方法，从而逐渐使用该语言构建神经网络，以能够识别人类手写的字母，特别是让其像专家所开发的网络那样地工作。第3部分是扩展，介绍如何将神经网络的性能提升到工业应用的层级，甚至让其在 Raspberry Pi 上工作。</p></div>`

var (
	//authorRegexp    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\s\S]*?<a.*?>([^<]+)</a>`)
	//publisherRegexp = regexp.MustCompile(`<span class="pl">出版社:</span>[\s\S]*?<a.*?>([^<]+)</a>`)
	//pagesRegexp = regexp.MustCompile(`<span class="pl">页数:</span> (\d+)<br>`)
	//priceRegexp = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br>`)
	//scoreRegexp = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+)</strong>`)
	introRegexp = regexp.MustCompile(`<div class="intro">[\s\S]*?<p>([^<]+)</p></div>`)
)

func main() {
	match := introRegexp.FindString(str)
	fmt.Println("==", match)

	allSubMatch := introRegexp.FindAllSubmatch([]byte(str), -1)
	for _, subMath := range allSubMatch {
		for i, matchGroup := range subMath {
			if i == 0 {
				// 属于整个表达式匹配的内容
				continue
			}
			fmt.Printf("\n group[%d]:%s", i, matchGroup)
		}
	}
}
