package idiomSolitaire

/*
	@Title : Calculate
	@Description : 对传入的切片和字母进行单词接龙功能
	@auth : 李江渝
	@param : char string从char字母开头的单词遍历，charSlice []string 存储键盘输入的单词
	@return : 单词接龙，输出以char字母开头单词的单词接龙
*/

import (
"fmt"
)

func IdiomSolitaire(char string,charSlice []string) {
	for i := 0; i < len(charSlice); i++ {
		//遍历切片
		if string(charSlice[i][0]) == char {
			//取出单词的首字母和char比较
			for j := i; j < len(charSlice) - 1; j++ {
				//如果为true，进入for循环
				k := len(charSlice[j]) - 1 		//取出当前字母的的首字母

				/*比较上一个单词和下一个单词的字母是否相同
				如果相同，则输出的当前单词
				*/

				if charSlice[j][k] == charSlice[j+1][0] {
					fmt.Print(" ", charSlice[j])
				}
				/*
					如果不同，则输出当单词
					上面的if语句输出的是当前的单词 而后面的单词要重新判断一次
				*/
				if charSlice[j][k] != charSlice[j+1][0]{
					fmt.Print(" ", charSlice[j])
					//退出当前循环，进入下一个循环，看是否有另外以char单词开头的单词接龙
					break
				}
			}
			//将不同的单词接龙换行输出
			fmt.Println()
		}
	}
}
