package basedata

import (
	"fmt"
	"regexp"
)

type DateTimeField struct {
	ValueField
	strValue string
}

func (f *DateTimeField) GetName() string {
	return f.strName
}
func (f *DateTimeField) SetName(name string) {
	f.strName = name
}
func (f *DateTimeField) GetValue() string {
	return f.strValue
}
func (f *DateTimeField) SetValue(value string) {
	if len(value) > 0 && len(value) != 19 {
		str := fmt.Sprintf("Invalid dateTime format,check len:%s", len(value))
		panic(str)
	}
	//reg := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}`)
	reg := regexp.MustCompile(`[0-9]{4}-0[1-9]|1[0-2]-0[1-9]|[1-2][0-9]|3[0-1] [0-9]{2}:[0-9]{2}:[0-9]{2}`)

	if isMatch := reg.MatchString(value); isMatch == false {
		str := fmt.Sprintf("Invalid dateTime format:%s", value)
		panic(str)
	}
	f.strValue = value
}
func (f *DateTimeField) ToXML(indentSize int) string {
	var strXML = "<DTF N=\"" + f.strName + "\""
	strXML += " V=\"" + f.strValue + "\"/>"
	return strXML
}
func (f *DateTimeField) ToXMLWithIndent(indentSize int) string {
	var strIndent = ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}

	var strXML = strIndent + "<DTF N=\"" + f.strName + "\""
	strXML += " V=\"" + f.strValue + "\"/>\r\n"
	return strXML
}
func (f *DateTimeField) ToJSON() string {
	return "\"" + f.strName + "\"" + ":\"" + f.strValue + "\""
}
func (f *DateTimeField) ToJSONString() string {
	return "\\\"" + f.strName + "\\\"" + ":\\\"" + f.strValue + "\\\""
}
