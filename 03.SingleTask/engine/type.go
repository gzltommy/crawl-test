package engine

type Request struct {
	Url      string
	ParseFun func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}