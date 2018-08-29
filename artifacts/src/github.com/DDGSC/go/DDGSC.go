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
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
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
	ID                  string   `json:"id"`                 // 金融机构ID
	Name                string   `json:"name"`               // 金融机构名称
	Address             string   `json:"address"`            // 金融机构地址
	Project_Involvement []string `json:"projectInvolvement"` // 参与的项目ID列表
}

// 项目
type Project struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID           string            `json:"id"`          // 项目ID
	Name         string            `json:"name"`        // 项目名称
	Description  string            `json:"description"` // 项目简介
	DDR          string            `json:"ddr"`         // 尽职调查报告ID
	Core_Firm    []string          `json:"coreFirm"`    // 核心企业ID列表
	Updown_Firm  []string          `json:"updownFirm"`  // 上下游企业ID列表
	Progress     map[string]string `json:"progress"`    // 项目进展 (时间+项目进展描述)
	Bid_Info     string            `json:"bidInfo"`     // 招标信息
	Winner_FI    string            `json:"winnerFI"`    // 中标金融机构
	Credit_Limit float64           `json:"creditLimit"` // 授信额度
	Used_Limit   float64           `json:"usedLimit"`   // 已用额度
	Capital_Flow map[string]string `json:"capitalFlow"` // 资金流信息 (时间+信息)
	Cargo_Flow   map[string]string `json:"cargoFlow"`   // 货物流信息 (时间+信息)
}

// 尽职调查报告 Due	Diligence Report
type DDR struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID            string `json:"id"`
	Balance_Sheet string `json:"balanceSheet"` // 资产负债表_ID
	Description   string `json:"description"`  // 其他描述
}

// 资产负债表
type Balance_Sheet struct {
	ID                 string   `json:"id"`
	LRFS               string   `json:"lrfs"`              // 法人代表家族史 legal representative family history
	Actual_Controllers []string `json:"actualControllers"` // 实际控制人列表
}

// 招标信息
type Bid struct {
	//ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	ID           string       `json:"id"`
	Start_Date   string       `json:"startDate"`   // 发起时间
	End_Date     string       `json:"end_date"`    // 结束时间
	Project      string       `json:"project"`     // 所属项目ID
	Involved_FIs []FI         `json:"involvedFIs"` // 参与的金融机构列表
	Offers       map[FI]Offer `json:"offers"`      // 金融机构报价
	Winner_FI    FI           `json:"winnerFI"`    // 中标银行
}

// 金融机构报价
type Offer struct {
	ID            string  `json:"id"`
	Loan_Amount   int64   `json:"loanAmount"`   // 放款金额
	Interest_Rate float64 `json:"interestRate"` // 利率
}

// DDGSCChainCode implementation
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

	//if function == "add" {
	//	return t.add(stub, args)
	//} else if function == "delete" {
	//	// Deletes an entity from its state
	//	//return t.delete(stub, args)
	//} else if function == "query" {
	//	// queries an entity state
	//	return t.query(stub, args)
	//} else if function == "move" {
	//	// Deletes an entity from its state
	//	//return t.move(stub, args)
	//}

	switch function {
	case "addEnterprise":
		return t.addEnterprise(stub, args)
	case "addFI":
		return t.addFI(stub, args)
	case "addProject":
		return t.addProject(stub, args)
	case "addDDR":
		return t.addDDR(stub, args)
	case "addBalanceSheet":
		return t.addBalanceSheet(stub, args)
	case "addBid":
		return t.addBid(stub, args)
	case "addOffer":
		return t.addOffer(stub, args)
	case "queryEnterprise":
		return t.queryEnterprise(stub, args)
	case "queryFI":
		return t.queryFI(stub, args)
	case "queryProject":
		return t.queryProject(stub, args)
	case "queryDDR":
		return t.queryDDR(stub, args)
	case "queryBalanceSheet":
		return t.queryBalanceSheet(stub, args)
	case "queryBid":
		return t.queryBid(stub, args)
	case "queryOffer":
		return t.queryOffer(stub, args)
	default:
		logger.Errorf("Unknown action, check the first argument, got: %v", args[0])
		return shim.Error(fmt.Sprintf("Unknown action, check the first argument, got: %v", args[0]))
	}

	//logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	//return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))

}

