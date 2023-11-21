package work

import (
	"errors"
	"fmt"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/parse"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/parse/zhengai"
	"github.com/gzltommy/crawl-test/05.DistributeCrawl/types"
	"log"
)

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url   string
	Parse SerializeParser
}

type ParseResult struct {
	Items    []types.Item
	Requests []Request
}

func SerializeResult(r types.ParseResult) ParseResult {
	result := ParseResult{Items: r.Items}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}
func SerializeRequest(r types.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parse: SerializeParser{
			Name: name,
			Args: args,
		},
	}
}

func DeserializeResult(r ParseResult) types.ParseResult {
	result := types.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializeing:%v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}

	return result
}
func DeserializeRequest(r Request) (types.Request, error) {
	p, err := deserializeParse(r.Parse)
	if err != nil {
		return types.Request{}, err
	}
	return types.Request{
		Url:    r.Url,
		Parser: p,
	}, nil

}
func deserializeParse(p SerializeParser) (parse.Parser, error) {
	switch p.Name {
	case "ParseCityList":
		return parse.NewParser(zhengai.CityList, "ParseCityList"), nil
	case "ParseCity":
		return parse.NewParser(zhengai.City, "ParseCity"), nil

	case "ParseUserProfile":
		if useName, ok := p.Args.(string); ok {
			return zhengai.NewParseUserProfile(useName), nil
		} else {
			return nil, fmt.Errorf("invilid args:%v", p.Args)
		}
	case "NilParse":
		return parse.NilParse{}, nil
	default:
		return nil, errors.New("unknown parse name")

	}
}
