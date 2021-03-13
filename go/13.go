package main
import(
	"fmt"
	"time"
	"runtime"
)
/*
	时间
	func 函数名(参数)(返回值){函数体}
*	函数的参数和返回值都是可选的

*/
func main(){
	test1()
	test2()
	test3()
	test4()
	test5()
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
	时间格式和时间戳
*/
func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	now:=time.Now()
	fmt.Printf("current time:%v,format %d-%02d-%02d %02d:%02d:%02d\n",now,now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println("Unix time:",now.Unix())
	fmt.Println("UnixNano time:",now.UnixNano())
	timestampDemo2(now.Unix())

}
//可以使用unix将时间戳转换为时间格式
func timestampDemo2(timestamp int64){
	now:=time.Unix(timestamp,0)//将时间戳转为时间格式
	fmt.Printf("in timestampDemo2 current time:%v,format %d-%02d-%02d %02d:%02d:%02d\n",now,now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
}

/*
1、	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
2、time.Duration表示1纳秒，time.Second表示1秒


*/
func test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	now:=time.Now()
	//func (t Time) Add(d Duration) Time
	fmt.Printf("current time+1Hour:%v\n",now.Hour())

	//求两个时间差 func (t Time) Sub(u time) Duration


	
}
//3、使用time.Tick(时间间隔)来设置定时器，定时器的本质时一个通道(channel)
func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	ticker:=time.Tick(time.Second)//定义一个1s间隔的定时器
	for i:=range ticker{
		fmt.Println(i)//每秒都会执行的任务
		break
	}
}
//时间格式化
func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	now:=time.Now()
	//格式化的模板为go的出生日期为2006年1月2号15点04分 Mon Jan
	//24小时
	fmt.Println(now.Format("2006-01-02 15:04:05:000 Mon Jan"))
	//12小时
	fmt.Println(now.Format("2006-01-02 15:04:05:000 PM Mon Jan"))
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("15:04 2006-01-02"))
	fmt.Println(now.Format("2006-01-02"))

}
/*

*/
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	now:=time.Now()
	fmt.Println(now)
	//加载时区
	loc,err:=time.LoadLocation("Asia/Shanghai")
	if err!=nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj,err:=time.ParseInLocation("2006/01/02 15:04:05","2019/08/04 14:15:20",loc)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))//-13824h58m58.829775s

}