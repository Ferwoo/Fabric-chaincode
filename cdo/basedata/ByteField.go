package basedata
import (
	"strconv"
)

type ByteField struct{
	ValueField
	byValue byte
}

func (b *ByteField)SetValue(byValue byte){
	b.byValue = byValue
}
func (b *ByteField)GetValue()byte{
	return b.byValue
}
func (b *ByteField)ToXML()string{
	strbXML := "<BYF N=\""+b.GetName()+"\">"
	strbXML += " V=\""+ strconv.FormatUint(uint64(b.byValue),10) + "\"/>"
	return strbXML
}
func (b *ByteField)ToXMLWithIndent(nIndentSize int)string{
	strIndent := ""
	for i := 0; i < nIndentSize; i++ {
		strIndent += "\t"
	}
	strXML := strIndent + "<BYF N=\"" + b.strName + "\""
	strXML += " V=\"" + strconv.FormatUint(uint64(b.byValue),10) + "\"/>\r\n"
	return strXML
}
func (b *ByteField) toJSONString() string {
	strJson := "\\\"" + b.strName + "\\\"" + ":" + strconv.FormatUint(uint64(b.byValue),10) + "\""
	return strJson
}
func (b *ByteField) ToJSON() string {
	return "\"" + b.strName + "\"" + ":" + strconv.FormatUint(uint64(b.byValue),10)
}



































