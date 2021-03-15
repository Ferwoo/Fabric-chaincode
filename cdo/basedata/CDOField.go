package basedata

import (
	"cdo/common"
	"fmt"
)

type CDOField struct {
	ValueField
	CdoValue CDO
}

func (c *CDOField)GetType()int{
	c.nType = common.CDO_TYPE
	return c.nType
}
func (c *CDOField)SetName(strName string){
	c.strName = strName
}
func (c *CDOField)GetName()string{
	return c.strName
}
func(c *CDOField)SetValue(cdoValue CDO){
	c.CdoValue = cdoValue
}
func (c *CDOField)GetValue()CDO{
	return c.CdoValue
}
func (c *CDOField)ToXML(strbXML string){
	strbXML += fmt.Sprintf("<CDOF N=\"%s\">%s</CDOF>\r\n",c.strName,c.CdoValue.ToXMLWithStr(strbXML))
}
func (c *CDOField)toXMLWithIndent(indentSize int, strbXML string){
	strIndent := ""
	for i:=0; i< indentSize;i++{
		strIndent += "\t"
	}
	strbXML += fmt.Sprintf("%s<CDOF N=\"%s\">\r\n",strIndent,c.strName)
	c.CdoValue.toXMLWithIndent("\t",&strbXML)
	strbXML += strIndent +"</CDOF>\r\n"
}

func (c *CDOField)toJSON()string{
	strJSON := "\""+c.strName+"\""+":"+c.CdoValue.ToJSON()
	return strJSON
}
func (c *CDOField)toJSONString()string{
	return "\\\""+c.GetName()+"\\\""+":"+c.CdoValue.ToJSON()
}