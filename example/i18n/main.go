package main

import (
	"fmt"
	"github.com/Bmixo/gotag"
	"github.com/Bmixo/gotag/tagdata"
)

var tranList = []tagdata.TranslateData{
	&MsgTran{
		ENData: map[string]string{
			"查询成功": "query suss",
		},
	},
}

type MsgTran struct {
	ENData map[string]string
	ZHData map[string]string
}

func (m *MsgTran) Translate(lang string, data string) (result string) {

	if lang == "en-us" {
		if _, ok := m.ENData[data]; ok {
			return m.ENData[data]
		}
	} else if lang == "zh-cn" {
		if _, ok := m.ZHData[data]; ok {
			return m.ZHData[data]
		}
	}
	return data
}
func (m *MsgTran) GetId() (id string) {
	return "msg"
}

type Response struct {
	Code    int         `json:"code" translate:""`
	Msg     string      `json:"msg" translate:"msg"`
	Version string      `json:"version"`
}

func main() {
	tran := gotag.Init(tranList)

	resp := Response{
		Code:    200,
		Msg:     "查询成功",
		Version: "1.0.0",
	}
	result, err := tran.Marshal("en-us", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
