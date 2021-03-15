package basedata

import (
	"fmt"
	"regexp"
)

type DateField struct {
	ValueField
	strValue string
}


func (f *DateField) GetName() string {
	return f.strName
}
func (f *DateField) SetName(name string) {
	f.strName = name
}
func (f *DateField) GetValue() string {
	return f.strValue
}
func (f *DateField) SetValue(value string) {
	if len(value) > 0 && len(value)!=10 {
		str := fmt.Sprintf("Invalid date format:%s", value)
		panic(str)
	}
	reg := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}`)
	if isMatch := reg.MatchString(value);isMatch ==false {
		str := fmt.Sprintf("Invalid date format:%s", value)
		panic(str)
	}
	f.strValue = value
}
func (f *DateField) ToXML(indentSize int) string {
	var strXML ="<DF N=\"" + f.strName + "\""
	strXML += " V=\"" + f.strValue + "\"/>"
	return strXML
}
func (f *DateField) ToXMLWithIndent(indentSize int) string {
	var strIndent = ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}

	var strXML = strIndent + "<DF N=\"" + f.strName + "\""
	strXML += " V=\"" + f.strValue + "\"/>\r\n"
	return strXML
}
func (f *DateField) ToJSON() string {
	return "\""+f.strName+"\""+":\""+f.strValue+"\""
}
func (f *DateField) ToJSONString() string {
	return "\\\""+f.strName+"\\\""+":\\\""+f.strValue+"\\\""
}
