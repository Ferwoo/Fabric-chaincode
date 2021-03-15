package basedata

import (
	. "cdo/common"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

const (
	SERVICE_NAME = "strServiceName"
	TRANS_NAME   = "strTransName"
)

type FieldId struct {
	nType           int
	strFieldId      string
	strMainFieldId  string
	strIndexFieldId string
}

type CDO struct {
	FieldId
	hmItem    map[string]ObjectExt
	alItem    []ObjectExt
	alFieldId []string
}

func indexOf(obj []string,key string)int{
	for k,v := range obj {
		if v == key {
			return k
		}
	}
	return -1
}
func (c *CDO) putItem(strKey string, objExt ObjectExt) {
	tempKey := strKey //strings.ToLower(strKey)			//mod by syl 2018-01-09
	if _, ok := c.hmItem[tempKey]; ok {
		c.hmItem[tempKey]=objExt
		if i := indexOf(c.alFieldId,tempKey); i >= 0 {
			c.alItem[i]    = objExt
			c.alFieldId[i] = tempKey
		}
	} else {
		c.hmItem[tempKey] = objExt
		c.alItem = append(c.alItem, objExt)
		c.alFieldId = append(c.alFieldId, tempKey)
	}
}

func (c *CDO) ParseFieldId(strFieldId string) *FieldId {
	length := len(strFieldId)

	// element of array
	if strings.HasSuffix(strFieldId, "]") {
		nMatchIndex := FindMatchedChar(length, strFieldId)
		if nMatchIndex < 1 {
			return nil
		}
		fieldId := &FieldId{
			nType:          ARRAY,
			strMainFieldId: strFieldId[0:nMatchIndex],
		}
		if len(fieldId.strMainFieldId) == 0 {
			return nil
		}
		fieldId.strIndexFieldId = strFieldId[nMatchIndex+1 : length-1]
		if len(fieldId.strIndexFieldId) == 0 {
			return nil
		}
		return fieldId
	}
	for i := length - 1; i >= 0; i-- {
		ch := string(strFieldId[i])
		if ch == "." { //multistage
			fieldId := &FieldId{}
			fieldId.nType = MULTISTAGE
			fieldId.strMainFieldId = strFieldId[0:i]
			if len(fieldId.strMainFieldId) == 0 {
				return nil
			}
			fieldId.strFieldId = strFieldId[i+1:]
			if len(fieldId.strFieldId) == 0 {
				return nil
			}
			return fieldId
		}
	}
	//simple
	fieldId := &FieldId{}
	fieldId.nType = SIMPLE
	fieldId.strFieldId = strFieldId
	return fieldId
}

func (c *CDO) ProcessXML(decoder *xml.Decoder) {
	isStr := false
	isStrAF := false
	isCdo   := false
	isCdoAf := false
	strAfName := ""
	strCdoAfName := ""
	cdosValue := []CDO{}
	strsValue := []string{}

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			//fmt.Println("xml.StartElement...")
			strTag := token.Name.Local
			attr := token.Attr
			if len(token.Attr) != 2 {
				//fmt.Printf("Attr[%s],len is[%d],token.Attr[%s]\n",strTag, len(token.Attr),token.Attr)
			}
			switch strTag {
			case "BF":
				//fmt.Println("Enter BF...")
				bValue := false
				if attr[1].Name.Local == "V" && strings.ToLower(attr[1].Value) == "true" {
					bValue = true
				} else if attr[1].Name.Local == "V" && strings.ToLower(attr[1].Value) == "false" {
					bValue = false
				} else {
					panic("parse xml error:unexpected boolean value " + attr[1].Value + " under " + token.Name.Local)
				}
				objExt := ObjectExt{NType: BOOLEAN_TYPE, BValue: bValue}
				c.putItem(attr[0].Value, objExt)
			case "BYF":
				//fmt.Println("Enter BYF...")
				byValue, err := strconv.ParseUint(attr[1].Value, 10, 8)
				if err != nil {
					panic("Convert string to uint8 err" + err.Error())
				}
				objExt := ObjectExt{NType: BYTE_TYPE, ByValue: byte(byValue)}
				c.putItem(attr[0].Value, objExt)
			case "SF":
				//fmt.Println("Enter SF...")
				sValue, err := strconv.ParseInt(attr[1].Value, 10, 16)
				if err != nil {
					panic("Convert string to short err" + err.Error())
				}
				objExt := ObjectExt{NType: SHORT_TYPE, ShValue: int16(sValue)}
				c.putItem(attr[0].Value, objExt)
			case "NF":
				//fmt.Println("Enter NF...")
				nValue, err := strconv.Atoi(attr[1].Value)
				if err != nil {
					panic("Convert string to int err" + err.Error())
				}
				objExt := ObjectExt{NType: INTEGER_TYPE, NValue: nValue}
				c.putItem(attr[0].Value, objExt)
			case "LF":
				//fmt.Println("Enter LF...")
				lValue, err := strconv.ParseInt(attr[1].Value, 10, 64)
				if err != nil {
					panic("Convert string to long err" + err.Error())
				}
				objExt := ObjectExt{NType: LONG_TYPE, LValue: lValue}
				c.putItem(attr[0].Value, objExt)
			case "FF":
				//fmt.Println("Enter FF...")
				fValue, err := strconv.ParseFloat(attr[1].Value, 32)
				if err != nil {
					panic("Convert string to float32 err" + err.Error())
				}
				objExt := ObjectExt{NType: FLOAT_TYPE, FValue: float32(fValue)}
				c.putItem(attr[0].Value, objExt)
			case "DBLF":
				//fmt.Println("Enter DBLF...")
				dValue, err := strconv.ParseFloat(attr[1].Value, 64)
				if err != nil {
					panic("Convert string to double err" + err.Error())
				}
				objExt := ObjectExt{NType: DOUBLE_TYPE, DblValue: dValue}
				c.putItem(attr[0].Value, objExt)
			case "STRF":
				//fmt.Println("Enter STRF...")
				objExt := ObjectExt{NType: STRING_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
				//fmt.Println("nType:", STRING_TYPE, " strValue:", attr[1].Value)
			case "DF":
				//fmt.Println("Enter DF...")
				objExt := ObjectExt{NType: DATE_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
			case "TF":
				//fmt.Println("Enter TF...")
				objExt := ObjectExt{NType: TIME_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
			case "DTF":
				//fmt.Println("Enter DTF...")
				objExt := ObjectExt{NType: DATETIME_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
			case "CDOF":
				cdoValue := NewCDO()
				cdoValue.ProcessXML(decoder)
				objExt := ObjectExt{NType: CDO_TYPE, CdoValue: cdoValue}
				c.putItem(attr[0].Value, objExt)
				//fmt.Println("Enter CDOF...")
			case "BAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				bsValue := make([]bool,len(strValueArr))
				for i:=0; i< len(strValueArr);i++{
					if strings.ToLower(strValueArr[i]) == "false" {
						bsValue[i] = false
					}else if strings.ToLower(strValueArr[i]) == "true"{
						bsValue[i] = true
					}
				}
				objExt := ObjectExt{NType:BOOLEAN_ARRAY_TYPE,BsValue:bsValue}
				c.putItem(attr[0].Value,objExt)
			case "BYAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				bysValue := make([]byte,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					byValue, err := strconv.ParseUint(strValueArr[i], 10, 8)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					bysValue[i] = byte(byValue)
				}
				objExt := ObjectExt{NType:BYTE_ARRAY_TYPE,BysValue:bysValue}
				c.putItem(attr[0].Value,objExt)
			case "SAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				shsValue := make([]int16,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					bValue, err := strconv.ParseInt(strValueArr[i], 10, 16)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					shsValue[i] = int16(bValue)
				}
				objExt := ObjectExt{NType:SHORT_ARRAY_TYPE,ShsValue:shsValue}
				c.putItem(attr[0].Value,objExt)
			case "NAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				nsValue := make([]int,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					bValue, err := strconv.Atoi(strValueArr[i])
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					nsValue[i] = bValue
				}
				objExt := ObjectExt{NType:INTEGER_ARRAY_TYPE,NsValue:nsValue}
				c.putItem(attr[0].Value,objExt)
			case "LAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				lsValue := make([]int64,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					lValue, err := strconv.ParseInt(strValueArr[i], 10, 64)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					lsValue[i] = lValue
				}
				objExt := ObjectExt{NType:LONG_ARRAY_TYPE,LsValue:lsValue}
				c.putItem(attr[0].Value,objExt)
			case "FAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				fsValue := make([]float32,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					fValue, err := strconv.ParseFloat(strValueArr[i], 32)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					fsValue[i] = float32(fValue)
				}
				objExt := ObjectExt{NType:FLOAT_ARRAY_TYPE,FsValue:fsValue}
				c.putItem(attr[0].Value,objExt)
			case "DBLAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				dbsValue := make([]float64,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					dValue, err := strconv.ParseFloat(strValueArr[i], 64)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					dbsValue[i] = dValue
				}
				objExt := ObjectExt{NType:DOUBLE_ARRAY_TYPE,DblsValue:dbsValue}
				c.putItem(attr[0].Value,objExt)
			case "STRAF":
				strAfName = attr[0].Value
				isStrAF = true
				fmt.Println("STRAF",attr[0].Value)
			case "STR":
				isStr = true
			case "DAF":fallthrough
			case "TAF":fallthrough
			case "DTAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				dataType := NONE_TYPE
				if strTag == "DAF"{
					dataType = DATE_ARRAY_TYPE
				}else if strTag == "TAF"{
					dataType = TIME_ARRAY_TYPE
				}else{
					dataType = DATETIME_ARRAY_TYPE
				}
				objExt := ObjectExt{NType:dataType,StrsValue:strValueArr}
				c.putItem(attr[0].Value,objExt)
			case "CDO":
				if isCdoAf {
					isCdo = true
					cdoValue := NewCDO()
					cdoValue.ProcessXMLCDOAF(decoder)
					cdosValue = append(cdosValue,cdoValue)
				}
			case "CDOAF":
				strCdoAfName = attr[0].Value
				cdosValue = []CDO{}				//Moddify by syl 2018-03-30
				isCdoAf = true
			}
		case xml.EndElement:
			//fmt.Printf("Token of '%s' end\n",token.Name.Local)
			if (token.Name.Local == "CDOF" && isCdo == false) || (token.Name.Local == "CDO" && isCdoAf == false && isCdo == true) { //|| {//( isCdoAf == false  /*&& isCdo == true*/) {
				//fmt.Println("xml.EndElement")
				//isCdo = false
				return
			}
			if token.Name.Local == "CDOAF" {
				objExt := ObjectExt{NType: CDO_ARRAY_TYPE, CdosValue: cdosValue}
				//fmt.Println("cdoaf name:",strCdoAfName,"cdosValue:",objExt.CdosValue)
				c.putItem(strCdoAfName, objExt)
				//isCdoAf = false
				strCdoAfName = ""
			}
			if token.Name.Local == "STRAF" {
				objExt := ObjectExt{NType: STRING_ARRAY_TYPE, StrsValue: strsValue}
				//fmt.Println("strsValue:",objExt.StrsValue)
				//fmt.Println("STRAF strsValue:",strsValue)
				c.putItem(strAfName, objExt)
				strAfName = ""
				strsValue = []string{""}
				isStrAF = false
				//isStr = false
			}
			if isStr{
				isStr = false
			}
		case xml.CharData:
			content := string([]byte(token))
			//fmt.Printf("This is the content: %v\n", content)
			if isStrAF && isStr {
				strsValue = append(strsValue, content)
				//fmt.Println("strsValue:",strsValue)
			}
		default:
			//fmt.Println("xml decode err:",token)
		}
	}
}

