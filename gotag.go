package gotag

import (
	"github.com/Bmixo/gotag/json"
	"github.com/Bmixo/gotag/tagdata"
)

type Translate struct {
	Data        map[string]tagdata.TranslateData
	jsonMarshal *json.TranslateDataMap
}

func Init(translateDataList []tagdata.TranslateData) *Translate {
	translateData := Translate{
		Data:        map[string]tagdata.TranslateData{},
		jsonMarshal: nil,
	}

	for i := 0; i < len(translateDataList); i++ {
		translateData.Data[translateDataList[i].GetId()] = translateDataList[i]
	}

	jsonDecode := json.Init(&translateData.Data)
	translateData.jsonMarshal = jsonDecode
	return &translateData
}

func (m *Translate) Marshal(lang string, data interface{}) ([]byte, error) {
	return m.jsonMarshal.Marshal(lang, data)
}
