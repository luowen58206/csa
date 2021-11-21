package main

/*  @Title : main.go
@Description : CSA-Homework
@Author : 李江渝
@Update : 2021/11/7
*/

import (
	"example.com/m/HomeworkCsa/First/idiomSolitaire"
	"example.com/m/HomeworkCsa/First/testCal"

	"fmt"
)

func main()  {
	//@Homework01
	//定义两个操作数
	var n1,n2 float64
	//定义一个操作符
	var operation string

	//从终端读取两个操作数及操作符并分别赋值给之前定义的变量
	fmt.Println("input n1 , n2 ,operation seq ='space' ")
	fmt.Scanf("%f %f %s",&n1,&n2,&operation)

	//调用函数得到操作后的结果并打印
	result := testCal.Calculate(n1,n2,operation)
	fmt.Printf("%.2f %s %.2f = %.4f", n1, operation, n2, result)
	fmt.Println()

	//@Homework02
	//定义一个string切片来存储单词
	var charSlice []string
	//定义一个变量来接收键盘的输入
	var newChar string
	//定义一个字母
	var char string

	//以空格间隔，将键盘输入的单词存储到切片
	fmt.Println("Please input some word : (seq = '\\n')")
	for  newChar != "nil" {
		fmt.Scanf("%s ",&newChar)
		charSlice = append(charSlice, newChar)
	}

	//存储读入的字母，从以这个字母开头的单词遍历
	fmt.Println("Please input a char : ")
	fmt.Scanln(&char)

	//传入参数
	idiomSolitaire.IdiomSolitaire(char,charSlice)
}
