package basedata
import (
	"strconv"
)

type ByteArrayField struct {
	ArrayField
	bysValue []byte
}

func (b *ByteArrayField)SetValue(bysValue []byte){
	b.bysValue = bysValue
}
func (b *ByteArrayField)GetValue() []byte{
	return b.bysValue
}
func (b *ByteArrayField)GetValueAt( index int)byte{
	return b.bysValue[index]
}
func (b *ByteArrayField)SetValueAt( index int, byValue byte){
	b.bysValue[index]=byValue
}
func (b *ByteArrayField)GetLength()int{
	return len(b.bysValue)
}
func (b *ByteArrayField)ToXML(indentSize int )string{
	strXML := "<BYAT N=\""+b.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(b.bysValue);i++{
		if i >0 {
			strXML += ","
		}
		strXML += strconv.FormatUint(uint64(b.bysValue[i]),10)
	}
	strXML +="\"/>"
	return strXML
}
func (b *ByteArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<BYAT N=\""+b.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(b.bysValue);i++{
		if i>0 {
			strXML +=","
		}
		strXML += strconv.FormatUint(uint64(b.bysValue[i]),10)
	}
	strXML +="\"/>\r\n"
	return strXML
}
func (b *ByteArrayField)toJSONString()string{
	strJson := "\\\""+b.strName+"\\\""+":"+"["
	for i:=0;i<len(b.bysValue);i++{
		sign := ""
		if i == len(b.bysValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJson += strconv.FormatUint(uint64(b.bysValue[i]),10) + sign
	}
	strJson+="]"
	return strJson
}
func (b *ByteArrayField)ToJSON()string{
	strJSON := "\""+b.strName+"\""+":"+"["
	for i:=0;i<len(b.bysValue);i++{
		sign := ""
		if i == len(b.bysValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON += strconv.FormatUint(uint64(b.bysValue[i]),10) + sign
	}
	strJSON +="]"
	return strJSON
}











