package pkg20

import "fmt"

//包变量可见性
var a = 100 //首字符小写，只能在包内使用

const Mode=1 //首字符大写，可在包外调用

//首字符小写，只能在当前包内使用
type person struct{
	name string
}
//首字符大写，可在包外使用
func Add(x, y int)int{
	return x+y
}
//首字符小写，只能在包内使用
func age(){
	var Age=18//局部函数变量，只能在当前函数中使用，所以外部包不可见
	fmt.Println(Age)
}
type Student struct{
	Name string //可见包外访问
	class string //可在包内访问
}
type Payer interface{
	init() //仅限包内访问的方法
	Pay() //可在包外访问的方法
}