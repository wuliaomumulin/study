package main
import(
	"fmt"
)

/*
* 定义了const
*   linux权限就是八进制数，0664
*   0x开头是16进制数
*   go不直接定义二进制数
*/
func main(){
	var i1 = 101
	fmt.Printf("%d\n",i1)
	fmt.Printf("%o\n",i1)
	fmt.Printf("%x\n",i1)
	//八进制
	i2 := 077
	fmt.Printf("%d\n",i2)
	//math.MaxFloat32();
	var f1 = 1.23456
	fmt.Printf("%T\n",f1)/*输出格式，默认go语言中的小数都为float64*/
	fmt.Printf("%T\n",float32(1.23456))/*显示声明32位,float64与float32不能强制转换*/
	/* bool值位孤立类型  T type ,v value */
	b1 := true
	var b2 bool
	fmt.Printf("%T\n",b1)
	fmt.Printf("%T value:%v\n",b2,b2)

	fmt.Printf("%b\n",b1);//错误示范
	fmt.Printf("hello,%#v\n","world");//在值前面加一个#号相当于值获得描述符，如果值是字符串，那么会给值用""包起来

	/*
		fmt占位符
		%T 查看类型
		%d 十进制 %b 二进制 %o 八进制 %x 十六进制
		%s 字符串 
		%v 值
	*/
}