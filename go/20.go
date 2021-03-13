package main
import(
	"fmt"
	"runtime"
	//路经名称，不要带有文件名称
	t1 "./pkg20"
)
/*
包
一、我们可以根据急需要创建自己的包。一个包简单的可以理解为一个存放.go文件的文件夹，该文件夹下面的.go文件都要在代码的第一行添加如下代码，声明该文件归属的包
1、一个文件夹下面直接包含的文件只能归属一个package,同样一个package的文件不能在多个文件夹下
2、包名可以不与文件夹的名称相同，包名不能包含-符号
3、包名为main的包为应用程序的入口包，这种包编译后得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件
4、import导入的包会基于GOPATH/src/路径中计算的

二、单行导入：
	import "包1"
	import "包2"
三、多行导入:
	import (
    	"包1"
    	"包2"
	)
四、设置别名
	import 别名 "包的路径"

五、匿名导入
如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。匿名导入包与其他方式导入的包一样都会被编译到可执行文件中
	import _ "包的路径"

六、init初始化函数
在go语言程序执行时导入包语句会自动出发包内部init函数的调用。需要注意的是:init()函数没有参数也没有返回值.init()函数在程序执行时自动调用执行，不能在代码中主动调用它.
	包中的执行机制:全局声明-->init()-->main()
	*在运行时，被最后导入的包会最先初始化并调用其init函数.

七、cannot find module for path

基于当前路径下的包导入需要需改GO111MODULE参数，GO111MODULE 有三种选项:off、on、auto,默认值是 auto
go env -w GO111MODULE=off

八、windows go env持久化文件位置
C:\Users\administratorn\AppData\Roaming\go
和GOPATH匹配差不多
*/
func main(){
	test1()

}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
如果想在一个包中引用另一个包里的标识符(如变量、常量、类型、函数等)时，该标识符必须是对外可见的(public)。go语言中只需将标识符的首字母大写就可以让标识符对外可见了。
 */
func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	fmt.Println(t1.Add(100,200),t1.Mode)
}