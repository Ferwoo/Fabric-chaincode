package basedata

import (
	"fmt"
	"strconv"
)

type DoubleArrayField struct {
	ArrayField
	dblsValue []float64
}
func (l *DoubleArrayField) GetType () string{
	return "DoubleArrayField"
}
func (l *DoubleArrayField)SetName( strName string){
	l.strName = strName
}
func (l *DoubleArrayField)GetName()string{
	return l.strName
}
func (l *DoubleArrayField)GetLength()int{
	return len(l.dblsValue)
}
func (l *DoubleArrayField)SetValue(lsValue []float64) {
	l.dblsValue = lsValue
}
func (l *DoubleArrayField)GetValue()[]float64 {
	return l.dblsValue
}
func (l *DoubleArrayField)SetValueAt(nIndex int, lValue float64){
	l.dblsValue[nIndex] = lValue
}
func (l *DoubleArrayField)GetValueAt(nIndex int)float64{
	return l.dblsValue[nIndex]
}
func (l *DoubleArrayField)ToXML(indentSize int )string{
	strXML := "<DBLAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.dblsValue);i++{
		if i >0 {
			strXML += ","
		}
		strXML += fmt.Sprintf("%f",l.dblsValue[i])
	}
	strXML +="\"/>"
	return strXML
}
func (l *DoubleArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<DBLAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.dblsValue);i++{
		if i>0 {
			strXML +=","
		}
		strXML += fmt.Sprintf("%f",l.dblsValue[i])
	}
	strXML +="\"/>\r\n"
	return strXML
}
func (l *DoubleArrayField)ToJSONString()string{
	strJson := "\\\""+l.strName+"\\\""+":"+"["
	for i:=0;i<len(l.dblsValue);i++{
		sign := ""
		if i == len(l.dblsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJson += strconv.FormatFloat(l.dblsValue[i],'f',-1,64) + sign
	}
	strJson+="]"
	return strJson
}
func (l *DoubleArrayField)ToJSON()string{
	strJSON := "\""+l.strName+"\""+":"+"["
	for i:=0;i<len(l.dblsValue);i++{
		sign := ""
		if i == len(l.dblsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON += strconv.FormatFloat(l.dblsValue[i],'f',-1,64) + sign
	}
	strJSON +="]"
	return strJSON
}
