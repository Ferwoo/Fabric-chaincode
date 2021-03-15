package basedata

import "strconv"

type IntegerArrayField struct {
	ArrayField
	nsValue []int

}

func (b *IntegerArrayField)GetType()string{
	return "IntegerArrayField"
}
func (b *IntegerArrayField)SetValue(shsValue []int){
	b.nsValue = shsValue
}
func (b *IntegerArrayField)GetValue() []int{
	return b.nsValue
}
func (b *IntegerArrayField)GetValueAt( index int)int{
	return b.nsValue[index]
}
func (b *IntegerArrayField)SetValueAt( index int, shValue int){
	b.nsValue[index]=shValue
}
func (b *IntegerArrayField)GetLength()int{
	return len(b.nsValue)
}
func (b *IntegerArrayField)ToXML(indentSize int )string{
	strXML := "<NAF N=\""+b.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(b.nsValue);i++{
		if i >0 {
			strXML += ","
		}
		strXML += strconv.Itoa(b.nsValue[i])
	}
	strXML +="\"/>"
	return strXML
}
func (b *IntegerArrayField)ToXMLWithIndent(indentSize int)string{
	strIndent := ""
	for i:=0; i<indentSize;i++{
		strIndent += "\t"
	}
	strXML := strIndent +"<NAF N=\""+b.strName+"\""
	strXML += " V=\""
	for i:=0; i<len(b.nsValue);i++{
		if i>0 {
			strXML +=","
		}
		strXML += strconv.Itoa(b.nsValue[i])
	}
	strXML +="\"/>\r\n"
	return strXML
}
func (b *IntegerArrayField)toJSONString()string{
	strJson := "\\\""+b.strName+"\\\""+":"+"["
	for i:=0;i<len(b.nsValue);i++{
		sign := ""
		if i == len(b.nsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJson +=strconv.Itoa(b.nsValue[i])+sign
	}
	strJson+="]"
	return strJson
}
func (b *IntegerArrayField)ToJSON()string{
	strJSON := "\""+b.strName+"\""+":"+"["
	for i:=0;i<len(b.nsValue);i++{
		sign := ""
		if i == len(b.nsValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON +=strconv.Itoa(b.nsValue[i])+sign
	}
	strJSON +="]"
	return strJSON
}

