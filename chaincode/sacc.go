package main

//导入chaincode必要的依赖包
import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

//shim.ChaincodeStubInterface用于接收参数(参数内容由传入的参数决定)

//实现初始化init()函数
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	//获取传入的参数
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("错误：传入了错误数量的参数")
	}

	//将初始状态存入账本
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintln("错误：创建失败", args[0]))
	}
	return shim.Success(nil)
}

//调用智能合约
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	//获取被调用的方法名称和参数
	fn, args := stub.GetFunctionAndParameters()
	//根据方法名称执行指定方法
	var result string //定义返回的结果集
	var err error     //定义返回的错误
	if fn == "set" {
		result, err = set(stub, args)
	} else {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	//调用方法成功后将返回的结果集以字节的方式传出
	return shim.Success([]byte(result))
}

//定义你自己要调用的方法(这里给出Invoke中的set,get)
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("错误：希望传入的参数为键值对的形式")
	}

	//如果成功就像初始化成功一样参数入栈
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("错误：设置失败")
	}
	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	//根据键获取值,只需要传入一个参数
	if len(args) != 1 {
		return "", fmt.Errorf("错误：希望传入的参数个数为1")
	}
	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("错误：根据键获取值失败")
	}
	if value == nil {
		return "", fmt.Errorf("错误：未找到与键对应的值")
	}
	return string(value), nil
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Println("错误：启动失败：%s", err)
	}
}

//编译chaincode
//go get -u --tags nopkcs11 github.com/hyperledger/fabric/core/chaincode/shim
//go build --tags nopkcs11