//// add new data to the blockchain
//func (t *DDGSCChainCode) add(stubInterface shim.ChaincodeStubInterface, args []string) pb.Response {
//	return shim.Success(nil)
//}
//
//// Query callback representing the query of a chaincode
//func (t *DDGSCChainCode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//	//var A string // Entities
//	//var err error
//	//
//	//if len(args) != 1 {
//	//	return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
//	//}
//	//
//	//A = args[0]
//	//
//	//// Get the state from the ledger
//	//Avalbytes, err := stub.GetState(A)
//	//if err != nil {
//	//	jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
//	//	return shim.Error(jsonResp)
//	//}
//	//
//	//if Avalbytes == nil {
//	//	jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
//	//	return shim.Error(jsonResp)
//	//}
//	//
//	//jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
//	//logger.Infof("Query Response:%s\n", jsonResp)
//	return shim.Success(Avalbytes)
//}
//
//func (t *DDGSCChainCode) move(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//	//// must be an invoke
//	//var A, B string    // Entities
//	//var Aval, Bval int // Asset holdings
//	//var X int          // Transaction value
//	//var err error
//	//
//	//if len(args) != 3 {
//	//	return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
//	//}
//	//
//	//A = args[0]
//	//B = args[1]
//	//
//	//// Get the state from the ledger
//	//// will be nice to have a GetAllState call to ledger
//	//Avalbytes, err := stub.GetState(A)
//	//if err != nil {
//	//	return shim.Error("Failed to get state")
//	//}
//	//if Avalbytes == nil {
//	//	return shim.Error("Entity not found")
//	//}
//	//Aval, _ = strconv.Atoi(string(Avalbytes))
//	//
//	//Bvalbytes, err := stub.GetState(B)
//	//if err != nil {
//	//	return shim.Error("Failed to get state")
//	//}
//	//if Bvalbytes == nil {
//	//	return shim.Error("Entity not found")
//	//}
//	//Bval, _ = strconv.Atoi(string(Bvalbytes))
//	//
//	//// Perform the execution
//	//X, err = strconv.Atoi(args[2])
//	//if err != nil {
//	//	return shim.Error("Invalid transaction amount, expecting a integer value")
//	//}
//	//Aval = Aval - X
//	//Bval = Bval + X
//	//logger.Infof("Aval = %d, Bval = %d\n", Aval, Bval)
//	//
//	//// Write the state back to the ledger
//	//err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
//	//if err != nil {
//	//	return shim.Error(err.Error())
//	//}
//	//
//	//err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
//	//if err != nil {
//	//	return shim.Error(err.Error())
//	//}
//	return shim.Success(nil)
//}
//
//// Deletes an entity from state
//func (t *DDGSCChainCode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//	//if len(args) != 1 {
//	//	return shim.Error("Incorrect number of arguments. Expecting 1")
//	//}
//	//
//	//A := args[0]
//	//
//	//// Delete the key from the state in ledger
//	//err := stub.DelState(A)
//	//if err != nil {
//	//	return shim.Error("Failed to delete state")
//	//}
//
//	return shim.Success(nil)
//}

// todo: 审理各个struct里头的参数，确定返回的信息的格式

// Add
func (t *DDGSCChainCode) addEnterprise(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	/*
		0	ID
		1	Name
		2	Legal_Personality
		3	Registered_Capbital
		4	Date_of_Establishment
		5	Business_Scope
		6	Basic_FI_Name
		7	Basic_FI_Account
	*/

	if len(args) != 8 {
		return shim.Error("Incorrect arguments, please check your arguments")
	}

	ID := args[0]
	Name := args[1]
	Legal_Personality := args[2]
	Registered_Capbital := args[3]
	Date_of_Establishment := args[4]
	Business_Scope := args[5]
	Basic_FI_Name := args[6]
	Basic_FI_Account := args[7]

	IDCheck, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get enterprise: " + err.Error())
	} else if IDCheck != nil {

		fmt.Println("This enterprise already exists.\nID: " + ID + "\nName: " + Name + "\n")
		return shim.Error("This enterprise already exists.\nID: " + ID + "\nName: " + Name + "\n")
	}
	Enterprise := &Enterprise{ID, Name, Legal_Personality, Registered_Capbital, Date_of_Establishment, Business_Scope, Basic_FI_Name, Basic_FI_Account}

	Enterprise_JSON_Byte, err := json.Marshal(Enterprise)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ID, Enterprise_JSON_Byte)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
