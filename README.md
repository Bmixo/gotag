# gotag


gotag 是一个通过自定义golang struct中的tag来轻松实现各种功能的工具 如 i18n

## 为什么使用 gotag ？

目前如果要在go实现接口的国际化翻译，基本上现存的框架都是通过定义一些kv对字段进行匹配翻译，但是这样就涉及到程序的多个地方改动，如果需要无缝集成国际化，最简单的方法就是在返回结果上进行翻译，gotag实现了使用go原生tag标签对需要翻译的字段进行标记，然后对被标记字段进行翻译和对翻译内容进行函数判断处理，如需要定义美元、人民币金额翻译结果 $ 10 和 10元，这种场景就可以轻松通过gotag实现

## Example
```
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


```
运行结果
```
{"code":200,"msg":"query suss","version":"1.0.0"}
```
更多例子见example目录

## 项目状态

gotag 目前已经商业项目测试、生产环境使用过。

master 分支用于发布稳定版本，dev 分支用于开发，您可以尝试下载最新的 release 版本进行测试。


## 为 gotag 做贡献

gotag 是一个免费且开源的项目，我们欢迎任何人为其开发和进步贡献力量。

* 在使用过程中出现任何问题，可以通过 [issues](https://github.com/Bmixo/gotag/issues) 来反馈。
* Bug 的修复可以直接提交 Pull Request 到 dev 分支。
* 如果是增加新的功能特性，请先创建一个 issue 并做简单描述以及大致的实现方法，提议被采纳后，就可以创建一个实现新特性的 Pull Request。
* 欢迎对说明文档做出改善，帮助更多的人使用 gotag，特别是英文文档。
* 贡献代码请提交 PR 至 dev 分支，master 分支仅用于发布稳定可用版本。
