package lv1

/*
	@Title : fourInterface
	@Description : 定义接口方法，通过结构体变量来实现接口变量
	@auth : 李江渝
	@param : nil
	@return : nil
*/

import "fmt"

// Dove 声明鸽子接口
type Dove interface {
	GuGuGu()
}

// Repeater 声明复读机接口
type Repeater interface {
	Repeater(string)
}

// LemonExtract 声明了柠檬精接口
type LemonExtract interface {
	LemonExtract()
}

// Weird 声明了真香怪
type Weird interface {
	Weird(string)
}

// Person 声明一个人的结构体
type Person struct {
	Name string
	Age int
	Gender string
}

func (p *Person) GuGuGu() {
	fmt.Printf("%s  又鸽了\n",p.Name)
}

func (p *Person) Repeater(word string) {
	fmt.Printf("%s is a repeater says: %s\n",p.Name,word)
}

func (p *Person) LemonExtract() {
	fmt.Printf("%s  酸死了\n",p.Name)
}

func (p *Person) Weird(sth string) {
	fmt.Printf("%s  对着%s说: 哎，真香\n",p.Name,sth)
}

func Lv1()  {
	//声明一个Person实例 来调用各种Person方法
	var person = &Person{
		Name: "golang",
		Age: 999,
		Gender: "you guess",
	}
	//实例调用结构体方法
	person.GuGuGu()
	person.Repeater("i am a bad duck!")
	person.LemonExtract()
	person.Weird("iphone 13 proMax")
}