func (c *CDO) ProcessXMLCDOAF(decoder *xml.Decoder) {
	isStr := false
	isStrAF := false
	isCdo   := false
	isCdoAf := false
	strAfName := ""
	strCdoAfName := ""
	cdosValue := []CDO{}
	strsValue := []string{}

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			//fmt.Println("xml.StartElement...")
			strTag := token.Name.Local
			attr := token.Attr
			if len(token.Attr) != 2 {
				//fmt.Printf("Attr[%s],len is[%d],token.Attr[%s]\n",strTag, len(token.Attr),token.Attr)
			}
			switch strTag {
			case "BF":
				//fmt.Println("Enter BF...")
				bValue := false
				if attr[1].Name.Local == "V" && strings.ToLower(attr[1].Value) == "true" {
					bValue = true
				} else if attr[1].Name.Local == "V" && strings.ToLower(attr[1].Value) == "false" {
					bValue = false
				} else {
					panic("parse xml error:unexpected boolean value " + attr[1].Value + " under " + token.Name.Local)
				}
				objExt := ObjectExt{NType: BOOLEAN_TYPE, BValue: bValue}
				c.putItem(attr[0].Value, objExt)
			case "BYF":
				//fmt.Println("Enter BYF...")
				byValue, err := strconv.ParseUint(attr[1].Value, 10, 8)
				if err != nil {
					panic("Convert string to uint8 err" + err.Error())
				}
				objExt := ObjectExt{NType: BYTE_TYPE, ByValue: byte(byValue)}
				c.putItem(attr[0].Value, objExt)
			case "SF":
				//fmt.Println("Enter SF...")
				sValue, err := strconv.ParseInt(attr[1].Value, 10, 16)
				if err != nil {
					panic("Convert string to short err" + err.Error())
				}
				objExt := ObjectExt{NType: SHORT_TYPE, ShValue: int16(sValue)}
				c.putItem(attr[0].Value, objExt)
			case "NF":
				//fmt.Println("Enter NF...")
				nValue, err := strconv.Atoi(attr[1].Value)
				if err != nil {
					panic("Convert string to int err" + err.Error())
				}
				objExt := ObjectExt{NType: INTEGER_TYPE, NValue: nValue}
				c.putItem(attr[0].Value, objExt)
			case "LF":
				//fmt.Println("Enter LF...")
				lValue, err := strconv.ParseInt(attr[1].Value, 10, 64)
				if err != nil {
					panic("Convert string to long err" + err.Error())
				}
				objExt := ObjectExt{NType: LONG_TYPE, LValue: lValue}
				c.putItem(attr[0].Value, objExt)
			case "FF":
				//fmt.Println("Enter FF...")
				fValue, err := strconv.ParseFloat(attr[1].Value, 32)
				if err != nil {
					panic("Convert string to float32 err" + err.Error())
				}
				objExt := ObjectExt{NType: FLOAT_TYPE, FValue: float32(fValue)}
				c.putItem(attr[0].Value, objExt)
			case "DBLF":
				//fmt.Println("Enter DBLF...")
				dValue, err := strconv.ParseFloat(attr[1].Value, 64)
				if err != nil {
					panic("Convert string to double err" + err.Error())
				}
				objExt := ObjectExt{NType: DOUBLE_TYPE, DblValue: dValue}
				c.putItem(attr[0].Value, objExt)
			case "STRF":
				//fmt.Println("Enter STRF...")
				objExt := ObjectExt{NType: STRING_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
				//fmt.Println("nType:", STRING_TYPE, " strValue:", attr[1].Value)
			case "DF":
				//fmt.Println("Enter DF...")
				objExt := ObjectExt{NType: DATE_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
			case "TF":
				//fmt.Println("Enter TF...")
				objExt := ObjectExt{NType: TIME_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
			case "DTF":
				//fmt.Println("Enter DTF...")
				objExt := ObjectExt{NType: DATETIME_TYPE, StrValue: attr[1].Value}
				c.putItem(attr[0].Value, objExt)
			case "CDOF":
				cdoValue := NewCDO()
				cdoValue.ProcessXML(decoder)
				objExt := ObjectExt{NType: CDO_TYPE, CdoValue: cdoValue}
				c.putItem(attr[0].Value, objExt)
				//fmt.Println("Enter CDOF...")
			case "BAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				bsValue := make([]bool,len(strValueArr))
				for i:=0; i< len(strValueArr);i++{
					if strings.ToLower(strValueArr[i]) == "false" {
						bsValue[i] = false
					}else if strings.ToLower(strValueArr[i]) == "true"{
						bsValue[i] = true
					}
				}
				objExt := ObjectExt{NType:BOOLEAN_ARRAY_TYPE,BsValue:bsValue}
				c.putItem(attr[0].Value,objExt)
			case "BYAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				bysValue := make([]byte,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					byValue, err := strconv.ParseUint(strValueArr[i], 10, 8)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					bysValue[i] = byte(byValue)
				}
				objExt := ObjectExt{NType:BYTE_ARRAY_TYPE,BysValue:bysValue}
				c.putItem(attr[0].Value,objExt)
			case "SAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				shsValue := make([]int16,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					bValue, err := strconv.ParseInt(strValueArr[i], 10, 16)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					shsValue[i] = int16(bValue)
				}
				objExt := ObjectExt{NType:SHORT_ARRAY_TYPE,ShsValue:shsValue}
				c.putItem(attr[0].Value,objExt)
			case "NAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				nsValue := make([]int,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					bValue, err := strconv.Atoi(strValueArr[i])
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					nsValue[i] = bValue
				}
				objExt := ObjectExt{NType:INTEGER_ARRAY_TYPE,NsValue:nsValue}
				c.putItem(attr[0].Value,objExt)
			case "LAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				lsValue := make([]int64,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					lValue, err := strconv.ParseInt(strValueArr[i], 10, 64)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					lsValue[i] = lValue
				}
				objExt := ObjectExt{NType:LONG_ARRAY_TYPE,LsValue:lsValue}
				c.putItem(attr[0].Value,objExt)
			case "FAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				fsValue := make([]float32,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					fValue, err := strconv.ParseFloat(strValueArr[i], 32)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					fsValue[i] = float32(fValue)
				}
				objExt := ObjectExt{NType:FLOAT_ARRAY_TYPE,FsValue:fsValue}
				c.putItem(attr[0].Value,objExt)
			case "DBLAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				dbsValue := make([]float64,len(strValueArr))
				for i := 0; i<len(strValueArr); i++ {
					dValue, err := strconv.ParseFloat(strValueArr[i], 64)
					if err != nil {
						fmt.Println("parseUint err:",err.Error())
					}
					dbsValue[i] = dValue
				}
				objExt := ObjectExt{NType:DOUBLE_ARRAY_TYPE,DblsValue:dbsValue}
				c.putItem(attr[0].Value,objExt)
			case "STRAF":
				strAfName = attr[0].Value
				isStrAF = true
				fmt.Println("STRAF",attr[0].Value)
			case "STR":
				isStr = true
			case "DAF":fallthrough
			case "TAF":fallthrough
			case "DTAF":
				strValueArr := []string{}
				if attr[1].Value == ""{
					strValueArr = []string{""}
				}else{
					strValueArr = strings.Split(attr[1].Value,",")
				}
				dataType := NONE_TYPE
				if strTag == "DAF"{
					dataType = DATE_ARRAY_TYPE
				}else if strTag == "TAF"{
					dataType = TIME_ARRAY_TYPE
				}else{
					dataType = DATETIME_ARRAY_TYPE
				}
				objExt := ObjectExt{NType:dataType,StrsValue:strValueArr}
				c.putItem(attr[0].Value,objExt)
			case "CDO":
				if isCdoAf {
					isCdo = true
					cdoValue := NewCDO()
					cdoValue.ProcessXMLCDOAF(decoder)
					cdosValue = append(cdosValue,cdoValue)
				}
			case "CDOAF":
				strCdoAfName = attr[0].Value
				isCdoAf = true
			}
		case xml.EndElement:
			//fmt.Printf("CDOAF Token of '%s' end\n", token.Name.Local)
			if token.Name.Local == "CDO" || (token.Name.Local == "CDO" && isCdo == true && isCdoAf == false) {
				//fmt.Println("xml.EndElement")
				isCdo = false
				return
			}
			if token.Name.Local == "CDOAF" {
				objExt := ObjectExt{NType: CDO_ARRAY_TYPE, CdosValue: cdosValue}
				//fmt.Println("cdoaf name:",strCdoAfName,"cdosValue:",objExt.CdosValue)
				c.putItem(strCdoAfName, objExt)
				isCdoAf = false
				strCdoAfName = ""
			}
			if token.Name.Local == "STRAF" {
				objExt := ObjectExt{NType: STRING_ARRAY_TYPE, StrsValue: strsValue}
				//fmt.Println("strsValue:",objExt.StrsValue)
				//fmt.Println("STRAF strsValue:",strsValue)
				c.putItem(strAfName, objExt)
				strAfName = ""
				strsValue = []string{""}
				isStrAF = false
				//isStr = false
			}
			if isStr{
				isStr = false
			}
		case xml.CharData:
			content := string([]byte(token))
			//fmt.Printf("This is the content: %v\n", content)
			if isStrAF && isStr {
				//strContent := "<STR>"+content+"</STR>"
				strsValue = append(strsValue, content)
				//fmt.Println("strsValue:",strsValue)
			}
		default:
			//fmt.Println("xml decode err:",token)
		}
	}
}

