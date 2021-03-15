package basedata

import "strconv"

type IntegerField struct {
	ValueField
	nValue  int
}

func (i *IntegerField)SetName(name string){
	i.strName = name
}
func (i *IntegerField)GetName()string{
	return i.strName
}
func (i *IntegerField)SetValue(nValue int){
	i.nValue = nValue;
}
func (i *IntegerField)GetValue()int{
	return i.nValue
}
func (i *IntegerField)ToXML()string{
	var strXML = "<NF N=\"" + i.strName + "\"";
	strXML += " V=\"" + strconv.Itoa(i.nValue) + "\"/>";
	return strXML;
}
func (i *IntegerField)ToXMLWithIndent(nIndentSize int) string{
	var strIndent = "";
	for idx := 0; idx < nIndentSize; idx++{
		strIndent += "\t";
	}
	var strXML = strIndent + "<NF N=\"" + i.strName + "\"";
	strXML += " V=\"" + strconv.Itoa(i.nValue) + "\"/>\r\n";
	return strXML;
}
func (i *IntegerField)ToJSONString()string{
	return "\\\""+i.strName+"\\\""+":"+strconv.Itoa(i.nValue)
}
func (i *IntegerField)ToJSON()string{
	return "\""+i.strName+"\""+":"+strconv.Itoa(i.nValue)
}
