package main
import(
	"fmt"
	"os"
)

func modifyArray(x [3]int){
	x[0] = 100
}
func modifyArray2(x [3][2]int){
	x[2][0] =100
}

/*
* Array操作
* 
*/
func main(){
	a := [...]int{1:1,2:2,3:5}
	fmt.Println(a)
	fmt.Printf("type of a:%T\n",a)

	var a1 = [...]string{"北京","上海","深圳"}
	//一维数组遍历
	for i:=0;i<len(a1);i++{
		fmt.Printf("%s",a1[i])
	}
	//第二种
	for index,value := range a1{
		fmt.Printf("%d=%s\t",index,value)//\t制表符
	}
	fmt.Println("----------多维数组的操作------------")
	/*
	* 多维数组的操作
	*/
	var a2 = [3][2]string{
		{"北京","上海"},
		{"广州","深圳"},
		{"成都","重庆"},
	}
	fmt.Println(a2)
	fmt.Println(a2[2][1])//重庆
	//遍历
	for _,v1 := range a2{
		for _,v2 := range v1{
			fmt.Printf("%s\t",v2)
		}
		fmt.Println()
	}
	//多维数组只支持第一层使用...来推导长度，内层不支持
	/*
	支持的写法
	a = [...][2]string{
		{"北京","上海"},
		{"广州","深圳"},
		{"成都","重庆"},
	}
	不支持的写法
	a1 = [3][...]string{
		{"北京","上海"},
		{"广州","深圳"},
		{"成都","重庆"},
	}
	*/

	/*数组是值类型,
	类似php作用域的概念

	*/
	var a3=[3]int{10,20,30}
	modifyArray(a3)//在modify中修改的是a的副本的x
	fmt.Println(a3)//10,20,30
	a4 :=[3][2]int{
		{1,1},
		{1,1},
		{1,1},
	}
	modifyArray2(a4)//在modify中修改的是a4的副本的x
	fmt.Println(a4)//[[1,1],[1,1],[1,1]]

	var a5 = [...]int{1,3,5,7,9}
	//累加
	var sum = 0
	for _,v := range a5{sum += v}
	fmt.Println(sum)
	//找出指定值的下标和为8的两个值的下标
	sum = 8;
	var a6 = [...]int{1,3,5,7,8}
	for i,j := range a6{
		for l,k := range a6[i+1:]{
			if(j+k==sum){fmt.Printf("(%d,%d)\t",i,l+i+1)} 
		}
	}
	/*
	Fprint系列函数将内容输出到一个io.writer接口类型的变量w中
	*/
	fmt.Fprintln(os.Stdout,"向标准输出写入内容")
	fileObj,err := os.OpenFile("./xx.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err!= nil {
		fmt.Println("打开文件出错,err:",err)
		return 
	}
	name := "找事情"
	//向打开的文件句柄写入内容
	fmt.Fprintf(fileObj,"往文件中写入信息:%s",name)
	/**
	* Sprint函数会将传入的数据拼接并返回一个字符串
	*/
	fmt.Println(fmt.Sprint("深入","理解"))
	//fmt.Errorf("这是一个错误:%v",111)
}