package basedata

import (

)
type ValueFielder interface{
	ToXML(strbXML string)
	ToXMLWithIndent( nIndentSize int, strbXML string)
	GetObject()ObjectExt
	GetObjectValue()interface{}
}
type ValueField struct{
	Field
}