package basedata

import (
	. "cdo/common"
)
type StringField struct {
	ValueField
	strValue string
}

func (f *StringField) GetName() string {
	return f.strName
}
func (f *StringField) SetName(name string) {
	f.strName = name
}
func (f *StringField) GetValue() string {
	return f.strValue
}
func (f *StringField) SetValue(value string) {
	f.strValue = value
}
func (f *StringField) ToXML(indentSize int) string {
	var strXML ="<STRF N=\"" + f.strName + "\""
	strXML += " V=\"" + EncodeToXMLText(f.strValue) + "\"/>"
	return strXML
}
func (f *StringField) ToXMLWithIndent(indentSize int) string {
	var strIndent = ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}

	var strXML = strIndent + "<STRF N=\"" + f.strName + "\""
	strXML += " V=\"" + EncodeToXMLText(f.strValue) + "\"/>\r\n"
	return strXML
}
func (f *StringField) ToJSON() string {
	return "\""+f.strName+"\""+":\""+f.strValue+"\""
}
func (f *StringField) ToJSONString() string {
	return "\\\""+f.strName+"\\\""+":\\\""+f.strValue+"\\\""
}
