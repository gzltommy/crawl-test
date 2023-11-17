package types

type Request struct {
	Url      string
	ParseFun func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
