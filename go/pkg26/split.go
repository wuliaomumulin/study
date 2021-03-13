package pkg26

import(
	//"fmt"
	"runtime"
	"strings"
)
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
一、测试函数
*/
func Split(s,sep string)(result []string){
	result=make([]string,0,strings.Count(s,sep)+1)//优化Split函数
	i:=strings.Index(s,sep)
	//fmt.Println(len(sep),sep)
	
	for i>-1{
		result=append(result,s[:i])
		//s=s[i+1:]
		s=s[i+len(sep):]//修复sep多个字符的问题
		i=strings.Index(s,sep)
	}
	result=append(result,s)
	return
}
/*
1、性能比较函数
Fib是一个计算第n个斐波那契数的函数
*/
func Fib(n int)int{
	if n<2{
		return n
	}
	return Fib(n-1)+Fib(n-2)
}
