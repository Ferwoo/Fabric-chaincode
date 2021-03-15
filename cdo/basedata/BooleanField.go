package basedata

import (
	"strconv"
	."cdo/common"
)

type BooleanField struct {
	ValueField
	bValue  bool
}

func (b *BooleanField) GetValue() bool {
	return b.bValue
}
func (b *BooleanField) SetValue(bValue bool) {
	b.bValue = bValue
}
func (b *BooleanField) ToXML() string {
	strXML := "<BF N=\"" + b.GetName() + "\""
	strXML += " V=\"" +  strconv.FormatBool(b.bValue) + "\"/>"
	return strXML
}
func (b *BooleanField) ToXMLWithIndent(indentSize int) string {
	strIndent := ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}
	strXML := strIndent + "<BF N=\"" + b.strName + "\""
	strXML += " V=\"" + strconv.FormatBool(b.bValue) + "\"/>\r\n"
	return strXML
}
func (b *BooleanField)GetObjectValue()interface{}{
	return b.bValue
}
func (b *BooleanField)GetObject()ObjectExt{
	obj := new(ObjectExt)
	obj.NType = b.GetType()
	obj.BValue = b.bValue
	return *obj
}
func (b *BooleanField) ToJSONString() string {
	strJson := "\\\"" + b.strName + "\\\"" + ":" + strconv.FormatBool(b.bValue) + "\""
	return strJson
}
func (b *BooleanField) ToJSON() string {
	return "\"" + b.strName + "\"" + ":" + strconv.FormatBool(b.bValue)
}
func NewBoolField(strFieldName string,bValue bool)BooleanField{
	vField := new(BooleanField)
	vField.SetName(strFieldName)
	vField.SetType(BOOLEAN_TYPE)
	vField.bValue = bValue
	return *vField
}
