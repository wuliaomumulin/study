package main
import(
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"io"
	"runtime"
)
/*
	文件操作

*/
func main(){
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
	文件操作
*/
func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	//只读模式下打开文件
	file,err:=os.Open("./13.go")
	if err!=nil {
		fmt.Println(err)
		return
	}
	file.Close()

}
/*
	func (f *File) Read(b []byte) (n int,err error)
	他接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾会返回0和io.EOF.
*/
func test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	file,err:=os.Open("./context.md")
	if err!=nil {
		fmt.Println("open file failed!,err:",err)
		return
	}
	defer file.Close()
	//使用Read方法读取数据
	var tmp = make([]byte,128)
	n,err:=file.Read(tmp)
	if err==io.EOF {
		fmt.Println("file read end")
		return
	}
	if err!=nil {
		fmt.Println("file read failed,err:",err)
		return
	}
	fmt.Printf("read %d byte data\n",n)
	fmt.Println(string(tmp[:n]))
}
//使用for循环读取文件中的数据
func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	
	file,err:=os.Open("./context.md")


	if err!=nil {
		fmt.Println("open file failed!,err:",err)
		return
	}
	defer file.Close()
	//循环读取文件
	var tmp [128]byte
	var content []byte


	for {

		n,err:=file.Read(tmp[:])
		if err==io.EOF {
			fmt.Println("file read end")
			goto label
		}
		if err!=nil {
			fmt.Println("file read failed,err:",err)
			return
		}

		fmt.Printf("read %d byte data\n",n)
		//fmt.Println(string(tmp[:n]))//这里输出的

		content=append(content,tmp[:n]...)


	}
	//
	label:
	fmt.Println("aaa")	
	fmt.Println(string(content))

}
//使用bufio读取文件
func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	file,err:=os.Open("./context.md")


	if err!=nil {
		fmt.Println("open file failed!,err:",err)
		return
	}
	defer file.Close()
	
	reader:=bufio.NewReader(file)
	for{
		line,err:=reader.ReadString('\n')//注意是字符
		if err==io.EOF {
			if len(line)!=0 {
				fmt.Println("bb",line)
			}
			fmt.Println("file read end")
			break
		}
		if err!=nil {
			fmt.Println("file read failed,err:",err)
			return
		}
		fmt.Print("aa",line)
	}
}
/*
ioutil读取整个文件
*/
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	content,err:=ioutil.ReadFile("./context.md")
	if err!=nil {
		fmt.Println("open file failed!,err:",err)
		return
	}
	fmt.Println(string(content))

}
/*
文件写入操作
func OpenFile(name string, flag int, perm FileMode) (*File, error) {}
1、文件名
2、flag
os.O_WRONLY 只写
os.O_CREATE 创建文件
os.O_RDONLY 只读
os.O_RDWR 读写
os.O_TRUNC 清空
os.O_APPEND 追加
3、perm:文件权限，一个8进制数,r(读)04,w(写)02,x(执行)01


Write和WriteString
*/
func test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	file,err:=os.OpenFile("./context.md",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err!=nil {
		fmt.Println("open file failed!,err:",err)
		return
	}
	defer file.Close()
	str:="hello,你好"
	file.Write([]byte(str))//写入字节切片数据
	file.WriteString("hello,你好")//直接写入数据

}
//bufio.NewWrite
func test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	file,err:=os.OpenFile("./context.md",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err!=nil {
		fmt.Println("open file failed!,err:",err)
		return
	}
	defer file.Close()

	writer:=bufio.NewWriter(file)
	for i:=0;i<10;i++{
		writer.WriteString("hello,你好\n")//直接写入数据
	}
	writer.Flush()//将缓存中的数据写入文件

}
/**
ioutil.WriteFile
*/
func test8(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	str:="hello,你好"
	err:=ioutil.WriteFile("./context.md",[]byte(str),0666)
	if err!=nil {
		fmt.Println("write file failed!,err:",err)
		return
	}
}