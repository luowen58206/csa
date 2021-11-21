package lv1

/*
	@Title : lv3
	@Description : 使用channel代替加锁操作
	@auth : 李江渝
	@param : nil
	@return : nil
*/

import (
	"fmt"
)

//定义map类型的管道变量来存放阶乘的值
var (
	myRes = make(chan map[int]int, 20)
)

/*
@title: 求阶乘的函数
@param: n n的阶乘
@return: nil
 */
func factorial(n int)  {
	//定义一个map[int]int来接收阶乘函数
	var res = make(map[int]int,20)
	res[n] = 1
	for i := 1; i <= n; i++ {
		res[n] *= i
	}
	//将阶乘的值写入管道中
	myRes <- res
}

func Lv1()  {
	//求1~20以内每个数的阶乘
	for i := 1; i <= 20; i++ {
		//启动一个go协程
		go factorial(i)
	}
	for i := 0; i < 20; i++{
		//答应从管道里面取出的值
		fmt.Println(<-myRes)
	}
}