package lv2

/*
	@Title : Receiver
	@Description : 对传入的变量进行类型判断
	@auth : 李江渝
	@param : v interface 可以接收任何变量的参数
	@return : nil
*/



import (
	"example.com/m/HomeworkCsa/second02/lv1"
	"fmt"
)

func Receiver(v interface{})  {
	switch v.(type) {
	case int32:
		fmt.Println("this is int32 type")
	case string:
		fmt.Println("this is string type")
	case float64:
		fmt.Println("this is float64 type")
	case lv1.Person:
		fmt.Println("this is Person struct type")
	case map[int]string:
		fmt.Println("this is map type")
	default:
		fmt.Println("no this type")
	}
}
func Lv2()  {
	//定义不同类型的参数变量 交给方法Receiver来判断

	//int类型
	var num int32 = 18
	Receiver(num)

	//float64类型
	var num02  = 18.02
	Receiver(num02)

	//string类型
	var str  = "i am a bad duck"
	Receiver(str)

	//map类型
	var myMap map[int]string
	myMap = make(map[int]string,10)
	Receiver(myMap)

	//Person结构体类型
	var person lv1.Person
	Receiver(person)

}
