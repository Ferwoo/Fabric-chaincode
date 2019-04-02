package bean

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"fmt"
	"time"
	"encoding/json"
)

//合同全集详情
//本条记录主键key由成员ID和合同ID联合组成，具备唯一性

type Compact struct {
	Timestamp          int64    "json:timestamp"         // 本条记录创建时间戳
	UID                string   "json:uid"               //用户唯一ID
	LoanAmount         string   "json:loanAmount"        //用户到款金额
	ApplyDate          string   "json:applyDate"         //申请日期
	CompactStartDate   string   "json:compactStartDate"  //贷款开始日期
	CompactEndDate     string   "json:compactEndDate"    //贷款计划终止日期
	RealEndDate        string   "json:realEndDate"       //贷款实际终止日期

	}

	//贷款操作
	//args：UID，贷款金额、申请日期、贷款开始日期、贷款计划终止日期、合同ID
	//name：成员名称
func Loan(stub shim.ChaincodeStubInterface,args []string ,name string) error  {
	if len(args)!=6{
		return fmt.Errorf("贷款参数格式错误，长度应为6 ")
	}
	if len(args[0])!=32{
		return fmt.Errorf("UID长度错误，应为32位")
	}
	if len(args[2])!=14{
		return fmt.Errorf("贷款申请日期格式错误，应为14位")
	}
	if len(args[3])!=14{
		return fmt.Errorf("贷款开始日期格式错误，应为14位")
	}
	if len(args[4])!=14{
		return fmt.Errorf("贷款计划终止日期格式错误，应为14位")
	}

	var compact Compact
	compact.UID =args[0]
	compact.LoanAmount=args[1]
	compact.ApplyDate=args[2]
	compact.CompactStartDate=args[3]
	compact.CompactEndDate=args[4]
	compact.Timestamp=time.Now().Unix()

	compactJsonBytes ,err:=json.Marshal(&compact)//json序列化
	if err!=nil{
		return fmt.Errorf("Json serialize Compact fail while Loan, compact id ="+args[5])
	}

	//生成合同联合主键,具备唯一性，方便查询
	key,err:=stub.CreateCompositeKey("Compact",[]string{name,args[5]})
	if err!=nil{
		return fmt.Errorf("Failed to CreateCompositeKey while Loan")
	}
	//保存合同的信息
	err=stub.PutState(key,compactJsonBytes)
	if err !=nil{
		return fmt.Errorf("Failed to PutState while Loan,compact id ="+args[5])
	}
	return nil
}
