package basedata

import "strconv"

type ShortField struct {
	ValueField
	shValue int16
}

func (s *ShortField) GetName() string {
	return s.strName
}
func (s *ShortField) SetName(name string) {
	s.strName = name
}
func (s *ShortField) GetValue() int16 {
	return s.shValue
}
func (s *ShortField) SetValue(value int16) {
	s.shValue = value
}
func (s *ShortField) ToXML(indentSize int) string {
	return "<SF N=\"" + s.strName + "\"" + " V=\"" + strconv.Itoa(int(s.shValue)) + "\"/>"
}
func (s *ShortField) ToXMLWithIndent(indentSize int) string {
	var strIndent = ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}
	var strXML = strIndent + "<SF N=\"" + s.strName + "\""
	strXML += " V=\"" + strconv.Itoa(int(s.shValue)) + "\"/>\r\n"
	return strXML
}
func (s *ShortField) ToJSON() string {
	return "\"" + s.strName + "\"" + ":" + strconv.Itoa(int(s.shValue))
}
func (s *ShortField) ToJSONString() string {
	return "\\\"" + s.strName + "\\\"" + ":" + strconv.Itoa(int(s.shValue))
}