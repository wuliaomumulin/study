package pkg21

import(
	"fmt"
	"runtime"
	//"encoding/json"//解析json的时候会可见该属性，但是正常在包外还是无法访问
)
//返回调用者函数名称
func Getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
//定义一个接口
type Sayer interface{
	Say()
}
//定义两个结构体
type Dog struct{}
type Cat struct{}
/*
实现两个结构体的方法
接口实现了所有方法，就是实现了这个接口
 */
func (d Dog) Say(){
	fmt.Println("汪汪汪")
}
func (c Cat) Say(){
	fmt.Println("喵喵喵")
}

/*	
	值接收者和指针接收者实现接口的区别
*/
type Mover interface{
	Move()
}
//值接收者
func (d Dog) Move(){
	fmt.Println("狗会动")
}

type Mover1 interface{
	Movep()
}
//指针类型的接收者
func (d *Dog) Movep(){
	fmt.Println("指针类型的狗会动")
}

//dog既可以实现Say接口，也可以实现Move接口
type Dog1 struct{
	Name string
}
type Sayer1 interface{
	Say1()
}
type Mover11 interface{
	Move1()
}
func (d Dog1)Say1(){
	fmt.Printf("%s会叫汪汪汪\n",d.Name)
}
func (d Dog1) Move1(){
	fmt.Printf("%s会动\n",d.Name)
}
//多个类型实现同一接口
type Mover2 interface{
	Move2()
}
type Dog2 struct{
	Name string
}
type Car2 struct{
	//brand string `json:"brand"`//解析json的时候会可见该属性，但是正常在包外还是无法访问
	Brand string
}
func (d Dog2) Move2(){
	fmt.Printf("%s会动\n",d.Name)
}
func (c Car2) Move2(){
	fmt.Printf("%s速度70迈\n",c.Brand)
}
//并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或结构体来实现
//洗衣机
type washingMachine interface{
	wash()
	dry()
}
//甩干器
type dryer struct{}
//实现甩干器的dry方法
func (d dryer) dry(){
	fmt.Println("甩一甩")
}
//海尔洗衣机
type haier struct{
	dryer //嵌入甩干器
}
//实现washingMachine接口的wash方法
func (h haier) wash(){
	fmt.Println("洗一洗")
}
/*
接口嵌套
*/
type Animal interface{
	Sayer1
	Mover11
}
/*
使用空接口实现可以接收任意类型的函数参数
*/
func show(a interface{}){
	fmt.Printf("type:%T,value:%v\n",a,a)
}
//空接口作为map的值
//使用空接口实现可以保存任意值的字典
func Test8(){
	var studentInfo=make(map[string]interface{})
	studentInfo["name"] = "张三"
	studentInfo["age"] = 20
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}