package basedata

import "strconv"

type LongArrayField struct {
	ArrayField
	lsValue []int64
}
func (l *LongArrayField) GetType () string{
	return "LongArrayField"
}
func (l *LongArrayField)SetName( strName string){
	l.strName = strName
}
func (l *LongArrayField)GetName()string{
	return l.strName
}
func (l *LongArrayField)GetLength()int{
	return len(l.lsValue)
}
func (l *LongArrayField)SetValue(lsValue []int64) {
	l.lsValue = lsValue
}
func (l *LongArrayField)GetValue()[]int64 {
	return l.lsValue
}
func (l *LongArrayField)SetValueAt(nIndex int, lValue int64){
	l.lsValue[nIndex] = lValue
}
func (l *LongArrayField)GetValueAt(nIndex int)int64{
	return l.lsValue[nIndex]
}
func (l *LongArrayField)ToXML(indentSize int )string{
	strXML := "<LAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.lsValue);i++{
		if i >0 {
			strXML += ","
		}
		strXML += strconv.FormatInt(l.lsValue[i],10)
	}
	strXML +="\"/>"
	return strXML
}
func (l *LongArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<LAF N=\""+l.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(l.lsValue);i++{
		if i>0 {
			strXML +=","
		}
		strXML += strconv.FormatInt(l.lsValue[i],10)
	}
	strXML +="\"/>\r\n"
	return strXML
}
func (l *LongArrayField)toJSONString()string{
	strJson := "\\\""+l.strName+"\\\""+":"+"["
	for i:=0;i<len(l.lsValue);i++{
		sign := ""
		if i == len(l.lsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJson +=strconv.FormatInt(l.lsValue[i],10)+sign
	}
	strJson+="]"
	return strJson
}
func (l *LongArrayField)ToJSON()string{
	strJSON := "\""+l.strName+"\""+":"+"["
	for i:=0;i<len(l.lsValue);i++{
		sign := ""
		if i == len(l.lsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON +=strconv.FormatInt(l.lsValue[i],10)+sign
	}
	strJSON +="]"
	return strJSON
}
