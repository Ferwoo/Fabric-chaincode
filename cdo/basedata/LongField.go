package basedata

import "strconv"

type LongField struct {
	ValueField
	lValue  int64
}

func (l *LongField)SetName(name string){
	l.strName = name
}
func (l *LongField)GetName()string{
	return l.strName
}
func (l *LongField)SetValue(nValue int64){
	l.lValue = nValue;
}
func (l *LongField)GetValue()int64{
	return l.lValue
}
func (l *LongField)ToXML()string{
	var strXML = "<LF N=\"" + l.strName + "\"";
	strXML += " V=\"" + strconv.FormatInt(l.lValue,10) + "\"/>";
	return strXML;
}
func (l *LongField)ToXMLWithIndent(nIndentSize int) string{
	var strIndent = "";
	for idx := 0; idx < nIndentSize; idx++{
		strIndent += "\t";
	}
	var strXML = strIndent + "<LF N=\"" + l.strName + "\"";
	strXML += " V=\"" +strconv.FormatInt(l.lValue, 10) + "\"/>\r\n";
	return strXML;
}
func (l *LongField)ToJSONString()string{
	return "\\\""+l.strName+"\\\""+":"+strconv.FormatInt(l.lValue,10)
}

func (l *LongField)ToJSON()string{
	return "\""+l.strName+"\""+":"+strconv.FormatInt(l.lValue,10)
}
