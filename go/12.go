package main
import(
	"fmt"
	"runtime"
	"strings"
	"errors"
)
/*
	函数
	func 函数名(参数)(返回值){函数体}
*	函数的参数和返回值都是可选的

*/
func main(){
	fmt.Println(inttest1(10,20))
	fmt.Println(test2(10,20,15,26,12))
	fmt.Println(test3(21,11,23,14,16,41,33))
	fmt.Println(test4(31,45))
	sum,sub := test5(22,34)
	fmt.Println(sum,sub)
	fmt.Println(test6("This program is distributed in the hope that it will be useful"))
	test7()
	test8()
	test9()
	test10()
	test11()
	test12()
	test13()
	test14()
	test15()
	test16()
	test17()
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
	类型简写
	函数的参数中如果相邻变量的类型相同，则可以省略类型，例如
*/
func inttest1(x,y int) int{
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	return x+y
}
//可变参数是指函数的参数数量固定，通过在参数名后加...来标识
func test2(x ...int) int{
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	fmt.Println(x)//x是一个切片
	sum:=0
	for _,v:=range x{
		sum+=v
	}
	return sum
}
//固定参数搭配可变参数使用时，可变参数要放在固定参数后面
func test3(x int,y ...int)int{
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	fmt.Println(x,y)

	sum :=x
	for _,v := range y{
		sum+=v
	}
	return sum
	//本质上函数的可变参数是通过切片来实现的
}
//多返回值
func test4(x,y int)(int,int){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	return (x+y),(x-y)
}
/*
返回值命名
函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
*/
func test5(x,y int)(sum,sub int){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	sum=x+y
	sub=x-y
	return
}
/*
当我们的一个函数返回值类型时slice时，nil可以看作是一个有效的slice,没必要显示返回一个为0的切片
*/
func test6(x string)[]string{
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	if x=="" {
		return nil//没必要返回[]int{}
	}
	//将返回一个切片
	return strings.Split(x," ")

}
/*
函数进阶
	1、全局变量和局部变量
		a、在if、for、switch结构体中定义的变量只在当前结构体生效;
		b、就近原则，全局变量和局部变量重名，优先访问局部变量
	2、类型函数定义
		type calculation func(int,int)int

*/
type calculation func(int,int)int
func add(x,y int)int{
	return x+y
}
func sub(x,y int)int{
	return x-y
}

func test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var c calculation//声明一个calculation类型的变量
	c = add
	fmt.Printf("type of c:%T\n",c)//type of c:main.calculation
	fmt.Println(c(1,2))//像调用add一样调用c

	f:=sub//将函数sub赋值给变量f
	fmt.Printf("type of f:%T\n",f)//type of f:func(int int) int
	fmt.Println(f(2,2))//像调用sub一样调用f


}
/*
高阶函数
1、函数作为参数
*/
func calc(x,y int,op func(int,int) int) int{
	return op(x,y)
}
//2、函数作为返回值
func do(s string) (func(int,int) int,error){
	switch s{
		case "+":
			return add,nil
		case "-":
			return sub,nil
		default:
			err:=errors.New("error operator")
			return nil,err
	}
}

func test8(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	//1、函数作为参数
	fmt.Println(calc(10,20,add))
	//2、函数作为返回值
	var f,_=do("-")
	fmt.Println(f(1,2))
}
/*
* 函数匿名和闭包
* func(参数)(返回值){函数体}
*/
func test9(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	f:=func(x,y int){
		fmt.Println(x+y)
	}
	f(21,22)//通过变量调用匿名函数

	//自调用匿名函数
	func(x,y int){fmt.Println(x+y)}(20,21)
}
/*
闭包指的是一个函数和与其相关的引用环境而组成的实体.简单来说,闭包=函数+引用环境，例如
*/
func adder() func(int) int{
	var x int
	return func(y int) int{
		x+=y
		return x
	}
}
/**
	闭包1
*/
func test10(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var f = adder()
	for i:=1;i<5;i++ {fmt.Println(f(i))}

	f1 := adder()
	fmt.Println(f1(40));
	fmt.Println(f1(50));
	/*
	 变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。在f的生命周期内，变量x也一直有效。
	*/
}
/**
	闭包2
*/
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string{
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
	return name
	}
}

func test11(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	jpgfunc:=makeSuffixFunc(".jpg")
	txtfunc:=makeSuffixFunc(".txt")
	fmt.Println(jpgfunc("test"))
	fmt.Println(txtfunc("test"))
}
/**
	闭包3
*/
func calc1(base int) (func(int) int,func(int) int) {
	add :=func(i int) int{
		base+=i
		return base
	}

	sub:=func(i int) int{
		base-=i
		return base
	}

	return add,sub
}
func test12(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	add,sub:=calc1(10)

	fmt.Println(add(1),sub(2))//11 9
	fmt.Println(add(3),sub(4))//12 8
	fmt.Println(add(6),sub(5))//14 9

}
/**
 defer会将其后面跟随的语句进行延迟处理,在defer归属的函数即将返回时，将延迟处理语句按defer定义的逆序进行执行。也就是说，先被defer的语句最后被执行，最后被defer的语句最先被执行
 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理，文件关闭、解锁及记录时间
*/
func test13(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	fmt.Println("begin")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
func test14a() int{
	x:=5
	defer func(){
		x++
	}()
	return x
}
func test14b() (x int){
	defer func(){
		x++
	}()
	return x
}
func test14c() (y int){
	x := 5
	defer func(){
		x++
	}()
	return x
}
func test14d() (x int){
	defer func(x int){
		x++
	}(x)
	return 5
}
/**
5
1
5
5
*/
func test14(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	fmt.Println(test14a())
	fmt.Println(test14b())
	fmt.Println(test14c())
	fmt.Println(test14d())
}
func calc2(index string,a,b int) int{
	ret:=a+b
	fmt.Println(index,a,b,ret)
	return ret
}
/**
A 1 2 3
B 10 2 12
BB 10 12 22
AA 1 3 4
*/
func test15(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	x:=1
	y:=2
	defer calc2("AA",x,calc2("A",x,y))
	x=10
	defer calc2("BB",x,calc2("B",x,y))
	y=20
}
/**
	错误处理panic/recover,panic可以在任何地方引发,但recover只有在defer调用的函数中有效。
	revocer()必须搭配defer使用
	defer一定要在可能引发panic的语句之前定义
*/
func test16(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	func(){fmt.Println("func A")}()
	func(){
		defer func(){
			err:=recover()
			//如果程序出现了panic错误，可以通过recover回复过来
			if err!=nil {
				fmt.Println("recover in B")
			}
		}()
		panic("panic in B")
	}()
	func(){fmt.Println("func C")}()
}
/**

*/

var (
	coins = 50
	users = []string{
		"Matthew","Sarah","Augustus","Heidi","Emille","Peter","Giana","Adriano","Aaron","Elizabeth",
	}
	distribution = make(map[string]int,len(users))
)

func test17(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	fmt.Println(dispatchCoin())
}
func dispatchCoin() int{	
	left:=coins
	for _,name := range users {
		e := strings.Count(name,"e") + strings.Count(name,"E")
		i := strings.Count(name,"i") + strings.Count(name,"I")
		o := strings.Count(name,"o") + strings.Count(name,"O")
		u := strings.Count(name,"u") + strings.Count(name,"U")
		sum:=e+i*2+o*3+u*4
		distribution[name] = sum
		left -= sum
	}
	return left
}