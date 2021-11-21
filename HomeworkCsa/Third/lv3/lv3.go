package lv3

/*
	@Title : lv3
	@Description : 计算50000以内的素数
	@auth : 李江渝
	@param : nil
	@return : nil
*/

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// Lv3 本例使用了未优化的埃氏筛法求素数
func Lv3()  {
	start := time.Now()
	//为50000以内的每个值附上一个bool的变量，默认为false
	var countSlice = make([]bool,50000)	//默认值为 false
	var primes []int
	//开启最大处理核数
	runtime.GOMAXPROCS(16)
	wg.Add(1)
	go func() {
		for i := 2; i < 50000; i++ {
			//因为2是素数，除去后面为2的倍数被标记为true的数，下一个3也是素数
			if countSlice[i] == false {
				primes = append(primes,i)

				//将每个已经筛选出来的素数的倍数的bool值定义为true
				for j := i*2 ; j < 50000;j=j+i {
					countSlice[j] = true
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("time = ",time.Now().Sub(start))
	fmt.Println(primes)
}
