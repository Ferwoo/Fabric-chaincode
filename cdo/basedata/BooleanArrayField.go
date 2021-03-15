package basedata

import (
	"strconv"
)

type BooleanArrayField struct {
	ArrayField
	bsValue  []bool
}
func (b *BooleanArrayField)SetValue(bsValue []bool){
	b.bsValue = bsValue
}
func (b *BooleanArrayField)GetValue()[]bool{
	return b.bsValue
}
func (b *BooleanArrayField)SetValueAt(nIndex int,bValue bool){
	b.bsValue[nIndex] = bValue
}
func (b *BooleanArrayField)GetValueAt(nIndex int)bool{
	return b.bsValue[nIndex]
}
func (b *BooleanArrayField)GetLength()int{
	return len(b.bsValue)
}
func (b *BooleanArrayField)ToXML()string{
	strbXML := "<BAF N=\""+b.GetName()+"\" V=\""
	for i:=0;i<len(b.bsValue);i++{
		if i>0{
			strbXML+=","
		}
		strbXML += strconv.FormatBool(b.bsValue[i])
	}
	strbXML += "\"/>"
	return strbXML
}
func (b *BooleanArrayField)ToXMLWithIndent(nIndentSize int )string{
	strIndent := ""
	for i := 0; i < nIndentSize; i++ {
		strIndent += "\t"
	}
	strbXML := strIndent+"<BAF N=\""+b.GetName()+"\" V=\""
	for i:=0; i<len(b.bsValue);i++{
		if i>0{
			strbXML+=","
		}
		strbXML += strconv.FormatBool(b.bsValue[i])
	}
	strbXML += "\"/>\r\n"
	return strbXML
}
func (b *BooleanArrayField)GetObjectValue()interface{}{
	return b.bsValue
}

func (b *BooleanArrayField)ToJSONString()string{
	strJSON := "\\\""+b.GetName()+"\\\":["
	length := len(b.bsValue)
	for i:=0;i<length;i++{
		sign := ""
		if i == length-1{
			sign=""
		}else{
			sign = ","
		}
		strJSON += strconv.FormatBool(b.bsValue[i])+sign
	}
	strJSON += "],"
	return strJSON
}

func (b *BooleanArrayField)ToJSON()string{
	strJSON := "\""+b.GetName()+"\":["
	length := len(b.bsValue)
	for i:=0;i<length;i++{
		sign := ""
		if i == length-1{
			sign=""
		}else{
			sign = ","
		}
		strJSON += strconv.FormatBool(b.bsValue[i])+sign
	}
	strJSON += "]"
	return strJSON
}





















