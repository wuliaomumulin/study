package main
import(
	"fmt"
	"runtime"
)
/*
	指针
*/
func main(){
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}

func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称


	a:= 10
	b:= &a
	fmt.Printf("a:%d,b:%p\n",a,&a)//a:10,b:0xc00000c0d0
	fmt.Printf("b:%p,type:%T\n",b,b)//b:0xc00000c0d0,type:*int
	fmt.Println(&b)//0xc000006028
}
//指针取值
func test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	//指针取值
	a:=10
	b:=&a//取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n",b)//type of b:*int
	c:=*b//指针取值(根据指针去内存取值)
	fmt.Printf("type of c:%T\n",c)//type of c:int
	fmt.Printf("value of c:%v\n",c)//value of c:10
}
//指针传值
func modify(x int){
	x=100
}

func modify2(x *int){
	*x=100
}

func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	a:=10
	modify(a)
	fmt.Println(a)//10
	modify2(&a)
	fmt.Println(a)//100
}
//func new(Type) *Type
func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	a:=new(int)
	b:=new(bool)
	fmt.Printf("type a is %T\n",a)//type a is *int
	fmt.Printf("type b is %T\n",b)//type b is *bool
	fmt.Println(*a)//0
	fmt.Println(*b)//false
}
//指针类型使用之前，不单单需要声明，还需要初始化，用new()函数来初始化
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	var a *int
	a=new(int)
	*a=10
	fmt.Println(*a)//10

}
/*
make也是用来内存分配的，区别于new,它只用于slice、map和chan的内存创建，而且他返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
func make(t Type,size ...IntegerType) Type
*/
func test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var b map[string]int
	b=make(map[string]int)
	b["lilin"]=100
	fmt.Println(b)
}
/*
new和make的区别:
1、二者都是用来做内存分配的;
2、make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
3、而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
*/