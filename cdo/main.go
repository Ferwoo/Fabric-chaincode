package main

import (
	"cdo/basedata"
	"fmt"
)
func main(){
	cdo := basedata.NewCDO()
	cdo.SetTimeValue("birthTime","23:55:51")
	fmt.Println("birthTime:",cdo.GetTimeValue("birthTime"))
	return
	//cdo.SetStringValue("strName","zhangsan")
	//cdo.SetBooleanValue("Old",true)
	//
	//cdoChild := basedata.NewCDO()
	//cdoChild.SetStringValue("strClass","Computer")
	//cdoChild.SetBooleanValue("Marray",true)
	//
	//cdo.SetCDOValue("StudentInfo",cdoChild)
	//fmt.Println(cdo.ToXMLWithoutPara())
	//tmp,err := cdo.XmlToCDO(cdo.ToXMLWithoutPara())
	//if err != nil {
	//	fmt.Println("err:",err.Error())
	//}
	////fmt.Println("unmarshal:",tmp.BF)
	////fmt.Println("unmarshal:",tmp.STRF)
	//fmt.Println("unmarshal:",tmp)
	//fmt.Printf("unmarshal:%+v\n",tmp.ArrayField)
	////fmt.Println("unmarshal:",tmp.ArrayField[0])
	////input := `<Person><FirstName>Xu</FirstName><LastName>Xinhua</LastName></Person>`
	//fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVVV")
	input := `
	<?xml version="1.0" encoding="UTF-8"?>
	<CDO>
		<CDOF N="studentinfo">
			<CDO>
				<STRF N="strclass" V="Math"/>
				<BF N="marray" V="false"/>
				<CDOF N="classinfo">
					<CDO>
						<STRF N="classname" V="English"/>
						<STRF N="teacher" V="Lucy"/>
						<NF N="age" V="29"/>
					</CDO>
				</CDOF>
			</CDO>
		</CDOF>
		<STRF N="school" V="MIT"/>
			<STRAF N="strarray">
				<STR>123</STR>
				<STR>456</STR>
			</STRAF>
		<STRF N="strname" V="zhangsan"/>
		<BF N="old" V="true"/>
	</CDO>
	`
	cdo.FromXMLDecode(input)
	//tempStr := cdo.GetCDOValue("studentinfo")
	//fmt.Printf("cdo:%+v\n",tempStr.GetStringValue("strclass"))
	fmt.Printf("cdo.ToXML:\n%s\n",cdo.ToXMLWithoutPara())
	fmt.Print("\nExit main...")

}


