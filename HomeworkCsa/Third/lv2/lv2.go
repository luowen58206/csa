package lv2

/*
	@Title : lv2
	@Description : 开启三个goroutine，顺序执行共routine
	@auth : 李江渝
	@param : chan bool
	@return : nil
*/

import (
	"fmt"
	"sync"
)

//定义全局变量、等待所有的共routine完成在退出程序
var wg = sync.WaitGroup{}

// Lv2 利用channel会阻塞和关闭后仍可继续读取的特性实现顺序执行
func Lv2()  {
		//定义三个bool类型的channel
		chA := make(chan bool,1)
		chB := make(chan bool,1)
		chC := make(chan bool,1)
		//在三个goroutine都在阻塞的情况下，给A的goroutine传入一个参数，使其不在阻塞
		chC <- true

		wg.Add(3)
		go printA(chC,chA)
		go printB(chA,chB)
		go printC(chB,chC)
		wg.Wait()

	}
//打印A
func printA(chC,chA chan bool)  {
	defer wg.Done()
	//每次向本channel中写入bool类型之后，关闭channel 不影响后面想要读取本channel的值
	defer close(chA)
	for i := 0; i < 10; i++ {
		<-chC
		fmt.Println("A")
		chA<-true
	}
}
//打印B
func printB(chA,chB chan bool)  {
	defer wg.Done()
	defer close(chB)
	for i := 0; i < 10; i++ {
		<-chA
		fmt.Println("B")
		chB<-true
	}
}
//打印C
func printC(chB,chC chan bool)  {
	defer wg.Done()
	defer close(chC)
	for i := 0; i < 10; i++ {
		<-chB
		fmt.Println("C")
		chC<-true
	}
}