package model

import (
	"encoding/json"
)

type UserProfile struct {
	Name          string
	Age           int
	Marry         string
	Constellation string
	Height        int
	Weight        int
	Salary        string
}

func (p *UserProfile) String() string {
	buf, _ := json.Marshal(p)
	return string(buf)
}
