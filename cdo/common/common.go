package common

import (
)

const (
	SIMPLE     = iota //简单类型
	MULTISTAGE        //多级类型
	ARRAY             //数组元素
)
const (
	NONE_TYPE = iota
	BOOLEAN_TYPE
	BYTE_TYPE
	SHORT_TYPE
	INTEGER_TYPE
	LONG_TYPE
	FLOAT_TYPE
	DOUBLE_TYPE
	STRING_TYPE
	DATE_TYPE
	TIME_TYPE
	DATETIME_TYPE
	CDO_TYPE
	RECORD_TYPE

	BOOLEAN_ARRAY_TYPE = iota - RECORD_TYPE - 1 + 101
	BYTE_ARRAY_TYPE
	SHORT_ARRAY_TYPE
	INTEGER_ARRAY_TYPE
	LONG_ARRAY_TYPE
	FLOAT_ARRAY_TYPE
	DOUBLE_ARRAY_TYPE
	STRING_ARRAY_TYPE
	DATE_ARRAY_TYPE
	TIME_ARRAY_TYPE
	DATETIME_ARRAY_TYPE
	CDO_ARRAY_TYPE
	RECORD_SET_TYPE
)

func FindMatchedChar(nIndex int, strText string) int {
	if nIndex < 0 {
		return -1
	}
	chChar := string(strText[nIndex])
	chFind := ""
	switch chChar {
	case "(":
		chFind = ")"
	case "{":
		chFind = "}"
	case "[":
		chFind = "]"
	case ")":
		chFind = "("
	case "]":
		chFind = "["
	case "}":
		chFind = "{"
	default:
		return -1
	}

	nStartIndex := -1
	nEndIndex := -1
	nCount := 0
	switch chChar {
	case "(":
		fallthrough
	case "[":
		fallthrough
	case "{":
		for i := nIndex + 1; i < len(strText); i++ {
			ch := string(strText[i])
			if ch == chChar {
				nCount++
			} else {
				if ch == chFind {
					if nCount == 0 {
						nEndIndex = i
						break
					} else {
						nCount--
					}
				}
			}
		}
		return nEndIndex
	case ")":
		fallthrough
	case "]":
		fallthrough
	case "}":
		for i := nIndex - 1; i >= 0; i-- {
			ch := string(strText[i])
			if ch == chChar {
				nCount++
			} else {
				if ch == chFind {
					nStartIndex = i
					break
				} else {
					nCount--
				}
			}
		}
		return nStartIndex
	default:
		return -1
	}
}

func EncodeToXMLText(strText string)string{
	strOutput := ""
	length := len(strText)
	for i:=0; i<length;i++{
		ch := string(strText[i])
		switch ch{
		case"&":strOutput+="&amp;"
		case"/":strOutput+="&#47;"
		case"'":strOutput+="&#039;"
		case">":strOutput+="&gt;"
		case"<":strOutput+="&lt;"
		case"\"":strOutput+="&quot;"
		case"\r":strOutput+="&#xd;"
		case"\n":strOutput+="&#xa;"
		default:strOutput+=ch
		}
	}
	return strOutput
}
