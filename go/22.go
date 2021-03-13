package main
import(
	"fmt"
	"reflect"
	"runtime"
)
/*
反射
反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身信息。
支持反射的语言可以在程序编译期将变量的反射信息，如字段名称，类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改他们。
go程序在运行期使用reflect包访问程序的反射信息
在上面我们介绍了空接口。空接口可以存储任意类型的变量，那么我们如何知道这个空接口保存的数据是什么呢？
反射就是在运行时动态的获取一个变量类型信息和值信息


*/
func main(){
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
reflect包
任意接口值都由是一个具体类型和具体类型的值两部分组成的。任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.Typeof和reflect.Valueof两个函数来获取任意对象的Value和Type.
使用reflect.Typeof()函数可以获得任意值的类型对象(reflect.Type),程序通过类型对象可以访问任意值的类型信息.
*/
func reflectType(x interface{}){
	v:=reflect.TypeOf(x)
	fmt.Printf("type:%v\n",v)
}
func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
}
/*
type name和type kind
在反射中关于类型还分为两种:类型(Type)和种类(Kind).因为在go语言中我们可以使用type关键字构造很多自定义类型，而种类(Kind)就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类(Kind).举个例子，我们定义了两种指针类型的两种结构体类型，通过反射查看他们的类型和种类:


type Kind uint

const (
	Invalid Kind = iota //非法类型
 	Bool                //布尔值
	Int                 //有符号整形
	Int8                //有符号8位整形
	Int16               //有符号16位整形
	Int32               //有符号23位整形
	Int64               //有符号64位整形
	Uint                //无符号整形
	Uint8               //无符号8位整形
	Uint16              //无符号16位整形
	Uint32              //无符号32位整形
	Uint64              //无符号64位整形
	Uintptr             //指针
	Float32             //单精度浮点数
	Float64             //双精度浮点数
	Complex64           //64位复数类型
	Complex128          //128位复数类型
	Array               //数组
	Chan                //通道
	Func                //函数
	Interface           //接口
	Map                 //映射
	Ptr                 //指针
	Slice               //切片
	String              //字符串
	Struct              //结构体
	UnsafePointer       //底层指针
)
*/
type MyInt2 int64
func reflectType2(x interface{}){
	v:=reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v\n",v.Name(),v.Kind())
}
func test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var a *float32 //指针
	var b MyInt2 //自定义类型
	var c rune //类型别名,代表一个utf-8字符，也可以表示为int32
	//向数组、切片、Map、指针等类型的变量，他们的.Name都是返回空
	reflectType2(a)//type:,kind:ptr
	reflectType2(b)//type:MyInt2,kind:int64
	reflectType2(c)//type:int32,kind:int32

	type person struct{
		name string
		age int
	}
	type book struct{title string}
	var d=person{
		name:"张三",age:29,
	}
	var e=book{title:"平凡的世界"}
	reflectType2(d)//type:person,kind:struct
	reflectType2(e)//type:book,kind:struct
}
/*
reflect.ValueOf返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值可以相互暗转。
 */
