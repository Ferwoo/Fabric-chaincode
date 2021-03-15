package basedata

import (

)
type ArrayFielder interface {
	GetObjectValueAt(nIndex int)interface{}
	GetObjectAt( nIndex int) ObjectExt
	GetLength()int
}
func NewArrayFielder()*ArrayField{
	return nil
}
type ArrayField struct {
	ValueField
}
func (a *ArrayField)GetObjectValueAt(nIndex int) interface{}{
	return ""
}
func (a *ArrayField)GetObjectAt(nIndex int) ObjectExt{
	return ObjectExt{}
}
func (a *ArrayField)GetLength()int{
	return 0
}

