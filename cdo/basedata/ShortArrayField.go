package basedata
import (
	"strconv"
)

type ShortArrayField struct {
	ArrayField
	shsValue []int16
}
func (b *ShortArrayField)SetValue(shsValue []int16){
	b.shsValue = shsValue
}
func (b *ShortArrayField)GetValue() []int16{
	return b.shsValue
}
func (b *ShortArrayField)GetValueAt( index int)int16{
	return b.shsValue[index]
}
func (b *ShortArrayField)SetValueAt( index int, shValue int16){
	b.shsValue[index]=shValue
}
func (b *ShortArrayField)GetLength()int{
	return len(b.shsValue)
}
func (b *ShortArrayField)ToXML(indentSize int )string{
	strXML := "<BYAT N=\""+b.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(b.shsValue);i++{
		if i >0 {
			strXML += ","
		}
		strXML += strconv.Itoa(int(b.shsValue[i]))
	}
	strXML +="\"/>"
	return strXML
}
func (b *ShortArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<BYAT N=\""+b.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(b.shsValue);i++{
		if i>0 {
			strXML +=","
		}
		strXML += strconv.Itoa(int(b.shsValue[i]))
	}
	strXML +="\"/>\r\n"
	return strXML
}
func (b *ShortArrayField)toJSONString()string{
	strJson := "\\\""+b.strName+"\\\""+":"+"["
	for i:=0;i<len(b.shsValue);i++{
		sign := ""
		if i == len(b.shsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJson +=strconv.Itoa(int(b.shsValue[i]))+sign
	}
	strJson+="]"
	return strJson
}
func (b *ShortArrayField)ToJSON()string{
	strJSON := "\""+b.strName+"\""+":"+"["
	for i:=0;i<len(b.shsValue);i++{
		sign := ""
		if i == len(b.shsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON +=strconv.Itoa(int(b.shsValue[i]))+sign
	}
	strJSON +="]"
	return strJSON
}