func (c *CDO) FromXML(strXML string) CDO {
	inputReader := strings.NewReader(strXML)
	decoder := xml.NewDecoder(inputReader)
	c.ProcessXML(decoder)
	return *c
}

func (c *CDO) ToXML() string {
	strbXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	c.toXMLWithIndent("", &strbXML)
	return strbXML
}

func (c *CDO) ToXMLWithStr(strbXML string) string {
	strbXML += "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	c.toXMLWithIndent("", &strbXML)
	return strbXML
}

func (c *CDO) outputField(strFieldId string, ext ObjectExt, strIndent string, strbXML *string) {
	if strIndent != "" {
		*strbXML += strIndent
	}
	switch ext.NType {
	case BOOLEAN_TYPE:
		//strbXML += "<BF N=\""+strFieldId+"\" V=\""+strconv.FormatBool(ext.GetBooleanValue())+"\"/>"
		*strbXML += fmt.Sprintf("<BF N=\"%s\" V=\"%t\"/>", strFieldId, ext.GetBooleanValue())
	case BYTE_TYPE:
		*strbXML += fmt.Sprintf("<BYF N=\"%s\" V=\"%s\"/>", strFieldId, strconv.FormatUint(uint64(ext.GetByteValue()),10))
	case SHORT_TYPE:
		*strbXML += fmt.Sprintf("<SF N=\"%s\" V=\"%s\"/>", strFieldId, strconv.Itoa(int(ext.GetShortValue())))
	case INTEGER_TYPE:
		*strbXML += fmt.Sprintf("<NF N=\"%s\" V=\"%s\"/>", strFieldId, strconv.Itoa(ext.GetIntegerValue()))
	case LONG_TYPE:
		*strbXML += fmt.Sprintf("<LF N=\"%s\" V=\"%s\"/>", strFieldId, strconv.FormatInt(ext.GetLongValue(), 10))
	case FLOAT_TYPE:
		*strbXML += fmt.Sprintf("<FF N=\"%s\" V=\"%s\"/>", strFieldId, strconv.FormatFloat(float64(ext.GetFloatValue()), 'f', -1, 32))
	case DOUBLE_TYPE:
		*strbXML += fmt.Sprintf("<DBLF N=\"%s\" V=\"%s\"/>", strFieldId, strconv.FormatFloat(ext.GetDoubleValue(), 'f', -1, 64))
	case STRING_TYPE:
		*strbXML += fmt.Sprintf("<STRF N=\"%s\" V=\"%s\"/>", strFieldId, ext.GetStringValue())
	case DATE_TYPE:
		*strbXML += fmt.Sprintf("<DF N=\"%s\" V=\"%s\"/>", strFieldId, ext.GetDateValue())
	case TIME_TYPE:
		*strbXML += fmt.Sprintf("<TF N=\"%s\" V=\"%s\"/>", strFieldId, ext.GetTimeValue())
	case DATETIME_TYPE:
		*strbXML += fmt.Sprintf("<DTF N=\"%s\" V=\"%s\"/>", strFieldId, ext.GetDateTimeValue())
	case CDO_TYPE:
		cdo := ext.GetValue().(CDO)
		if strIndent != "" {
			*strbXML += fmt.Sprintf("<CDOF N=\"%s\"/>\r\n", strFieldId)
			cdo.toXMLWithIndent(strIndent+"\t", strbXML)
			*strbXML += strIndent
		} else {
			*strbXML += fmt.Sprintf("<CDOF N=\"%s\">", strFieldId)
			cdo.toXMLWithIndent("", strbXML)
		}
		*strbXML += "</CDOF>"
	case BOOLEAN_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<BAF N=\"%s\" V=\"", strFieldId)
		bsValue := ext.GetBooleanArrayValue()
		for i := 0; i < len(bsValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strconv.FormatBool(bsValue[i])
		}
		*strbXML += "\"/>"
	case BYTE_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<BYAF N=\"%s\" V=\"", strFieldId)
		bysValue := ext.GetByteArrayValue()
		for i := 0; i < len(bysValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strconv.FormatUint(uint64(bysValue[i]),10)
		}
		*strbXML += "\"/>"
	case SHORT_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<SAF N=\"%s\" V=\"", strFieldId)
		bysValue := ext.GetShortArrayValue()
		for i := 0; i < len(bysValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strconv.Itoa(int(bysValue[i]))
		}
		*strbXML += "\"/>"
	case INTEGER_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<NAF N=\"%s\" V=\"", strFieldId)
		bysValue := ext.GetIntegerArrayValue()
		for i := 0; i < len(bysValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strconv.Itoa(bysValue[i])
		}
		*strbXML += "\"/>"
	case LONG_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<LAF N=\"%s\" V=\"", strFieldId)
		bysValue := ext.GetLongArrayValue()
		for i := 0; i < len(bysValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strconv.FormatInt(bysValue[i], 10)
		}
		*strbXML += "\"/>"
	case FLOAT_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<FAF N=\"%s\" V=\"", strFieldId)
		bysValue := ext.GetFloatArrayValue()
		for i := 0; i < len(bysValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strconv.FormatFloat(float64(bysValue[i]), 'f', -1, 32)
		}
		*strbXML += "\"/>"
	case DOUBLE_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<DBLAF N=\"%s\" V=\"", strFieldId)
		bysValue := ext.GetDoubleArrayValue()
		for i := 0; i < len(bysValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strconv.FormatFloat(bysValue[i], 'f', -1, 64)
		}
		*strbXML += "\"/>"
	case STRING_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<STRAF N=\"%s\">", strFieldId)
		strsValue := ext.GetStringArrayValue()
		fmt.Println("strsValue:",strsValue)
		if strIndent == "" {
			for i := 0; i < len(strsValue); i++ {
				*strbXML +=fmt.Sprintf("<STR>%s</STR>",EncodeToXMLText(strsValue[i]))
			}
		} else {
			*strbXML += "\r\n"
			for i := 0; i < len(strsValue); i++ {
				//TODO
				*strbXML += fmt.Sprintf("%s\t<STR>%s</STR>\r\n",strIndent,EncodeToXMLText(strsValue[i]))
			}
			*strbXML += strIndent
		}
		*strbXML += "</STRAF>"
	case DATE_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<DAF N=\"%s\" V=\"", strFieldId)
		strsValue := ext.GetDateArrayValue()
		for i := 0; i < len(strsValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strsValue[i]
		}
		*strbXML += "\"/>"
	case TIME_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<TAF N=\"%s\" V=\"", strFieldId)
		strsValue := ext.GetTimeArrayValue()
		for i := 0; i < len(strsValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strsValue[i]
		}
		*strbXML += "\"/>"
	case DATETIME_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<DTAF N=\"%s\" V=\"", strFieldId)
		strsValue := ext.GetDateTimeArrayValue()
		for i := 0; i < len(strsValue); i++ {
			if i > 0 {
				*strbXML += ","
			}
			*strbXML += strsValue[i]
		}
		*strbXML += "\"/>"
	case CDO_ARRAY_TYPE:
		*strbXML += fmt.Sprintf("<CDOAF N=\"%s\">", strFieldId)
		if strIndent == "" {
			*strbXML += "\r\n"
		}
		cdosValue := ext.GetCDOArrayValue()
		for i := 0; i < len(cdosValue); i++ {
			if strIndent != "" {
				cdosValue[i].toXMLWithIndent(strIndent+"\t", strbXML)
			} else {
				cdosValue[i].toXMLWithIndent("", strbXML)
			}
		}
		if strIndent != "" {
			*strbXML += strIndent + "</CDOAF>"
		} else {
			*strbXML += "</CDOAF>"
		}
	}
	if strIndent != "" {
		*strbXML += "\r\n"
	}
	return
}

