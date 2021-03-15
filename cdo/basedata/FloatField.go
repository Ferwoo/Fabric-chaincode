package basedata

import (
	"fmt"
	"strconv"
)

type FloatField struct {
	ValueField
	fValue  float32
}

func (f *FloatField) GetName() string {
	return f.strName
}
func (f *FloatField) SetName(name string) {
	f.strName = name
}
func (f *FloatField) GetValue() float32 {
	return f.fValue
}
func (f *FloatField) SetValue(value float32) {
	f.fValue = value
}
func (f *FloatField) ToXML(indentSize int) string {
	str := fmt.Sprintf("<FF N=\" + %s + \"" + " V=\"%f\"/>",f.strName,f.fValue)
	return str
}
func (f *FloatField) ToXMLWithIndent(indentSize int) string {
	var strIndent = ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}

	strXML := fmt.Sprintf("%s<FF N=\"%s\" V=\"%f\"/>\r\n",strIndent,f.strName,f.fValue)
	return strXML
}
func (f *FloatField) ToJSON() string {
	strJson := "\"" + f.strName + "\"" + ":" + strconv.FormatFloat(float64(f.fValue), 'f', -1, 32)
	return strJson
}
func (f *FloatField) ToJSONString() string {
	strJson :=  "\\\""+f.strName+"\\\":"+strconv.FormatFloat(float64(f.fValue),'f',-1,32)
	return strJson
}
