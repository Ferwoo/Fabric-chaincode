package basedata

import "strc
import "fmt"(
	"fmt"
	"strconv"
)

type DoubleField struct {
	ValueField
	dblValue float64
}

func (f *DoubleField) GetName() string {
	return f.strName
}
func (f *DoubleField) SetName(name string) {
	f.strName = name
}
func (f *DoubleField) GetValue() float64 {
	return f.dblValue
}
func (f *DoubleField) SetValue(value float64) {
	f.dblValue = value
}
func (f *DoubleField) ToXML(indentSize int) string {
	str := fmt.Sprintf("<DBLF N=\" + %s + \"" + " V=\"%f\"/>",f.strName,strconv.FormatFloat(f.dblValue,'f',-1,64))

	return str
}
func (f *DoubleField) ToXMLWithIndent(indentSize int) string {
	var strIndent = ""
	for i := 0; i < indentSize; i++ {
		strIndent += "\t"
	}

	strXML := fmt.Sprintf("%s<DBLF N=\"%s\" V=\"%f\"/>\r\n",strIndent,f.strName,strconv.FormatFloat(f.dblValue,'f',-1,64))
	return strXML
}
func (f *DoubleField) ToJSON() string {
	return "\"" + f.strName + "\"" + ":" + strconv.FormatFloat(f.dblValue, 'f', -1, 64)
}
func (f *DoubleField) ToJSONString() string {
	return "\\\""+f.strName + "\\\":" + strconv.FormatFloat(f.dblValue,'f',-1,64)
}
