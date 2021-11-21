package testCal

/*
	@Title : Calculate
	@Description : 对传入的两个数值进行operation操作
	@auth : 李江渝
	@param : n1,n2,operation n1,n2为操作数，operation为操作符
	@return : res float64  保存上面变量之间操作的结果
*/
import (
	"fmt"
	"math"
)

func Calculate(n1,n2 float64, operation string) (res float64) {

	//对operation进行判断，执行符号相应的功能
	switch operation {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	case "sqrt":
		fmt.Println("please neglect n2")
		res = math.Sqrt(n1)
	case "pow":
		res = math.Pow(n1, n2)
	case "exit":
		break
	default:
		fmt.Println("input a current operation !")
	}

	//返回结果
	return res
}

