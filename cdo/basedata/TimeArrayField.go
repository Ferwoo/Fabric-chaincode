package basedata

import (
	"regexp"
	"fmt"
)

type TimeArrayField struct {
	ArrayField
	strsValue []string
}
func (l *TimeArrayField) GetType () string{
	return "TimeArrayField"
}
func (l *TimeArrayField)SetName( strName string){
	l.strName = strName
}
func (l *TimeArrayField)GetName()string{
	return l.strName
}
func (l *TimeArrayField)GetLength()int{
	return len(l.strsValue)
}
func (l *TimeArrayField)SetValue(strsValue []string) {
	reg := regexp.MustCompile(`[0-9]{2}:[0-9]{2}:[0-9]{2}`)
	for i:=0;i<len(strsValue);i++{
		if len(strsValue[i])==0{
			continue
		}
		if len(strsValue[i])!=8 {
			str := fmt.Sprintf("Invalid time format:", strsValue)
			panic(str)
		}

		if isMatch := reg.MatchString(strsValue[i]); isMatch ==false {
			str := fmt.Sprintf("Invalid time format:", strsValue[i])
			panic(str)
		}
	}
	l.strsValue = strsValue
}
func (l *TimeArrayField)GetValue()[]string {
	return l.strsValue
}
func (l *TimeArrayField)SetValueAt(nIndex int, strsValue string){
	reg := regexp.MustCompile(`/[0-9]{2}:[0-9]{2}:[0-9]{2}/`)
	if len(strsValue)>0 && len(strsValue)!=10{
		str := fmt.Sprintf("Invalid time format:", strsValue)
		panic(str)
	}else{
		if isMatch := reg.MatchString(strsValue); isMatch ==false {
			str := fmt.Sprintf("Invalid time format:", strsValue)
			panic(str)
		}
	}
	l.strsValue[nIndex] = strsValue
}
func (l *TimeArrayField)GetValueAt(nIndex int)string{
	return l.strsValue[nIndex]
}
func (l *TimeArrayField)ToXML(indentSize int )string{
	strXML := "<TAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.strsValue);i++{
		if i > 0{
			strXML +=","
		}
		strXML += l.strsValue[i]
	}
	strXML +="\"/>"
	return strXML
}
func (l *TimeArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<TAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.strsValue);i++{
		if i > 0{
			strXML +=","
		}
		strXML += l.strsValue[i]
	}
	strXML += "\"/>\r\n"
	return strXML
}
func (l *TimeArrayField)ToJSONString()string{
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
func (l *TimeArrayField)ToJSON()string{
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
