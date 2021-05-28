package tagdata

type TranslateData interface {
	Translate(lan string, data string) (result string)
	GetId() (id string)
}