func (c *CDO) toXMLWithIndent(strIndent string, strbXML *string) {
	//strbXML := ""
	if strIndent != "" {
		*strbXML += strIndent + "<CDO>\r\n"
	} else {
		*strbXML += "<CDO>"
	}
	for k, v := range c.hmItem {
		if strIndent != "" {
			c.outputField(k, v, strIndent+"\t", strbXML)
		} else {
			c.outputField(k, v, "", strbXML)
		}
	}
	if strIndent != "" {
		*strbXML += strIndent + "</CDO>\r\n"
	} else {
		*strbXML += "</CDO>"
	}
}

func (c *CDO) ToXMLWithIndent() string {
	strbXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n>"
	c.toXMLWithIndent("", &strbXML)
	return strbXML
}

func (c *CDO) SetField(strName string, field ObjectExt) {
	c.putItem(strName, field)
}

func (c *CDO) GetIndexValue(strIndex string, cdoRoot CDO) int {
	nIndex := 0
	strIndex = strings.Trim(strIndex, " ")
	if strIndex[0] >= '0' && strIndex[0] <= '9' { //下载是数字，直接使用
		nIndex, _ = strconv.Atoi(strIndex)
	} else { //下载为字段Id, 获取字段值当索引
		objIndex := cdoRoot.getObject(strIndex)
		switch objIndex.NType {
		case BYTE_TYPE:
			nIndex = int(objIndex.ByValue)
		case SHORT_TYPE:
			nIndex = int(objIndex.ShValue)
		case INTEGER_TYPE:
			nIndex = objIndex.NValue
		case LONG_TYPE:
			nIndex = int(objIndex.LValue)
		default:
			panic("Invalid array index")
		}
	}
	return nIndex
}