func (t *DDGSCChainCode) addFI(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
		添加新的FI的时候默认是没有project ID的，所以只保留两个参数，在chaincode上做一次检查
		0	ID
		1	Name
		2	Address
		3	Project_Involvement	空	todo: updateProjectInvolvment
	*/

	if len(args) != 3 {
		return shim.Error("Incorrect arguments, please check your arguments")
	}

	ID := args[0]
	Name := args[1]
	Address := args[2]
	Project_Involvement := []string{}

	IDCheck, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get FI: " + err.Error())
	} else if IDCheck != nil {

		fmt.Println("This FI already exists.\nID: " + ID + "\nName: " + Name + "\n")
		return shim.Error("This FI already exists.\nID: " + ID + "\nName: " + Name + "\n")
	}
	FI := &FI{ID, Name, Address, Project_Involvement}

	FI_JSON_Byte, err := json.Marshal(FI)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ID, FI_JSON_Byte)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *DDGSCChainCode) addProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	/*
		0	ID
		1	Name
		2	Description
		3	DDR	空
		|	Core_Firm	空	todo updateCoreFirm
		|	Updown_Firm todo updateUpdownFirm
		|	Progress	空 todo updateProgress
		|	Bid_Info	空 todo updateBidInfo
		| 	Winner_FI	空 todo updateWinnerFI
		|	Credit_Limit	空 todo updateCreditLimit
		|	Used_Limit	空 todo updateUsedLimit
		|	Capital_Flow	空	todo updateCapitalFlow
		|	Cargo_Flow	空	todo updateCargoFlow
	*/

	if len(args) != 4 {
		return shim.Error("Incorrect arguments, please check your arguments")
	}

	ID := args[0]
	Name := args[1]
	Description := args[2]
	DDR := string("")
	Core_Firm := []string{}
	Updown_Firm := []string{}
	Progress := make(map[string]string)
	Bid_Info := string("")
	Winner_FI := string("")
	Credit_Limit := float64(0)
	Used_Limit := float64(0)
	Capital_Flow := make(map[string]string)
	Cargo_Flow := make(map[string]string)

	IDCheck, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Project: " + err.Error())
	} else if IDCheck != nil {

		fmt.Println("This Project already exists.\nID: " + ID + "\nName: " + Name + "\n")
		return shim.Error("This Project already exists.\nID: " + ID + "\nName: " + Name + "\n")
	}
	Project := &Project{ID, Name, Description, DDR, Core_Firm, Updown_Firm, Progress, Bid_Info, Winner_FI, Credit_Limit, Used_Limit, Capital_Flow, Cargo_Flow}

	Project_JSON_Byte, err := json.Marshal(Project)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ID, Project_JSON_Byte)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *DDGSCChainCode) addDDR(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
		新的尽职调查报告 balance sheet默认为空
		0	ID
		|	Balance_Sheet	空	 todo updateBalanceSheet
		1	Description
	*/

	if len(args) != 2 {
		return shim.Error("Incorrect arguments, please check your arguments")
	}

	ID := args[0]
	Balance_Sheet := string("")
	Description := args[1]

	IDCheck, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get DDR: " + err.Error())
	} else if IDCheck != nil {

		fmt.Println("This DDR already exists.\nID: " + ID + "\n")
		return shim.Error("This DDR already exists.\nID: " + ID + "\n")
	}
	DDR := &DDR{ID, Balance_Sheet, Description}

	DDR_JSON_Byte, err := json.Marshal(DDR)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ID, DDR_JSON_Byte)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *DDGSCChainCode) addBalanceSheet(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	/*
		新的产品负债表 实际控制人列表为空
		0	ID
		1	LRFS
		|	Actual_Controllers 空	todo: updateActualControllers
	*/

	if len(args) != 2 {
		return shim.Error("Incorrect arguments, please check your arguments")
	}

	ID := args[0]
	LRFS := args[1]
	Actual_Controllers := []string{}

	IDCheck, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get BalanceSheet: " + err.Error())
	} else if IDCheck != nil {

		fmt.Println("This BalanceSheet already exists.\nID: " + ID + "\n")
		return shim.Error("This BalanceSheet already exists.\nID: " + ID + "\n")
	}
	Balance_Sheet := &Balance_Sheet{ID, LRFS, Actual_Controllers}

	BalanceSheet_JSON_Byte, err := json.Marshal(Balance_Sheet)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ID, BalanceSheet_JSON_Byte)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
