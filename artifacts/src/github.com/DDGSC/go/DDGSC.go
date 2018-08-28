/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("DDGSC_cc0")

// 企业
type Enterprise struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	//the field tags are needed to keep case from bouncing around
	ID                    string `json:"id"`                  // 企业ID
	Name                  string `json:"name"`                // 企业名称
	Legal_Personality     string `json:"legalPersonality"`    // 法人代表
	Registered_Capital    string `json:"registeredCapital"`   // 注册资本
	Date_of_Establishment string `json:"dateOfEstablishment"` // 成立日期
	Business_Scope        string `json:"businessScope"`       // 营业范围
	Basic_FI_Name         string `json:"basicFIName"`         // 基本开户银行名称
	Basic_FI_Account      string `json:"basicFIAccount"`      // 基本开户银行账号
}

// 金融机构 Financial Institution
type FI struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID                  string `json:"id"`                 // 金融机构ID
	Name                string `json:"name"`               // 金融机构名称
	Address             string `json:"address"`            // 金融机构地址
	Project_Involvement string `json:"projectInvolvement"` // 参与的项目
}

// 项目
type Project struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID           string            `json:"id"`          // 项目ID
	Name         string            `json:"name"`        // 项目名称
	Description  string            `json:"description"` // 项目简介
	Core_Firm    []Enterprise      `json:"coreFirm"`    // 核心企业列表
	Updown_Firm  []Enterprise      `json:"updownFirm"`  // 上下游企业列表
	Progress     map[string]string `json:"progress"`    // 项目进展 (时间+项目进展描述)
	Bid_Info     Bid               `json:"bidInfo"`     // 招标信息
	Winner_FI    FI                `json:"winnerFI"`    // 中标金融机构
	Credit_Limit float64           `json:"creditLimit"` // 授信额度
	Used_Limit   float64           `json:"usedLimit"`   // 已用额度
	Capital_Flow map[string]string `json:"capitalFlow"` // 资金流信息
	Cargo_Flow   map[string]string `json:"cargoFlow"`   // 货物流信息
}

// 尽职调查报告 Due	Diligence Report
type DDR struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Balance_Sheet Balance_Sheet `json:"balanceSheet"` // 资产负债表
	Description   string        `json:"description"`  // 其他描述
}

// 资产负债表
type Balance_Sheet struct {
	LRFS               string   `json:"lrfs"`              // 法人代表家族史 legal representative family history
	Actual_Controllers []string `json:"actualControllers"` // 实际控制人列表
}

// 招标信息
type Bid struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Start_Date   string       `json:"startDate"`   // 发起时间
	End_Date     string       `json:"end_date"`    // 结束时间
	Project      Project      `json:"project"`     // 所属项目
	Involved_FIs []FI         `json:"involvedFIs"` // 参与的金融机构列表
	Offers       map[FI]Offer `json:"offers"`      // 金融机构报价
	Winner_FI    FI           `json:"winnerFI"`    // 中标银行
}

// 金融机构报价
type Offer struct {
	Loan_Amount   int64   `json:"loanAmount"`   // 放款金额
	Interest_Rate float64 `json:"interestRate"` // 利率
}

// DDGSCChainCode example simple Chaincode implementation
type DDGSCChainCode struct {
}

func (t *DDGSCChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	//logger.Info("########### example_cc0 Init ###########")
	//
	//_, args := stub.GetFunctionAndParameters()
	//var A, B string    // Entities
	//var Aval, Bval int // Asset holdings
	//var err error
	//
	//// Initialize the chaincode
	//A = args[0]
	//Aval, err = strconv.Atoi(args[1])
	//if err != nil {
	//	return shim.Error("Expecting integer value for asset holding")
	//}
	//B = args[2]
	//Bval, err = strconv.Atoi(args[3])
	//if err != nil {
	//	return shim.Error("Expecting integer value for asset holding")
	//}
	//logger.Info("Aval = %d, Bval = %d\n", Aval, Bval)
	//
	//// Write the state to the ledger
	//err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	//if err != nil {
	//	return shim.Error(err.Error())
	//}
	//
	//err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	//if err != nil {
	//	return shim.Error(err.Error())
	//}

	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *DDGSCChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### DDGSCChainCode Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	if function == "add" {
		return t.add(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	} else if function == "move" {
		// Deletes an entity from its state
		return t.move(stub, args)
	}

	//logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	//return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
	logger.Errorf("Unknown action, check the first argument, got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, got: %v", args[0]))
}

// add new data to the blockchain
func (t *DDGSCChainCode) add(stubInterface shim.ChaincodeStubInterface, strings []string) pb.Response {
	return shim.Success(nil)
}

func (t *DDGSCChainCode) move(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//// must be an invoke
	//var A, B string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	//var err error
	//
	//if len(args) != 3 {
	//	return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
	//}
	//
	//A = args[0]
	//B = args[1]
	//
	//// Get the state from the ledger
	//// TODO: will be nice to have a GetAllState call to ledger
	//Avalbytes, err := stub.GetState(A)
	//if err != nil {
	//	return shim.Error("Failed to get state")
	//}
	//if Avalbytes == nil {
	//	return shim.Error("Entity not found")
	//}
	//Aval, _ = strconv.Atoi(string(Avalbytes))
	//
	//Bvalbytes, err := stub.GetState(B)
	//if err != nil {
	//	return shim.Error("Failed to get state")
	//}
	//if Bvalbytes == nil {
	//	return shim.Error("Entity not found")
	//}
	//Bval, _ = strconv.Atoi(string(Bvalbytes))
	//
	//// Perform the execution
	//X, err = strconv.Atoi(args[2])
	//if err != nil {
	//	return shim.Error("Invalid transaction amount, expecting a integer value")
	//}
	//Aval = Aval - X
	//Bval = Bval + X
	//logger.Infof("Aval = %d, Bval = %d\n", Aval, Bval)
	//
	//// Write the state back to the ledger
	//err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	//if err != nil {
	//	return shim.Error(err.Error())
	//}
	//
	//err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	//if err != nil {
	//	return shim.Error(err.Error())
	//}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *DDGSCChainCode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//if len(args) != 1 {
	//	return shim.Error("Incorrect number of arguments. Expecting 1")
	//}
	//
	//A := args[0]
	//
	//// Delete the key from the state in ledger
	//err := stub.DelState(A)
	//if err != nil {
	//	return shim.Error("Failed to delete state")
	//}

	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *DDGSCChainCode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//var A string // Entities
	//var err error
	//
	//if len(args) != 1 {
	//	return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	//}
	//
	//A = args[0]
	//
	//// Get the state from the ledger
	//Avalbytes, err := stub.GetState(A)
	//if err != nil {
	//	jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
	//	return shim.Error(jsonResp)
	//}
	//
	//if Avalbytes == nil {
	//	jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
	//	return shim.Error(jsonResp)
	//}
	//
	//jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	//logger.Infof("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(DDGSCChainCode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
