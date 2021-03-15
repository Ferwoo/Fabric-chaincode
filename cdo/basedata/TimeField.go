package basedata

import (
	"fmt"
	"regexp"
)

type TimeField struct {
	ValueField
	strValue string
}

func (f *TimeField) GetName() string {
	return f.strName
}
func (f *TimeField) SetName(name string) {
	f.strName = name
}
func (f *TimeField) GetValue() string {
	return f.strValue
}
func (f *TimeField) SetValue(value string) {
	if len(value) > 0 && len(value)!=8 {
		str := fmt.Sprintf("Invalid date format:", value)
		panic(str)
	}
	reg := regexp.MustCompile(`[0-9]{2}:[0-9]{2}:[0-9]{2}`)
	if isMatch := reg.MatchString(value);isMatch ==false {
		str := fmt.Sprintf("Invalid date format:", value)
		panic(str)
	}
	f.strValue = value
}
func (f *TimeField) ToXML(indentSize int) string {
	var strXML ="<TF N=\"" + f.strName + "\""
	strXML += " V=\"" + f.strValue + "\"/>"
	return strXML
}
func (f *TimeField) ToXMLWithIndent(indentSize int) string {
	var strIndent = ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}

	var strXML = strIndent + "<TF N=\"" + f.strName + "\""
	strXML += " V=\"" + f.strValue + "\"/>\r\n"
	return strXML
}
func (f *TimeField) ToJSON() string {
	return "\""+f.strName+"\""+":\""+f.strValue+"\""
}
func (f *TimeField) ToJSONString() string {
	return "\\\""+f.strName+"\\\""+":\\\""+f.strValue+"\\\""
}
