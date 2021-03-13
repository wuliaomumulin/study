package main
import(
	pkg "./pkg26"
)
/*
单元测试之TDD
go语言中的测试依赖go test命令。编写测试代码和编写普通的Go代码过程时类似的。并不需要学习新的语法、规则和工具。
go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。
在*_test.go文件中有三种类型的函数:单元测试函数、基准测试函数/示例函数
  类型			格式						作用
测试函数  	函数名前缀为Test			测试程序的一些逻辑行为是否正确
基准函数		函数名前缀为Benchmark		测试函数的性能
示例函数		函数名前缀为Example		为文档提供示例程序

go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。

一、测试函数
	func TestName(t *testing.T){}
	*必须以Test开头，后缀名首字符必须大写
*/
func main(){
	pkg.Test1()
	//pkg.Test2()
	
	//pkg.Test2Server()
	//pkg.Test3()
	

}



