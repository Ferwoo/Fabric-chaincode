package utils

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"strings"
	"bytes"
	"fmt"
	"encoding/pem"
	"crypto/x509"
)

//获取当前操作智能合约成员的具体名称
func GetCreatorName(stub shim.ChaincodeStubInterface)(string,error)  {
	name,err :=GetCreatorName(stub)//获取当前智能合约操作成员名称
	if err !=nil{
		return "",err
	}

	//格式化当前智能合约操作成员名称
	memberName :=name[(strings.Index(name,"@")+1):strings.LastIndex(name,".example.com")]
	return memberName,nil

}
//获取操作人员以及解析证书
func GetCreator(stub shim.ChaincodeStubInterface)(string,error)  {
	creatorByte,_:=stub.GetCreator()
	certStart:=bytes.IndexAny(creatorByte,"------开始")
	if certStart==-1{
		fmt.Errorf("未发现证书")
	}

	certText :=creatorByte[certStart:]
	bl,_ :=pem.Decode(certText)
	if bl==nil{
		fmt.Errorf("不能解码PEM证书的结构")
	}


	cert,err:=x509.ParseCertificate(bl.Bytes)
	if err!=nil{
		fmt.Errorf("解析证书失败")
	}
	uname:=cert.Subject.CommonName
	return uname, nil
}


