package main
import(
	"fmt"
	"runtime"
	"time"
	"sort"
	"strings"
	"math/rand"
)
/*
	map是一种无序的基于key-value的数据结构，go语言的map为引用类型，必须初始化才能使用
	map[KeyType]ValueType
	初始化或者分配内容
	make(map[KeyType]ValueType,[cap])

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
	map的声明和使用
*/
func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称


	scoreMap:=make(map[string]int,8)
	scoreMap["zhangsan"]=100
	scoreMap["lisi"]=90
	fmt.Println(scoreMap)//map[lisi:90 zhangsan:100]
	fmt.Println(scoreMap["lisi"])//90
	fmt.Printf("type is a:%T",scoreMap)//type is a:map[string]int

}
//声明填充
func test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	userInfo:=map[string]string{
		"username":"zhangsan",
		"password":"123456",

	}
	fmt.Println(userInfo);//map[password:123456 username:zhangsan]
}
//判断某个值是否存在
func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	scoreMap:=map[string]int{
		"zhangsan":100,
		"lisi":80,
	}
	//如果key为ok为true,v为对应的值，不存在为false,v为值类型的零值
	v,ok:=scoreMap["zhangsan"]
	if ok {
		fmt.Println(v);//100
	}else{
		fmt.Println(v);//0
	}

	//for循环遍历map
	fmt.Println("for each map");
	for k,v:=range scoreMap{
		fmt.Printf("k=%v,v=%d&",k,v)
	}
	fmt.Println();
	fmt.Println("only each k");
	for k:=range scoreMap{
		fmt.Printf("k=%v&",k)
	}
	fmt.Println();
	fmt.Println("only each v");
	for _,v:=range scoreMap{
		fmt.Printf("k=%d&",v)
	}
}
//删除键值对
func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	scoreMap:=make(map[string]int,8)
	scoreMap["zhangsan"]=100
	scoreMap["lisi"]=90
	scoreMap["wangwu"]=80
	delete(scoreMap,"lisi")
	for k,v:=range scoreMap{
		fmt.Printf("k=%v,v=%d&",k,v)
	}

}
//按照指定顺序遍历map
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	rand.Seed(time.Now().UnixNano())//初始化随机种子
	var scoreMap=make(map[string]int,200)

	for i:=0;i<100;i++ {
		key:=fmt.Sprintf("stu%02d",i)//生成stu开头的字符串
		value:=rand.Intn(100)//生成0-99的随机数
		scoreMap[key]=value
	}	
	//取出map中的所有key存入切片keys
	var keys=make([]string,0,200)
	for key:=range scoreMap{
		keys=append(keys,key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _,key:=range keys{
		fmt.Printf("k=%02d,v=%v&",key,scoreMap[key])
	}

}
/*
元素为map类型的切片
*/
func test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var myslice = make([]map[string]string,3)
	for i,v:=range myslice{
		fmt.Printf("i:%d,value:%v ",i,v)
	}
	//after init
	myslice[0]=make(map[string]string,10)
	myslice[0]["name"] = "zhangsan"
	myslice[0]["pass"] = "123456"
	myslice[0]["address"] = "hebei"
	for i,v:=range myslice{
		fmt.Printf("i=%d,v=%v\n",i,v)
	}
}
/*
map中值为切片类型的操作
*/
func test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var sliceMap=make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key:="china"
	value,ok:=sliceMap[key]
	if !ok {
		value = make([]string,0,2)
	}
	value=append(value,"beijing","shanghai")
	sliceMap[key]=value
	fmt.Println(sliceMap)//map[china:[beijing,shanghai]]
}
/*
	统计单词个数
*/
func test8(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	str:="This program is distributed in the hope that it will be useful is"
	strslice :=strings.Split(str," ")
	fmt.Println(strslice)//map[china:[beijing,shanghai]]
	countmap:=make(map[string]int,20)
	for _,v:=range strslice{
		countmap[v]++
	}
	fmt.Println(countmap)
}