func reflectType3(x interface{}){
	v:=reflect.ValueOf(x)
	k:=v.Kind()
	switch k{
	case reflect.Int64:
		//v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64,value is %d\n",int64(v.Int()))
	case reflect.Float32:
		//v.Float()从反射中获取整型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32,value is %f\n",float32(v.Float()))
	case reflect.Float64:
		//v.Float()从反射中获取整型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64,value is %f\n",float64(v.Float()))	
	}
}
func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var a float32 =4.14
	var b int64 =100
	reflectType3(a)//type is float32,value is 4.140000
	reflectType3(b)//type is int64,value is 100
	//将int类型的原始值转换为reflect.Value类型
	fmt.Printf("type:%T\n",reflect.ValueOf(10))//type:reflect.Value
}
/*
通过反射设置变量的值
想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值
*/
func reflectSetValue41(x interface{}){
	v:=reflect.ValueOf(x)
	if v.Kind()==reflect.Int64 {
		v.SetInt(200)//修改的是副本,reflect包会引发panic
	}
}
func reflectSetValue42(x interface{}){
	v:=reflect.ValueOf(x)
	//反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind()==reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var a int64 = 100
	//reflectSetValue41(a)//panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	fmt.Println(a)
	reflectSetValue42(&a)
	fmt.Println(a)
}
/*
IsNil()常被用于判断指针是否为空，IsValid()常被用于判断返回值是否有效。
*/
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var a *int
	fmt.Println("var a *int IsNil:",reflect.ValueOf(a).IsNil())
	fmt.Println("nil IsValid:",reflect.ValueOf(a).IsValid())
	//实例化一个匿名结构体
	b :=struct{}{}
	//尝试从结构体中查找某个字段
	fmt.Println("不存在的结构体成员:",reflect.ValueOf(b).FieldByName("abc").IsValid())
	fmt.Println("不存在的结构体方法:",reflect.ValueOf(b).MethodByName("abc").IsValid())
	//map
	c:=map[string]int{}
	//尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键",reflect.ValueOf(c).MapIndex(reflect.ValueOf("abc")).IsValid())
}
/*
结构体反射
任意值通过reflect.TypeOf()获得反射对象信息后，如果他的类型是结构体，可以通过反射值对象(reflect.Type)的NumField()和Field()方法获得结构体成员的详细信息。

struct类型用来描述结构体中的一个字段的信息
type StructField struct {
	Name      string    //字段名称        
	PkgPath string      //pkgpath是非导出字段的包路径，对导出字段该字段为""
	Type      Type      // 类型
	Tag       StructTag // 标签
	Offset    uintptr   // 字段在结构体中的偏移量
	Index     []int     // 基于Type.FieldByIndex()时的索引切片
	Anonymous bool      // 是否匿名字段
}

当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段消息，也可以通过字段名去获取字段消息
 */
type student struct{
	//不需要导入encoding/json包
	name string `json:"name"`
	score int `json:"score"`
}
func test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	stu1:=student{name:"张三",score:90}
	t:=reflect.TypeOf(stu1)
	fmt.Println(t.Name(),t.Kind())
	//通过for循环遍历结构体的所有字段信息
	for i:=0;i<t.NumField();i++ {
		field:=t.Field(i)
		fmt.Printf("name:%s,index:%d,type:%v,json tag:%v\n",field.Name,field.Index,field.Type,field.Tag.Get("json"))
	}
	//通过字段名获取指定结构体字段信息
	if scoreField,ok:=t.FieldByName("Score");ok {
		fmt.Printf("name:%s,index:%d,type:%v,json tag:%v\n",scoreField.Name,scoreField.Index,scoreField.Type,scoreField.Tag.Get("json"))
	}
}
/*
遍历打印s包含的方法
 */
//注意函数首字母大写
func (s student) Study() string{
	msg:="好好学习，天天向上"
	fmt.Println(msg)
	return msg
}
func (s student) Sleep() string{
	msg:="保持睡眠，快快向大"
	fmt.Println(msg)
	return msg
}
func printMethod(x interface{}){
	t:=reflect.TypeOf(x)
	v:=reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	for i:=0;i<v.NumMethod();i++ {
		methodType:=v.Method(i).Type()
		fmt.Printf("method name:%s,method type:%s\n",t.Method(i).Name,methodType)
		//通过反射调用方法传递的参数必须是[]reflect.Value类型
		var args=[]reflect.Value{}
		v.Method(i).Call(args)
	}
}
func test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	stu1:=student{name:"张三",score:90}
 	printMethod(stu1)
}
/*
反射不应该被滥用的原因:
1、基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会触发panic,那很有可能是在代码写完的很长时间之后;
2、大量使用反射的代码通常难以理解;
3、反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级;
 */