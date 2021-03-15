package basedata

import (
	. "cdo/common"
)
type StringArrayField struct {
	ArrayField
	strsValue []string
}
func (l *StringArrayField) GetType () string{
	return "StringArrayField"
}
func (l *StringArrayField)SetName( strName string){
	l.strName = strName
}
func (l *StringArrayField)GetName()string{
	return l.strName
}
func (l *StringArrayField)GetLength()int{
	return len(l.strsValue)
}
func (l *StringArrayField)SetValue(lsValue []string) {
	l.strsValue = lsValue
}
func (l *StringArrayField)GetValue()[]string {
	return l.strsValue
}
func (l *StringArrayField)SetValueAt(nIndex int, lValue string){
	l.strsValue[nIndex] = lValue
}
func (l *StringArrayField)GetValueAt(nIndex int)string{
	return l.strsValue[nIndex]
}
func (l *StringArrayField)ToXML(indentSize int )string{
	strXML := "<STRAF N=\""+l.strName+"\">"
	strXML += " V=\""
	for i:=0; i<len(l.strsValue);i++{
		strXML += "<STR>" + EncodeToXMLText(l.strsValue[i]) + "</STR>"
	}
	strXML +="</STRAF>"
	return strXML
}
func (l *StringArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<STRAF N=\""+l.strName+"\">\r\n"
	strXML += " V=\""
	for i:=0; i<len(l.strsValue);i++{
		strXML += "\t<STR>" + EncodeToXMLText(l.strsValue[i])+"</STR>\r\n"
	}
	strXML += strIndent + "</STRAF>\r\n"
	return strXML
}
func (l *StringArrayField)toJSONString()string{
	strJson := "\\\""+l.strName+"\\\""+":"+"["
	for i:=0;i<len(l.strsValue);i++{
		sign := ""
		if i == len(l.strsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJson += l.strsValue[i]+sign
	}
	strJson+="]"
	return strJson
}
func (l *StringArrayField)ToJSON()string{
	strJSON := "\""+l.strName+"\""+":"+"["
	for i:=0;i<len(l.strsValue);i++{
		sign := ""
		if i == len(l.strsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON += "\""+l.strsValue[i]+"\""+sign
	}
	strJSON +="]"
	return strJSON
}