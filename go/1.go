package main
import(
	"fmt"
)
/*
* 定义了const
*/
const pi = 3.1415926
const (
	STATUSOK = 200
	NOTFOUND = 404
)
const (
	n1 = 100
	n2
	n3
)


/**
* 常量计数器 iota
*/

const (
	a1 = iota //重置为0
	a2
	a3 = 100 //插队,会使后面值为该值
	a4 = iota //值为次数，应该3
	_ //丢弃a4
	a6 //还是5
)

/**
* 多个常量申明在一行
*/
const (
	d1,d2 = iota+1,iota+2
	//没加一行常量申明，iota增1
	d3,d4 = iota+1,iota+2

)
/**
* 定义数量级
*/
const (
	_ = iota //扔掉0
	KB = 1 << (10*iota)//先将1转换位0001,然后左边填充10个零,就是0010 0000 0000 转换位10进制就是1024 2的10次方
	MB = 1 << (10*iota)//2的20次方
	TB = 1 << (10*iota)//30次方
	PB = 1 << (10*iota)
)
//010 000 000 000 

func main(){
	fmt.Println("pi",pi)
	fmt.Println("STATUSOK",STATUSOK)
	fmt.Println("NOTFOUND",NOTFOUND)
	fmt.Println("n1:",n1)
	fmt.Println("n2:",n2)
	fmt.Println("n3:",n3)
	fmt.Println("a1:",a1)
	fmt.Println("a2:",a2)
	fmt.Println("a3:",a3)
	fmt.Println("a6:",a6)
	fmt.Println("d1:",d1)
	fmt.Println("d2:",d2)
	fmt.Println("d3:",d3)
	fmt.Println("d4:",d4)
	fmt.Println("KB:",KB)
	fmt.Println("MB:",MB)
	fmt.Println("TB:",TB)
	fmt.Println("PB:",PB)


}