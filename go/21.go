package main
import(
//一、在包内引入的包，如果要在包外使用还需要再一次引入
	"fmt"
	"./pkg21"
)
/*
接口
interface定义了一个对象的行为规范，只定义不实现，由具体的对象来实现规范的细节
接口的定义
	type 接口类型名 interface{
		方法名1(参数列表1) 返回值列表1
		方法名2(参数列表2) 返回值列表2
	}
	*接口名：使用type将接口定义为自定义的类型名。在GO语言的接口在命名时，一般会在单词后面加er,如有写操作的接口叫writer,有字符串的接口叫Stringer等。接口名最好能突出该接口的类型定义。
	*方法名：当方法名首字母大写且接口类型首字母也是大写，这个方法可以被接口所在的包(package)之外的代码访问
	*参数列表、返回值列表：参数列表和返回值列表的参数变量名称可以省略;
	example:
		type writer interface{
		 Write([]byte) error
		}
	当看到这个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的Write方法来做一些事情

基于当前路径下的包导入需要需改GO111MODULE参数，GO111MODULE 有三种选项:off、on、auto,默认值是 auto
go env -w GO111MODULE=off
*/
func main(){
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
	test10()
}
/*
	接口实现
 */
func test1(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	var x pkg21.Sayer //声明一个Sayer类型的变量
	a:=pkg21.Cat{}//实例化一个Cat
	b:=pkg21.Dog{}//实例化一个Dog
	x=a//可以把Cat直接赋值给x
	x.Say()
	x=b//亦可以把Dog直接赋值给x
	x.Say()
}
/*
	值接收者和指针接收者实现接口
	从下面的代码我们发现，使用值接收者实现接口之后，不管是Dog结构体还是结构体指针*Dog类型的变量都可以赋值给该接口变量，因为Go语言中有对指针类型变量求值的语法糖,dog2指针内部会自动求值*Dog
 */
func test2(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	//值接收者
	var x pkg21.Mover
	var dog1 = pkg21.Dog{}//dog1是dog类型
	x=dog1
	var dog2 = &pkg21.Dog{}//dog2是*dog指针类型
	x=dog2
	x.Move()

	//指针接收者
	/**
	 * cannot use dog3 (type pkg21.Dog) as type pkg21.Mover1 in assignment:
        pkg21.Dog does not implement pkg21.Mover1 (Movep method has pointer receiver)
	  接口只接收*指针类型，不接收值类型，所以注释63、64两行
	 */
	var a pkg21.Mover1
	//var dog3=pkg21.Dog{}
	//a=dog3//x不可以接收Dog类型
	var dog4=&pkg21.Dog{}
	a=dog4
	a.Movep()
	
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "sb"
	} else {
		talk = "您好"
	}
	return
}
func test3(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
//dog既可以实现Say接口，也可以实现Move接口
func test4(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	//初始化接口
	var x pkg21.Sayer1 //声明一个Sayer类型的变量
	var y pkg21.Mover11
	//初始化结构体
	var a = pkg21.Dog1{Name:"五一"}
	//结构体赋值给接口
	x=a
	y=a
	//调用接口方法
	x.Say1()
	y.Move1()
}
//多个类型实现同一接口
func test5(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	var x pkg21.Mover2
	var a=pkg21.Dog2{Name:"五一"}
	var b=pkg21.Car2{Brand:"大众"}
	x=a
	x.Move2()
	x=b
	x.Move2()
}
/*
接口嵌套
*/
func test6(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	var x pkg21.Animal
	x=pkg21.Dog1{Name:"五一"}
	x.Move1()
	x.Say1()
}
/*
空接口是指没有定义任何方法的接口，因此任何类型都实现了空接口
空接口类型的变量可以存储任意类型的变量
*/
func test7(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	var x interface{}
	x="Hello 铃"
	fmt.Printf("type:%T,value:%v\n",x,x)
	i:=100
	x=i
	fmt.Printf("type:%T,value:%v\n",x,x)
	x=true
	fmt.Printf("type:%T,value:%v\n",x,x)
}
//空接口作为map的值
//使用空接口实现可以保存任意值的字典
func test8(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	pkg21.Test8()
}
/*类型断言
空接口可以保存任意类型的值，那我们如何获取其存储的具体数据呢
一个接口值是由一个具体类型和具体类型的值两部分组成的，这两部分分别称为接口的动态类型和动态值。

想要判断空接口的值这个时候可以使用类型断言:
x.(T)
x:表示类型为interface{}的变量
T:表示断言，x可能是的类型
该语法返回两个参数，第一个参数时x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，若为false则为断言失败。
*/
func test9(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	var x interface{}
	x="Hello 林"
	v,ok:=x.(string)
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("类型断言失败")
	}
}
/*
Switch方式的断言
*/
func justifyType(x interface{}){
	switch v:=x.(type){
	case string:
		fmt.Println("x is a string,value is ",v)
	case int:
		fmt.Println("x is a int,value is ",v)
	case bool:
		fmt.Println("x is a bool,value is ",v)
	default:
		fmt.Println("unsupport type!")
	}
}
func test10(){
	fmt.Println("---------",pkg21.Getfunctionname(),"-------------")//打印函数名称
	var x interface{}
	x="Hello 林"
	justifyType(x)
}
/*
只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口，不要为了写接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
*/