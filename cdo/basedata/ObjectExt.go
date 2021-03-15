package basedata

import (
	"strconv"
	"fmt"
	."cdo/common"
	"reflect"
)

type ObjectExt struct {
	BValue    bool
	ByValue   byte
	ShValue   int16
	NValue    int
	LValue    int64
	FValue    float32
	DblValue  float64
	StrValue  string
	CdoValue  CDO
	BsValue   []bool
	BysValue  []byte
	ShsValue  []int16
	NsValue   []int
	LsValue   []int64
	FsValue   []float32
	DblsValue []float64
	StrsValue []string
	CdosValue []CDO

	NType int
}

func (o *ObjectExt) GetValue() interface{} {
	switch o.NType {
	case BOOLEAN_TYPE:
		return o.BValue
	case BYTE_TYPE:
		return o.ByValue
	case SHORT_TYPE:
		return o.ShValue
	case INTEGER_TYPE:
		return o.NValue
	case LONG_TYPE:
		return o.LValue
	case FLOAT_TYPE:
		return o.FValue
	case DOUBLE_TYPE:
		return o.DblValue
	case STRING_TYPE:
		fallthrough
	case DATE_TYPE:
		fallthrough
	case TIME_TYPE:
		fallthrough
	case DATETIME_TYPE:
		return o.StrValue
	case CDO_TYPE:
		return o.CdoValue
	case RECORD_TYPE:
	case BOOLEAN_ARRAY_TYPE:
		return o.BsValue
	case BYTE_ARRAY_TYPE:
		return o.BysValue
	case SHORT_ARRAY_TYPE:
		return o.ShsValue
	case INTEGER_ARRAY_TYPE:
		return o.NsValue
	case LONG_ARRAY_TYPE:
		return o.LsValue
	case FLOAT_ARRAY_TYPE:
		return o.FsValue
	case DOUBLE_ARRAY_TYPE:
		return o.DblsValue
	case STRING_ARRAY_TYPE:
		fallthrough
	case DATE_ARRAY_TYPE:
		fallthrough
	case TIME_ARRAY_TYPE:
		fallthrough
	case DATETIME_ARRAY_TYPE:
		return o.StrsValue
	case CDO_ARRAY_TYPE:
		return o.CdosValue
	case RECORD_SET_TYPE:
	default:
		return "Invalid data type " + strconv.Itoa(o.NType)
	}
	return nil
}
func (o *ObjectExt) GetType() int {
	return o.NType
}

func (o *ObjectExt) GetLength() int {
	switch o.NType {
	case BOOLEAN_ARRAY_TYPE:
		return len(o.BsValue)
	case BYTE_ARRAY_TYPE:
		return len(o.BysValue)
	case SHORT_ARRAY_TYPE:
		return len(o.ShsValue)
	case INTEGER_ARRAY_TYPE:
		return len(o.NsValue)
	case LONG_ARRAY_TYPE:
		return len(o.LsValue)
	case FLOAT_ARRAY_TYPE:
		return len(o.FsValue)
	case DOUBLE_ARRAY_TYPE:
		return len(o.DblsValue)
	case STRING_ARRAY_TYPE:
		fallthrough
	case DATE_ARRAY_TYPE:
		fallthrough
	case TIME_ARRAY_TYPE:
		fallthrough
	case DATETIME_ARRAY_TYPE:
		return len(o.StrsValue)
	case CDO_ARRAY_TYPE:
	case RECORD_SET_TYPE:
	}
	return -1
}
func (o *ObjectExt) IsArrayType() bool {
	if o.NType >= BOOLEAN_ARRAY_TYPE {
		return true
	}
	return false
}

