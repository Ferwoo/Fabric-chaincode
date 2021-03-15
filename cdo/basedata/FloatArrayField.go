package basedata

import (
	"fmt"
	"strconv"
)

type FloatArrayField struct {
	ArrayField
	fsValue []float32
}
func (l *FloatArrayField) GetType () string{
	return "FloatArrayField"
}
func (l *FloatArrayField)SetName( strName string){
	l.strName = strName
}
func (l *FloatArrayField)GetName()string{
	return l.strName
}
func (l *FloatArrayField)GetLength()int{
	return len(l.fsValue)
}
func (l *FloatArrayField)SetValue(lsValue []float32) {
	l.fsValue = lsValue
}
func (l *FloatArrayField)GetValue()[]float32 {
	return l.fsValue
}
func (l *FloatArrayField)SetValueAt(nIndex int, lValue float32){
	l.fsValue[nIndex] = lValue
}
func (l *FloatArrayField)GetValueAt(nIndex int)float32{
	return l.fsValue[nIndex]
}
func (l *FloatArrayField)ToXML(indentSize int )string{
	strXML := "<FAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.fsValue);i++{
		if i >0 {
			strXML += ","
		}
		strXML += fmt.Sprintf("%f",l.fsValue[i])
	}
	strXML +="\"/>"
	return strXML
}
func (l *FloatArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<FAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.fsValue);i++{
		if i>0 {
			strXML +=","
		}
		strXML += fmt.Sprintf("%f",l.fsValue[i])
	}
	strXML +="\"/>\r\n"
	return strXML
}
func (l *FloatArrayField)toJSONString()string{
	strJson := "\\\""+l.strName+"\\\""+":"+"["
	for i:=0;i<len(l.fsValue);i++{
		sign := ""
		if i == len(l.fsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJson += strconv.FormatFloat(float64(l.fsValue[i]),'f',-1,32) + sign
	}
	strJson+="]"
	return strJson
}
func (l *FloatArrayField)ToJSON()string{
	strJSON := "\""+l.strName+"\""+":"+"["
	for i:=0;i<len(l.fsValue);i++{
		sign := ""
		if i == len(l.fsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON += strconv.FormatFloat(float64(l.fsValue[i]),'f',-1,32) + sign
	}
	strJSON +="]"
	return strJSON
}
