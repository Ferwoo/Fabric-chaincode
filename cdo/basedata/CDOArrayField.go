package basedata
import (
	"cdo/common"
	"fmt"
)

type CDOArrayField struct {
	ArrayField
	cdosValue []CDO
}

func (c *CDOArrayField)GetType()int{
	return common.CDO_ARRAY_TYPE
}
func (c *CDOArrayField)SetName(strName string){
	c.strName = strName
}
func (c *CDOArrayField)GetName()string{
	return c.strName
}
func (c *CDOArrayField)GetLength()int{
	return len(c.cdosValue)
}
func(c *CDOArrayField)SetValue(cdoValue []CDO){
	c.cdosValue = cdoValue
}
func (c *CDOArrayField)GetValue()[]CDO{
	return c.cdosValue
}
func (c *CDOArrayField)ToXML(nIndentSize int, strbXML string){
	strbXML ="<CDOAF N=\"" + c.strName + "\">"
	for i := 0; i < len(c.cdosValue); i++ {
		strbXML += c.cdosValue[i].ToXMLWithStr(strbXML)
	}
	strbXML += "</CDOAF>"
}
func (c *CDOArrayField)toXMLWithIndent(indentSize int, strbXML string){
	strIndent := ""
	for i:=0; i< indentSize;i++{
		strIndent += "\t"
	}
	strbXML += fmt.Sprintf("%s<CDOAF N=\"%s\">\r\n",strIndent,c.strName)
	for i := 0; i < len(c.cdosValue); i++ {
		strbXML += c.cdosValue[i].ToXMLWithIndent()
	}
	strbXML += strIndent +"</CDOAF>\r\n"
}

func (c *CDOArrayField)ToJSON()string{
	strJSON :="\""+c.GetName()+"\""+":"+"[";
	for i := 0; i < len(c.cdosValue); i++ {
		sign := ""
		if i == len(c.cdosValue)-1 {
			sign = ""
		}else{
			sign=","
		}
		strJSON+=""+c.cdosValue[i].ToJSON()+sign
	}
	strJSON+="]"
	return strJSON
}