// TODO java 中此处该文件以及cdo文件没有调用者，js中有
//func (c *CDO)GetField(strFieldId string)interface{}{
//	// strField 不能为一个数组
//	if strings.HasSuffix(strFieldId,"]") {
//		panic("Invalid FieldId "+strFieldId)
//	}
//	if strings.IndexAny(strFieldId,".") == -1 { //查找某个字符串中某子串第一次出现的位置
//		nArrayStartIndex := strings.IndexAny(strFieldId,"[")
//		nArrayEndIndex   := strings.IndexAny(strFieldId,"]")
//		if nArrayStartIndex < 0 && nArrayEndIndex < 0{
//			value,ok := c.hmItem[strFieldId]
//			if !ok {
//				return nil
//			}
//			return &value
//		}else{
//			if nArrayStartIndex >=0 && nArrayEndIndex > 0 && nArrayEndIndex-nArrayStartIndex>1{
//				panic("FieldId "+strFieldId+" is invalid")
//			}
//		}
//	}
//
//	//FieldId中带有 .
//	nDotIndex := strings.IndexAny(strFieldId,".")
//	nArrayStartIndex := strings.IndexAny(strFieldId,"[")
//	// 没有数组或 . 在 [ 前面，直接查找并返回下级
//	if nArrayStartIndex < 0 || nDotIndex < nArrayStartIndex {
//		field,ok := c.hmItem[strFieldId[0:nDotIndex]]
//		if !ok {
//			return nil
//		}else{
//
//			if reflect.TypeOf(field) != reflect.TypeOf(CDOField{}) {
//				panic("Invalid FieldId "+strFieldId)
//			}
//		}
//		v := reflect.ValueOf(field.GetValue())
//
//	}
//}
//func (c *CDO)GetFieldForJava( strFieldId string )ValueField{
//	objExt := c.getObject( strFieldId )
//	return c.createField(strFieldId, objExt);
//}

