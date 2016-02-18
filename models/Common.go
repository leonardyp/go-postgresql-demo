package models

type CommonResp struct {
	ErrorCode int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}
