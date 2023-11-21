package zhengai

import (
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/model"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"regexp"
	"strconv"
)

var (
	ageRe           = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
	marryRe         = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(已婚)</div>`)
	constellationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(天秤座)</div>`)
	heightRe        = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div>`)
	weightRe        = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
	salaryRe        = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)
	idRe            = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

// UserProfile 解析用户
// name 为上一级传递过来的
func UserProfile(contents []byte, url string, name string) types.ParseResult {

	//用户结构体
	profile := model.UserProfile{}
	profile.Name = name

	//年龄   string转换为int
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	//身高
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	//体重
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	//薪水
	profile.Salary = extractString(contents, salaryRe)

	//星座
	profile.Constellation = extractString(contents, constellationRe)
	if extractString(contents, marryRe) == "" {
		profile.Marry = "未婚"
	} else {
		profile.Marry = "已婚"
	}

	result := types.ParseResult{
		Items: []types.Item{
			{
				Url:     url,
				Type:    "zhengai",
				Id:      extractString([]byte(url), idRe),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ParseUserProfile struct {
	userName string
}

func (p *ParseUserProfile) Parse(contents []byte, url string) types.ParseResult {
	return UserProfile(contents, url, p.userName)
}

func (p *ParseUserProfile) Serialize() (name string, args interface{}) {
	return "ParseUserProfile", p.userName
}

func NewParseUserProfile(name string) *ParseUserProfile {
	return &ParseUserProfile{
		userName: name,
	}
}