//func (c *CDO)SetStringValue1(strFieldId,strValue string){
//	if strings.HasSuffix(strFieldId, "]") { //数组结尾，给数组赋值
//		strIndex := ""
//
//		// 查找 ] 对应的 [
//		nArrayEndIndex := len(strFieldId)-1
//		nArrayStartIndex := FindMatchedChar(nArrayEndIndex,strFieldId)
//		if nArrayStartIndex == -1 { // 没有找到对应的 ]
//			panic("Invalid FieldId"+strFieldId)
//		}
//		field := c.GetFieldValue(strFieldId[0:nArrayStartIndex])
//	}
//}

func (c *CDO) getChildObject(fieldId FieldId, cdoRoot CDO) ObjectExt {
	if fieldId.nType == SIMPLE { //简单类型
		objExt := c.hmItem[strings.ToLower(fieldId.strFieldId)]
		return objExt
	} else if fieldId.nType == MULTISTAGE {
		fieldIdMain := c.ParseFieldId(fieldId.strMainFieldId)
		if fieldIdMain.nType == ARRAY {
			objExt := c.getChildObject(*c.ParseFieldId(fieldIdMain.strMainFieldId), cdoRoot)
			nIndex := c.GetIndexValue(fieldIdMain.strIndexFieldId, cdoRoot)
			temp := objExt.GetValueAt(nIndex).(CDO)
			return temp.getObject(fieldId.strFieldId)
			//return (objExt.GetValueAt(nIndex).(CDO)).getObject(fieldId.strFieldId)
		}
		cdoMainField := c.getChildObject(*fieldIdMain, cdoRoot)
		if binary.Size(cdoMainField) == 0 {
			return ObjectExt{}
		}
		if cdoMainField.IsArrayType() && strings.EqualFold(fieldId.strFieldId, "length") == true {
			return ObjectExt{NType: INTEGER_TYPE, NValue: cdoMainField.GetLength()}
		}
		temp := cdoMainField.GetValue()
		return temp.(CDO).hmItem[strings.ToLower(fieldId.strFieldId)]
		//return (cdoMainField.GetValue().(CDO)).hmItem[strings.ToLower(fieldId.strFieldId)]
	} else { //数组元素
		nIndex := c.GetIndexValue(fieldId.strIndexFieldId, cdoRoot)
		fieldIdMain := c.ParseFieldId(fieldId.strMainFieldId)
		objExt := c.getChildObject(*fieldIdMain, cdoRoot)
		return objExt.GetValueAtExt(nIndex)
	}
}
func (c *CDO) getObject(strFieldId string) ObjectExt {
	objExt, ok := c.hmItem[strings.ToLower(strFieldId)]
	if ok {
		return objExt
	}
	fieldId := c.ParseFieldId(strFieldId)
	objExt = c.getChildObject(*fieldId, *c)
	if binary.Size(objExt) == 0 {
		panic("FieldId" + strFieldId + " not exists")
	}
	return objExt
}
func (c *CDO) getCDOValue(strFieldId string) CDO {
	objExt := c.getObject(strFieldId)
	return objExt.GetCDOValue()
}

func (c *CDO) SetObjectValue(fieldId FieldId, nType int, object interface{}, cdoRoot CDO) {
	if fieldId.nType == SIMPLE {
		obj := NewObjectExt(nType, object)
		c.putItem(fieldId.strFieldId, obj)
	} else if fieldId.nType == MULTISTAGE {
		cdoMain := c.getCDOValue(fieldId.strMainFieldId)
		cdoMain.SetField(fieldId.strFieldId,NewObjectExt(nType,object))
	} else {
		fieldIdMain := c.ParseFieldId(fieldId.strMainFieldId)
		nIndex := -1
		arrField := c.getChildObject(*fieldIdMain, *c)
		if binary.Size(arrField) == 0 {
			panic("FieldId " + fieldId.strMainFieldId + " not exist")
		}
		nIndex = c.GetIndexValue(fieldId.strIndexFieldId, cdoRoot)
		arrField.SetValueAt(nIndex, object)
	}
}
func (c *CDO) SetBooleanValue(strField string, bValue bool) {
	fieldId := c.ParseFieldId(strField)
	c.SetObjectValue(*fieldId, BOOLEAN_TYPE, bValue, *c)
}
func (c *CDO) SetByteValue(strFieldId string, value byte) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, BYTE_TYPE, value, *c)
}
func (c *CDO) SetShortValue(strFieldId string, value int16) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, SHORT_TYPE, value, *c)
}
func (c *CDO) SetIntegerValue(strFieldId string, value int) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, INTEGER_TYPE, value, *c)
}
func (c *CDO) SetLongValue(strFieldId string, value int64) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, LONG_TYPE, value, *c)
}
func (c *CDO) SetFloatValue(strFieldId string, value float32) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, FLOAT_TYPE, value, *c)
}
func (c *CDO) SetDoubleValue(strFieldId string, value float64) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, DOUBLE_TYPE, value, *c)
}
func (c *CDO) SetStringValue(strFieldId string, value string) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, STRING_TYPE, value, *c)
}
func (c *CDO) SetDateValue(strFieldId string, value string) {
	fieldId := c.ParseFieldId(strFieldId)
	reg := regexp.MustCompile(`[0-9]{4}-0[1-9]|1[0-2]-0[1-9]|[1-2][0-9]|3[0-1]`) //考虑加入平年闰年
	if isMatch := reg.MatchString(value);isMatch ==false {
		str := fmt.Sprintf("Invalid Time format:%s", value)
		panic(str)
	}
	c.SetObjectValue(*fieldId, DATE_TYPE, value, *c)
}
func (c *CDO) SetTimeValue(strFieldId string, value string) {
	fieldId := c.ParseFieldId(strFieldId)
	reg := regexp.MustCompile(`[0-2][0-3](:[0-5][0-9]){2}`)
	if isMatch := reg.MatchString(value);isMatch ==false {
		str := fmt.Sprintf("Invalid Time format:%s", value)
		panic(str)
	}
	c.SetObjectValue(*fieldId, TIME_TYPE, value, *c)
}
func (c *CDO) SetDateTimeValue(strFieldId string, value string) {
	fieldId := c.ParseFieldId(strFieldId)
	reg := regexp.MustCompile(`[0-9]{4}-0[1-9]|1[0-2]-0[1-9]|[1-2][0-9]|3[0-1] [0-1][0-9]|[20-23](:[0-5][0-9]){2}`) //考虑加入平年闰年
	if isMatch := reg.MatchString(value);isMatch ==false {
		str := fmt.Sprintf("Invalid Time format:%s", value)
		panic(str)
	}
	c.SetObjectValue(*fieldId, DATETIME_TYPE, value, *c)
}
func (c *CDO) SetCDOValue(strFieldId string, value CDO) {
	fieldId := c.ParseFieldId(strFieldId)
	c.SetObjectValue(*fieldId, CDO_TYPE, value, *c)
}

