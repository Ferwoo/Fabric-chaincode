package basedata

import (
	"regexp"
	"fmt"
)

type DateArrayField struct {
	ArrayField
	strsValue []string
}
func (l *DateArrayField) GetType () string{
	return "DateArrayField"
}
func (l *DateArrayField)SetName( strName string){
	l.strName = strName
}
func (l *DateArrayField)GetName()string{
	return l.strName
}
func (l *DateArrayField)GetLength()int{
	return len(l.strsValue)
}
func (l *DateArrayField)SetValue(strsValue []string) {
	reg := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}`)
	for i:=0;i<len(strsValue);i++{
		if len(strsValue[i])==0{
			continue
		}
		if len(strsValue[i])!=10 {
			str := fmt.Sprintf("Invalid date format:%s", strsValue)
			panic(str)
		}

		if isMatch := reg.MatchString(strsValue[i]); isMatch ==false {
			panic("Invalid date format:" + strsValue[i])
		}
	}
	l.strsValue = strsValue
}
func (l *DateArrayField)GetValue()[]string {
	return l.strsValue
}
func (l *DateArrayField)SetValueAt(nIndex int, strsValue string){
	reg := regexp.MustCompile(`[0-9]{4}-[01-12]{2}-[01-31]{2}`)
	if len(strsValue)>0 && len(strsValue)!=10{
		str := fmt.Sprintf("Invalid date format:%s", strsValue)
		panic(str)
	}else{
		if isMatch := reg.MatchString(strsValue); isMatch ==false {
			str := fmt.Sprintf("Invalid date format:%s", strsValue)
			panic(str)
		}
	}
	l.strsValue[nIndex] = strsValue
}
func (l *DateArrayField)GetValueAt(nIndex int)string{
	return l.strsValue[nIndex]
}
func (l *DateArrayField)ToXML(indentSize int )string{
	strXML := "<DAF N=\""+l.strName+"\""
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
func (l *DateArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<DAF N=\""+l.strName+"\""
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
func (l *DateArrayField)toJSONString()string{
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
func (l *DateArrayField)ToJSON()string{
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
