package basedata

type Fielder interface {
	SetType(nType int)
	GetType() int
	SetName(strName string)
	GetName() string
	ToJSON() string
	ToJSONString() string
	GetObjectValue() interface{}
}

type Field struct {
	nType   int
	strName string
}
func (f *Field)SetType(nType int){
	f.nType = nType
}
func (f *Field)GetType()int{
	return f.nType
}
func (f *Field)SetName(strName string){
	f.strName = strName
}
func (f *Field)GetName()string{
	return f.strName
}
func (f *Field)ToJSON()string{
	return ""
}
func (f *Field)toJSONString()string{
	return ""
}
func (f *Field)GetObjectValue()interface{}{
	return ""
}