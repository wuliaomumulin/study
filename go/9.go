package main
/*
日志输出
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
 */
import (
	"fmt"
	"log"
	"os"
)

func main() {

	flag1()
	flag2()

	log.Println("这是一条很普通的日志")
	v :="很普通的"
	log.Printf("这是一条%s日志",v)
	log.Fatalln("这是一条出发fatal的日志")
	log.Panicln("这是一条panic的日志")

}

func flag1(){
	//格式化日志格式
	log.SetFlags(log.Llongfile|log.Lmicroseconds|log.Ldate)
	//设置日志前缀
	log.SetPrefix("[test]");
	//配置日志输出位置，默认输出到控制台
	logFile,err := os.OpenFile("./1.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err!=nil {
		fmt.Println("open log file error:",err)
		return
	}
	log.SetOutput(logFile)

	log.Println("这是一条很普通的日志")
}
func flag2(){
	logger := log.New(os.Stdout,"<安帝科技>",log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义级别的logger记录的日志")
}