func (t *DDGSCChainCode) addBid(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	/*
		新的招标信息 实际控制人列表为空
		0	ID
		1	Start_Date
		2	End_Date
		3	Project
		|	Involved_FIs 空	todo: updateINVolvedFIs
		|	Offers	空	todo: updateWInner_FI
		|	Winner_FI	空	todo：updateWInner_FI
	*/

	if len(args) != 4 {
		return shim.Error("Incorrect arguments, please check your arguments")
	}

	ID := args[0]
	Start_Date := args[1]
	End_Date := args[2]
	Project := args[3]
	Involved_FIs := []FI{}
	Offers := make(map[FI]Offer)
	Winner_FI := FI{}

	IDCheck, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Bid: " + err.Error())
	} else if IDCheck != nil {

		fmt.Println("This Bid already exists.\nID: " + ID + "\n")
		return shim.Error("This Bid already exists.\nID: " + ID + "\n")
	}
	Bid := &Bid{ID, Start_Date, End_Date, Project, Involved_FIs, Offers, Winner_FI}

	Bid_JSON_Byte, err := json.Marshal(Bid)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ID, Bid_JSON_Byte)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *DDGSCChainCode) addOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
	/*
		新的招标信息 实际控制人列表为空
		0	ID
		1	Loan_Amount
		2	Interest_Rate
	*/

	if len(args) != 3 {
		return shim.Error("Incorrect arguments, please check your arguments")
	}

	ID := args[0]
	Loan_Amount, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}
	Interest_Rate, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	IDCheck, err := stub.GetState(ID)
	if err != nil {
		return shim.Error("Failed to get Offer: " + err.Error())
	} else if IDCheck != nil {

		fmt.Println("This Offer already exists.\nID: " + ID + "\n")
		return shim.Error("This Offer already exists.\nID: " + ID + "\n")
	}
	Offer := &Offer{ID, Loan_Amount, Interest_Rate}

	Offer_JSON_Byte, err := json.Marshal(Offer)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(ID, Offer_JSON_Byte)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Query
func (t *DDGSCChainCode) queryEnterprise(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	ID = args[0]
	valAsbytes, err := stub.GetState(ID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Enterprise does not exist: " + ID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

func (t *DDGSCChainCode) queryFI(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	ID = args[0]
	valAsbytes, err := stub.GetState(ID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"FI does not exist: " + ID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

func (t *DDGSCChainCode) queryProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	ID = args[0]
	valAsbytes, err := stub.GetState(ID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Project does not exist: " + ID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

func (t *DDGSCChainCode) queryDDR(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	ID = args[0]
	valAsbytes, err := stub.GetState(ID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"DDR does not exist: " + ID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

func (t *DDGSCChainCode) queryBalanceSheet(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	ID = args[0]
	valAsbytes, err := stub.GetState(ID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"BalanceSheet does not exist: " + ID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

func (t *DDGSCChainCode) queryBid(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	ID = args[0]
	valAsbytes, err := stub.GetState(ID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Bid does not exist: " + ID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

func (t *DDGSCChainCode) queryOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var ID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	ID = args[0]
	valAsbytes, err := stub.GetState(ID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Offer does not exist: " + ID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

func main() {
	err := shim.Start(new(DDGSCChainCode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