func (o *ObjectExt) GetBooleanValue() bool {
	if o.NType == BOOLEAN_TYPE {
		return o.BValue
	} else if o.NType == STRING_TYPE {
		ret, _ := strconv.ParseBool(o.StrValue)
		return ret
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetByteValue() byte {
	switch o.NType {
	case BYTE_TYPE:
		return o.ByValue
	case SHORT_TYPE:
		return byte(o.ShValue)
	case INTEGER_TYPE:
		return byte(o.NValue)
	case LONG_TYPE:
		return byte(o.LValue)
	case FLOAT_TYPE:
		return byte(o.FValue)
	case DOUBLE_TYPE:
		return byte(o.DblValue)
	case STRING_TYPE:
		//TODO return byte(o.StrValue)
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetShortValue() int16 {
	switch o.NType {
	case BYTE_TYPE:
		return int16(o.ByValue)
	case SHORT_TYPE:
		return o.ShValue
	case INTEGER_TYPE:
		return int16(o.NValue)
	case LONG_TYPE:
		return int16(o.LValue)
	case FLOAT_TYPE:
		return int16(o.FValue)
	case DOUBLE_TYPE:
		return int16(o.DblValue)
	case STRING_TYPE:
		ret, err := strconv.ParseInt(o.StrValue, 10, 16)
		if err != nil {
			panic(err.Error())
		}
		return int16(ret)
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetIntegerValue() int {
	switch o.NType {
	case BYTE_TYPE:
		return int(o.ByValue)
	case SHORT_TYPE:
		return int(o.ShValue)
	case INTEGER_TYPE:
		return o.NValue
	case LONG_TYPE:
		return int(o.LValue)
	case FLOAT_TYPE:
		return int(o.FValue)
	case DOUBLE_TYPE:
		return int(o.DblValue)
	case STRING_TYPE:
		ret, err := strconv.Atoi(o.StrValue)
		if err != nil {
			panic(err.Error())
		}
		return ret
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetLongValue() int64 {
	switch o.NType {
	case BYTE_TYPE:
		return int64(o.ByValue)
	case SHORT_TYPE:
		return int64(o.ShValue)
	case INTEGER_TYPE:
		return int64(o.NValue)
	case LONG_TYPE:
		return o.LValue
	case FLOAT_TYPE:
		return int64(o.FValue)
	case DOUBLE_TYPE:
		return int64(o.DblValue)
	case STRING_TYPE:
		ret, err := strconv.ParseInt(o.StrValue, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		return ret
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetFloatValue() float32 {
	switch o.NType {
	case BYTE_TYPE:
		return float32(o.ByValue)
	case SHORT_TYPE:
		return float32(o.ShValue)
	case INTEGER_TYPE:
		return float32(o.NValue)
	case LONG_TYPE:
		return float32(o.LValue)
	case FLOAT_TYPE:
		return o.FValue
	case DOUBLE_TYPE:
		return float32(o.DblValue)
	case STRING_TYPE:
		ret, err := strconv.ParseFloat(o.StrValue, 32)
		if err != nil {
			panic(err.Error())
		}
		return float32(ret)
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetDoubleValue() float64 {
	switch o.NType {
	case BYTE_TYPE:
		return float64(o.ByValue)
	case SHORT_TYPE:
		return float64(o.ShValue)
	case INTEGER_TYPE:
		return float64(o.NValue)
	case LONG_TYPE:
		return float64(o.LValue)
	case FLOAT_TYPE:
		return float64(o.FValue)
	case DOUBLE_TYPE:
		return o.DblValue
	case STRING_TYPE:
		ret, err := strconv.ParseFloat(o.StrValue, 64)
		if err != nil {
			panic(err.Error())
		}
		return ret
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetStringValue() string {
	switch o.NType {
	case BOOLEAN_TYPE:
		if o.BValue {
			return "true"
		}
		return "false"
	case BYTE_TYPE:
		return string(o.ByValue)
	case SHORT_TYPE:
		return strconv.Itoa(int(o.ShValue))
	case INTEGER_TYPE:
		return strconv.Itoa(o.NValue)
	case LONG_TYPE:
		return strconv.FormatInt(o.LValue, 10)
	case FLOAT_TYPE:
		return strconv.FormatFloat(float64(o.FValue), 'f', -1, 32)
	case DOUBLE_TYPE:
		return strconv.FormatFloat(o.DblValue, 'f', -1, 64)
	case STRING_TYPE:
		fallthrough
	case DATE_TYPE:
		fallthrough
	case TIME_TYPE:
		fallthrough
	case DATETIME_TYPE:
		return o.StrValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetDateValue() string {
	switch o.NType {
	case DATE_TYPE:
		return o.StrValue
	case DATETIME_TYPE:
		return string(o.StrValue[0:10])
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetTimeValue() string {
	switch o.NType {
	case TIME_TYPE:
		return o.StrValue
	case DATETIME_TYPE:
		return string(o.StrValue[11:])
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetDateTimeValue() string {
	switch o.NType {
	case DATETIME_TYPE:
		return o.StrValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetCDOValue() CDO {
	//TODO
	if o.NType == CDO_TYPE{
		return o.CdoValue
	}
	panic("Type cast failed")
}

func (o *ObjectExt) GetBooleanArrayValue() []bool {
	switch o.NType {
	case BOOLEAN_ARRAY_TYPE:
		return o.BsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetByteArrayValue() []byte {
	switch o.NType {
	case BYTE_ARRAY_TYPE:
		return o.BysValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetShortArrayValue() []int16 {
	switch o.NType {
	case SHORT_ARRAY_TYPE:
		return o.ShsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetIntegerArrayValue() []int {
	switch o.NType {
	case INTEGER_ARRAY_TYPE:
		return o.NsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetLongArrayValue() []int64 {
	switch o.NType {
	case LONG_ARRAY_TYPE:
		return o.LsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetFloatArrayValue() []float32 {
	switch o.NType {
	case FLOAT_ARRAY_TYPE:
		return o.FsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetDoubleArrayValue() []float64 {
	switch o.NType {
	case DOUBLE_ARRAY_TYPE:
		return o.DblsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetStringArrayValue() []string {
	switch o.NType {
	case STRING_ARRAY_TYPE:
		return o.StrsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetDateArrayValue() []string {
	switch o.NType {
	case DATE_ARRAY_TYPE:
		return o.StrsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetTimeArrayValue() []string {
	switch o.NType {
	case TIME_ARRAY_TYPE:
		return o.StrsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetDateTimeArrayValue() []string {
	switch o.NType {
	case DATETIME_ARRAY_TYPE:
		return o.StrsValue
	}
	str := fmt.Sprintf("Invalid data format:%s", strconv.Itoa(o.NType))
	panic(str)
}
func (o *ObjectExt) GetCDOArrayValue() []CDO {
	//TODO
	if o.NType == CDO_ARRAY_TYPE{
		return o.CdosValue
	}
	panic("Type cast failed")
}

func (o *ObjectExt) GetValueAt(nIndex int)interface{}{
	switch o.NType{
	case BOOLEAN_ARRAY_TYPE:
		return o.BsValue[nIndex]
	case BYTE_ARRAY_TYPE:
		return o.BysValue[nIndex]
	case SHORT_ARRAY_TYPE:
		return o.ShsValue[nIndex]
	case INTEGER_ARRAY_TYPE:
		return o.NsValue[nIndex]
	case LONG_ARRAY_TYPE:
		return o.LsValue[nIndex]
	case FLOAT_ARRAY_TYPE:
		return o.FsValue[nIndex]
	case DOUBLE_ARRAY_TYPE:
		return o.DblsValue[nIndex]
	case STRING_ARRAY_TYPE:
		return o.StrsValue[nIndex]
	case DATE_ARRAY_TYPE:
		return o.StrsValue[nIndex]
	case TIME_ARRAY_TYPE:
		return o.StrsValue[nIndex]
	case DATETIME_ARRAY_TYPE:
		return o.StrsValue[nIndex]
	case CDO_ARRAY_TYPE:
		return o.CdosValue[nIndex]
	}
	panic("Type cast failed")
}
func (o *ObjectExt) GetValueAtExt( nIndex int )ObjectExt {
	switch o.NType{
	case BOOLEAN_ARRAY_TYPE:
		return ObjectExt{NType:BOOLEAN_TYPE,BValue:o.BsValue[nIndex]}
	case BYTE_ARRAY_TYPE:
		return ObjectExt{NType:BYTE_TYPE,ByValue:o.BysValue[nIndex]}
	case SHORT_ARRAY_TYPE:
		return ObjectExt{NType:SHORT_TYPE,ShValue:o.ShsValue[nIndex]}
	case INTEGER_ARRAY_TYPE:
		return ObjectExt{NType:INTEGER_TYPE,NValue:o.NsValue[nIndex]}
	case LONG_ARRAY_TYPE:
		return ObjectExt{NType:LONG_TYPE,LValue:o.LsValue[nIndex]}
	case FLOAT_ARRAY_TYPE:
		return ObjectExt{NType:FLOAT_TYPE,FValue:o.FsValue[nIndex]}
	case DOUBLE_ARRAY_TYPE:
		return ObjectExt{NType:DOUBLE_TYPE,DblValue:o.DblsValue[nIndex]}
	case STRING_ARRAY_TYPE:
		return ObjectExt{NType:STRING_TYPE,StrValue:o.StrsValue[nIndex]}
	case DATE_ARRAY_TYPE:
		return ObjectExt{NType:DATE_TYPE,StrValue:o.StrsValue[nIndex]}
	case TIME_ARRAY_TYPE:
		return ObjectExt{NType:TIME_TYPE,StrValue:o.StrsValue[nIndex]}
	case DATETIME_ARRAY_TYPE:
		return ObjectExt{NType:DATETIME_TYPE,StrValue:o.StrsValue[nIndex]}
	case CDO_ARRAY_TYPE:
		return ObjectExt{NType:CDO_TYPE,CdoValue:o.CdosValue[nIndex]}
	}
	return ObjectExt{}
}
func (o *ObjectExt) SetValueAt(nIndex int, ext interface{} ){
	switch o.NType {
	case BOOLEAN_TYPE:
		o.BsValue[nIndex] = ext.(bool)
	case BYTE_ARRAY_TYPE:
		o.BysValue[nIndex] = ext.(byte)
	case SHORT_ARRAY_TYPE:
		o.ShsValue[nIndex] = ext.(int16)
	case INTEGER_ARRAY_TYPE:
		o.NsValue[nIndex] = ext.(int)
	case LONG_ARRAY_TYPE:
		o.LsValue[nIndex] = ext.(int64)
	case FLOAT_ARRAY_TYPE:
		o.FsValue[nIndex] = ext.(float32)
	case DOUBLE_ARRAY_TYPE:
		o.DblsValue[nIndex] = ext.(float64)
	case STRING_ARRAY_TYPE:fallthrough
	case DATE_ARRAY_TYPE:
	case TIME_ARRAY_TYPE:
	case DATETIME_ARRAY_TYPE:
		o.StrsValue[nIndex] = ext.(string)
	case CDO_ARRAY_TYPE:
		o.CdosValue[nIndex] = ext.(CDO)
	}
	panic("Type cast failed")
}

func NewObjectExt(nType int,objValue interface{})ObjectExt{
	objExt := ObjectExt{NType:nType}
	switch objValue.(type){
	case bool:
		if nType != BOOLEAN_TYPE{
			panic("Type and value not match")
		}
		objExt.BValue = objValue.(bool)
	case byte:
		if nType != BYTE_TYPE{
			panic("Type and value not match")
		}
		objExt.ByValue = objValue.(byte)
	case int16:
		if nType != SHORT_TYPE{
			panic("Type and value not match")
		}
		objExt.ShValue = objValue.(int16)
	case int:
		if nType != INTEGER_TYPE{
			panic("Type and value not match")
		}
		objExt.NValue = objValue.(int)
	case int64:
		if nType != LONG_TYPE{
			panic("Type and value not match")
		}
		objExt.LValue = objValue.(int64)
	case float32:
		if nType != FLOAT_TYPE{
			panic("Type and value not match")
		}
		objExt.FValue = objValue.(float32)
	case float64:
		if nType != DOUBLE_TYPE{
			panic("Type and value not match")
		}
		objExt.DblValue = objValue.(float64)
	case string:
		if nType != STRING_TYPE && nType != DATE_TYPE && nType != TIME_TYPE && nType != DATETIME_TYPE{
			panic("Type and value not match")
		}
		objExt.StrValue = objValue.(string)
	case CDO:
		if nType != CDO_TYPE{
			panic("Type and value not match")
		}
		objExt.CdoValue = objValue.(CDO)
	case []bool:
		if nType != BOOLEAN_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.BsValue = objValue.([]bool)
	case []byte:
		if nType != BYTE_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.BysValue = objValue.([]byte)
	case []int16:
		if nType != FLOAT_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.ShsValue = objValue.([]int16)
	case []int:
		if nType != INTEGER_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.NsValue = objValue.([]int)
	case []int64:
		if nType != LONG_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.LsValue = objValue.([]int64)
	case []float32:
		if nType != FLOAT_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.FsValue = objValue.([]float32)
	case []float64:
		if nType != DOUBLE_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.DblsValue = objValue.([]float64)
	case []string:
		if nType != STRING_ARRAY_TYPE && nType != DATE_ARRAY_TYPE &&
			nType != TIME_ARRAY_TYPE && nType != DATETIME_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.StrsValue = objValue.([]string)
	case []CDO:
		if nType != CDO_ARRAY_TYPE{
			panic("Type and value not match")
		}
		objExt.CdosValue = objValue.([]CDO)
	default:
		panic("Unsupported data type: "+reflect.TypeOf(objValue).Name())
	}
	return objExt
}
