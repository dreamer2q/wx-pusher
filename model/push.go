package model

import (
	"encoding"
	"encoding/json"
	"time"
)

type PushMsg struct {
	CreateTime time.Time
	Content    string
}

func (p *PushMsg) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *PushMsg) MarshalBinary() (data []byte, err error) {
	return json.Marshal(p)
}

var _ encoding.BinaryMarshaler = &PushMsg{}
var _ encoding.BinaryUnmarshaler = &PushMsg{}