func (c *CDO) SetBoolArrayValue(strFieldId string, value []bool) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, BOOLEAN_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetByteArrayValue(strFieldId string, value []byte) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, BYTE_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetShortArrayValue(strFieldId string, value []int16) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, SHORT_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetIntegerArrayValue(strFieldId string, value []int) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, INTEGER_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetFloatArrayValue(strFieldId string, value []float32) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, FLOAT_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetDoubleArrayValue(strFieldId string, value []float64) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, DOUBLE_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetLongArrayValue(strFieldId string, value []int64) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, LONG_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetStringArrayValue(strFieldId string, value []string) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, STRING_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetDateArrayValue(strFieldId string, value []string) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, DATE_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetTimeArrayValue(strFieldId string, value []string) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, TIME_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetDateTimeArrayValue(strFieldId string, value []string) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, DATETIME_ARRAY_TYPE, value, *c)
	}
}
func (c *CDO) SetCDOArrayValue(strFieldId string, value []CDO) {
	fieldId := c.ParseFieldId(strFieldId)
	if fieldId.nType == ARRAY {
		panic("INvalid FieldId " + strFieldId)
	} else {
		c.SetObjectValue(*fieldId, CDO_ARRAY_TYPE, value, *c)
	}
}

func (c *CDO) GetBooleanValue(strFieldId string) bool {
	objExt := c.getObject(strFieldId)
	return objExt.GetBooleanValue()
}
func (c *CDO) GetByteValue(strFieldId string) byte {
	objExt := c.getObject(strFieldId)
	return objExt.GetByteValue()
}
func (c *CDO) GetShortValue(strFieldId string) int16 {
	objExt := c.getObject(strFieldId)
	return objExt.GetShortValue()
}
func (c *CDO) GetIntegerValue(strFieldId string) int {
	objExt := c.getObject(strFieldId)
	return objExt.GetIntegerValue()
}
func (c *CDO) GetLongValue(strFieldId string) int64 {
	objExt := c.getObject(strFieldId)
	return objExt.GetLongValue()
}
func (c *CDO) GetFloatValue(strFieldId string) float32 {
	objExt := c.getObject(strFieldId)
	return objExt.GetFloatValue()
}
func (c *CDO) GetDoubleValue(strFieldId string) float64 {
	objExt := c.getObject(strFieldId)
	return objExt.GetDoubleValue()
}
func (c *CDO) GetStringValue(strFieldId string) string {
	objExt := c.getObject(strFieldId)
	return objExt.GetStringValue()
}
func (c *CDO) GetText(strFieldId string) string {
	objExt := c.getObject(strFieldId)
	return objExt.GetStringValue()
}
func (c *CDO) GetDateValue(strFieldId string) string {
	objExt := c.getObject(strFieldId)
	return objExt.GetDateValue()
}
func (c *CDO) GetTimeValue(strFieldId string) string {
	objExt := c.getObject(strFieldId)
	return objExt.GetTimeValue()
}
func (c *CDO) GetDateTimeValue(strFieldId string) string {
	objExt := c.getObject(strFieldId)
	return objExt.GetDateTimeValue()
}
func (c *CDO) GetCDOValue(strFieldId string) CDO {
	objExt := c.getObject(strFieldId)
	return objExt.GetCDOValue()
}

func (c *CDO) Exists(strFieldId string) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	if _, exist := c.hmItem[strings.ToLower(strFieldId)]; exist {
		return exist
	}
	c.getObject(strFieldId)
	return true
}
func (c *CDO) IsEmpty() bool {
	if len(c.hmItem) == 0 {
		return true
	}
	return false
}

func (c *CDO) FieldToJSON(strFieldId string, objExt ObjectExt) string {
	switch objExt.GetType() {
	case BOOLEAN_TYPE:
		vField1 := BooleanField{}
		vField1.SetName(strFieldId)
		vField1.SetType(objExt.GetType())
		vField1.bValue = objExt.GetBooleanValue()
		return vField1.ToJSON()
	case BYTE_TYPE:
		byField := ByteField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetByteValue())
		return byField.ToJSON()
	case SHORT_TYPE:
		byField := ShortField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetShortValue())
		return byField.ToJSON()
	case INTEGER_TYPE:
		byField := IntegerField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetIntegerValue())
		return byField.ToJSON()
	case LONG_TYPE:
		byField := LongField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetLongValue())
		return byField.ToJSON()
	case FLOAT_TYPE:
		byField := FloatField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetFloatValue())
		return byField.ToJSON()
	case DOUBLE_TYPE:
		byField := DoubleField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetDoubleValue())
		return byField.ToJSON()
	case STRING_TYPE:
		byField := StringField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetStringValue())
		return byField.ToJSON()
	case DATE_TYPE:
		byField := DateField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetDateValue())
		return byField.ToJSON()
	case TIME_TYPE:
		byField := TimeField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetTimeValue())
		return byField.ToJSON()
	case DATETIME_TYPE:
		byField := DateTimeField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetDateTimeValue())
		return byField.ToJSON()
	case CDO_TYPE:
		byField := CDOField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetCDOValue())
		return byField.toJSON()
	case BOOLEAN_ARRAY_TYPE:
		byField := BooleanArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetBooleanArrayValue())
		return byField.ToJSON()
	case BYTE_ARRAY_TYPE:
		byField := ByteArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetByteArrayValue())
		return byField.ToJSON()
	case SHORT_ARRAY_TYPE:
		byField := ShortArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetShortArrayValue())
		return byField.ToJSON()
	case INTEGER_ARRAY_TYPE:
		byField := IntegerArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetIntegerArrayValue())
		return byField.ToJSON()
	case LONG_ARRAY_TYPE:
		byField := LongArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetLongArrayValue())
		return byField.ToJSON()
	case FLOAT_ARRAY_TYPE:
		byField := FloatArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetFloatArrayValue())
		return byField.ToJSON()
	case DOUBLE_ARRAY_TYPE:
		byField := DoubleArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetDoubleArrayValue())
		return byField.ToJSON()
	case STRING_ARRAY_TYPE:
		byField := StringArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetStringArrayValue())
		return byField.ToJSON()
	case DATE_ARRAY_TYPE:
		byField := DateArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetDateArrayValue())
		return byField.ToJSON()
	case TIME_ARRAY_TYPE:
		byField := TimeArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetTimeArrayValue())
		return byField.ToJSON()
	case DATETIME_ARRAY_TYPE:
		byField := DateTimeArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetDateTimeArrayValue())
		return byField.ToJSON()
	case CDO_ARRAY_TYPE:
		byField := CDOArrayField{}
		byField.SetName(strFieldId)
		byField.SetType(objExt.GetType())
		byField.SetValue(objExt.GetCDOArrayValue())
		return byField.ToJSON()
	}
	return ""
}
func (c *CDO) createField(strFieldId string, objExt ObjectExt) ValueField {
	vf := ValueField{}
	switch objExt.GetType() {
	case BOOLEAN_TYPE:
		//	vf = ValueField(NewBoolField(strFieldId,objExt.GetBooleanValue()))
		//case BYTE_TYPE:
		//	vf=ByteField(strFieldId,objExt.getByteValue());
		//case SHORT_TYPE:
		//	vf=ShortField(strFieldId,objExt.getShortValue());
		//case INTEGER_TYPE:
		//	vf=IntegerField(strFieldId,objExt.getIntegerValue());
		//case LONG_TYPE:
		//	vf=LongField(strFieldId,objExt.getLongValue());
		//case FLOAT_TYPE:
		//	vf=FloatField(strFieldId,objExt.getFloatValue());
		//case DOUBLE_TYPE:
		//	vf=DoubleField(strFieldId,objExt.getDoubleValue());
		//case STRING_TYPE:
		//	vf=StringField(strFieldId,objExt.getStringValue());
		//case DATE_TYPE:
		//	vf=DateField(strFieldId,objExt.getDateValue());
		//case TIME_TYPE:
		//	vf=TimeField(strFieldId,objExt.getTimeValue());
		//case DATETIME_TYPE:
		//	vf=DateTimeField(strFieldId,objExt.getDateTimeValue());
		//case CDO_TYPE:
		//	vf=CDOField(strFieldId,objExt.getCDOValue());
		//case BOOLEAN_ARRAY_TYPE:
		//	vf=BooleanArrayField(strFieldId,objExt.getBooleanArrayValue());
		//case BYTE_ARRAY_TYPE:
		//	vf=ByteArrayField(strFieldId,objExt.getByteArrayValue());
		//case SHORT_ARRAY_TYPE:
		//	vf=ShortArrayField(strFieldId,objExt.getShortArrayValue());
		//case INTEGER_ARRAY_TYPE:
		//	vf=IntegerArrayField(strFieldId,objExt.getIntegerArrayValue());
		//case LONG_ARRAY_TYPE:
		//	vf=LongArrayField(strFieldId,objExt.getLongArrayValue());
		//case FLOAT_ARRAY_TYPE:
		//	vf=FloatArrayField(strFieldId,objExt.getFloatArrayValue());
		//case DOUBLE_ARRAY_TYPE:
		//	vf=DoubleArrayField(strFieldId,objExt.getDoubleArrayValue());
		//case STRING_ARRAY_TYPE:
		//	vf=StringArrayField(strFieldId,objExt.getStringArrayValue());
		//case DATE_ARRAY_TYPE:
		//	vf=DateArrayField(strFieldId,objExt.getDateArrayValue());
		//case TIME_ARRAY_TYPE:
		//	vf=TimeArrayField(strFieldId,objExt.getTimeArrayValue());
		//case DATETIME_ARRAY_TYPE:
		//	vf=DateTimeArrayField(strFieldId,objExt.getDateTimeArrayValue());
		//case CDO_ARRAY_TYPE:
		//	vf=CDOArrayField(strFieldId,objExt.getCDOArrayValue());
	}
	return vf
}
func (c *CDO) getFieldAt(nIndex int) string {
	strFieldId := c.alFieldId[nIndex]
	objExt := c.alItem[nIndex]
	return c.FieldToJSON(strFieldId, objExt)
}
func (c *CDO) ToJSON() string {
	strJSON := "{"
	nSize := len(c.hmItem)
	for i := 0; i < nSize; i++ {
		if i>0 {
			strJSON += ","
		}
		fieldItem := c.getFieldAt(i)
		strJSON += fieldItem
	}
	lastComma := strings.LastIndex(strJSON, ",")
	lengthJson := len(strJSON)
	if lastComma == lengthJson {
		strJSON = strJSON[:lengthJson-1]
	}
	strJSON += "}"
	return strJSON
}

func NewCDO() CDO {
	return CDO{
		hmItem: make(map[string]ObjectExt),
	}
}
