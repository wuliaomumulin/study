package main
import(
	"fmt"
	"runtime"
	"unsafe"
	"encoding/json"
)
/*
结构体
go里面没有类的概念，也不支持类的继承等面向对象的概念。go语言通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
go的基本类型有stirng,int,float,bool,另外可以通过type自定义类型
	如：type Myint int
类型别名:	
	type byte = uint8
	type rune = int32
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
	test9()
	test10()
	test11()
	test12()
	test13()
	test14()
	test15()
	test16()
	test17()
	test18()
}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
//类型别名和定性定义的区别
func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type NewInt int//类型定义
	type MyInt = int//类型别名
	var a NewInt
	var b MyInt
	fmt.Printf("type a is %T,type b is %T\n",a,b)//type a is main.NewInt,type b is int
}
//使用type和struct来定义结构体
type person struct {
	name,nickname string
	city string
	age int8
}
func test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var p1 person
	p1.name="zhangsan"
	p1.nickname="张三"
	p1.city="内蒙古"
	p1.age=29
	fmt.Printf("p1=%v\n",p1)//只输出值
	fmt.Printf("p1=%#v\n",p1)//键值都会输出
}
//匿名结构体
func test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var user struct{Name string;Age int}
	user.Name="李四"
	user.Age=27
	fmt.Printf("user=%#v\n",user)//user=struct { Name string; Age int }{Name:"李四", Age:27}	
}
//创建指针类型结构体会得到结构体地址，支持对结构体指针只用使用.来访问结构体成员
func test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var p1 = new(person)
	fmt.Printf("%T\n",p1)//*main.person
	fmt.Printf("p1=%#v\n",p1)//p1=&main.person{name:"", nickname:"", city:"", age:0}
	p1.name="baolisi"
	p1.nickname="鲍里斯"
	p1.city="china"
	p1.age=45
	fmt.Printf("p1=%#v\n",p1)//p1=&main.person{name:"baolisi", nickname:"鲍里斯", city:"china", age:45}
}
//取结构体地址实例化&
func test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var p1 = &person{}
	fmt.Printf("%T\n",p1)//*main.person
	fmt.Printf("p1=%#v\n",p1)//p1=&main.person{name:"", nickname:"", city:"", age:0}
	p1.name="tanya"
	//p1.name="tanya"其实底层是(*p1).name="tanya",这是go燕燕为我们实现的语法糖
	p1.nickname="谭雅"
	p1.city="china"
	p1.age=29
	fmt.Printf("p1=%#v\n",p1)//p1=&main.person{name:"tanya", nickname:"谭雅", city:"china", age:29}
}
/*
结构体初始
化	默认为对应其类型的零值

 */
func test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var p1 person
	fmt.Printf("p1=%#v\n",p1)//p1=&main.person{name:"", nickname:"", city:"", age:0}
	var p2 = person{name:"zhangsan",nickname:"张三",city:"beijing",age:25}
	var p3 = &person{name:"zhangsan",nickname:"张三",city:"beijing",age:25}
	var p4 = &person{name:"zhangsan",nickname:"张三",age:25}
	var p5 = &person{"zhangsan","张三","beijing",25}//键初始化时必须要所有字段、顺序一致，而且不能和键值初始化混用
	fmt.Printf("p2=%#v\n",p2)
	fmt.Printf("p3=%#v\n",p3)
	fmt.Printf("p4=%#v\n",p4)
	fmt.Printf("p5=%#v\n",p5)
}
//结构体占用一块连续的内存
func test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type test struct{
		a int8
		b int8
		c int8
		d int8
	}
	n:=test{1,2,3,4}
	fmt.Printf("n.a=%p,n.b=%p,n.c=%p,n.d=%p\n",&n.a,&n.b,&n.c,&n.d)
}
//空结构体是不占用内存的
func test8(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	var v struct{}
	fmt.Println(unsafe.Sizeof(v))//0
}
//正常使用
func test9(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type student struct{
		name string
		age int
	}
	m:=make(map[string]*student)//一个map切片
	stus:=[]student{
		{name:"张三",age:21},
		{name:"李四",age:24},
		{name:"王五",age:20},
	}
	for _,stu:=range stus{
		m[stu.name] = &stu
	}
	for k,v:=range m{
		fmt.Println(k,"=",v.name)//张三=张三
	}
}
/*
构造参数
	go没有构造参数，需要自己实现。
	因为下面的实现构造参数方法，返回值是struct类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造参数返回值是结构体指针类型
 */
func newPerson(name,nickname,city string,age int8) *person{
	return &person{
		name:name,
		nickname:nickname,
		city:city,
		age:age,
	}
}
func test10(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	p1:=newPerson("zhangsan","张三","beijing",25)
	fmt.Printf("%#v\n",p1)//&main.person{name:"zhangsan", nickname:"张三", city:"beijing", age:25}
	//方法于函数的区别是，方法属于特定的类型，函数不属于特定的类型，范围大小:方法<函数
	p1.Dream()
	//指针类型的接收者
	p1.SetAge(21)
	fmt.Println(p1.age)
	//值类型的接受者,值得改变只作用在SetAge2函数之内，之外并没有改变
	p1.SetAge2(22)//在函数内为22,但是在父函数里依然没有改变
	fmt.Println(p1.age)

	//任意类型添加方法
	var m1 MyInt
	m1.sayhello()
	m1=100
	fmt.Printf("m1 value is %#v,type is %T\n",m1,m1)		
}
/*
方法和接收者
	go语言中的方法Method是一种作用于特定类型变量的函数，叫做接收者Receiver.接收者的概念就类似于其他语言中的this或者self
	格式如下:
		func (接收者变量 接收者变量类型) 方法名(参数类型) (返回参数) {函数体}
			接收者变量：接受者的参数变量名再命名时，官方建议使用接收者类型名称首字母的小写，而不是self,this之类的命名,例如Person类型的接收者变量应该为p,Connector应该为c;
			接受者类型:可以是指针和非指针类型;
			方法名，参数列表、返回参数:具体格式和函数定义一致;
 */
func (p person) Dream(){
	fmt.Println(p.nickname,"的梦想是学好go语言")
}
/*
参数类型的接收者
	由一个结构体指针组成，由于指针的特性，调用该方法时修改接收者指针的任意成员变量，在方法结束后都是有效的。这种方法就很接近面向对象中的this和self.
 */
func (p *person) SetAge(age int8){
	p.age=age
}
/*
值类型的接收者
	
 */
func (p person) SetAge2(age int8){
	p.age=age
}
/*
什么时候需要指针类型呢？
1、需要修改接受者的值;
2、接受者是拷贝代价比较大的大对象;
3、保证一致性，如果有某个方法使用了指针接受者，那么其他的方法应该也使用指针接受者;


任意类型添加方法
1、我们可以基于内置任意类型使用type关键字定义任意的新的自定义类型，然后为我们的自定义类型添加任意方法

 */
type MyInt int
func  (m MyInt) sayhello(){
	fmt.Println("我是一个基于自定义类型的方法")
}
/*
结构体的匿名字短
匿名结构体默认采用类型名作为字段名，而字段名称必须唯一，所以一个结构体中同种类型的匿名字段只能有一个
*/

func test11(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type person struct{
		string
		int
	}
	p1:=person{"张三",30}
	fmt.Printf("p1 is %#v\n",p1)
	fmt.Println(p1.string,p1.int)
}
/*
结构体的嵌套
*/
func test12(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type Address struct{
		province string
		city string
	}
	type User struct{
		name,gender string
		address Address
	}
	user1:=User{name:"张三",gender:"男",address:Address{province:"北京市",city:"昌平区"}}
	fmt.Printf("user1 is %#v\n",user1)
}
/*
结构体嵌套的匿名字段
1、当访问结构体成员时会先在结构体中查找字段，找不到再去嵌套结构体的匿名字段中查找
*/
func test13(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type Address struct{
		province string
		city string
	}
	type User struct{
		name,gender string
		Address
	}
	var user1 User
	user1.name="张三"
	user1.gender="男"
	user1.Address.province="北京市"//匿名字段默认使用类型名作为字段
	user1.city="昌平区"//匿名字段可以省略
	fmt.Printf("user1 is %#v\n",user1)
}
/*
结构体嵌套匿名字段重冲突
1、嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名
*/
func test14(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type Address struct{
		province,city,createtime string
	}
	type Email struct{
		account,createtime string
	}
	type User struct{
		name,gender string
		Address
		Email
	}

	var user1 User
	user1.name="张三"
	user1.gender="男"
	user1.Address.province="北京市"//匿名字段默认使用类型名作为字段
	user1.city="昌平区"//匿名字段可以省略
	user1.Address.createtime="2021-03-06"
	user1.Email.createtime="2021-03-05"
	fmt.Printf("user1 is %#v\n",user1)
}
/*继承实现*/
type animal struct{
	name string
}
func (a animal) move(){
	fmt.Printf("%s会动",a.name)
}
type dog struct{
	feet uint8
	animal
}
func (d dog) wang(){
	fmt.Printf("%s会发出声音",d.name)
}
func test15(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	dog1 := &dog{feet:4,animal:animal{name:"贝贝"}}
	fmt.Printf("dog1 is %#v\n",dog1)
}
/*
结构体字段可见性
 */
func test16(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称

	type Student struct{
		Id int
		Gender string
		Name string
	}
	type Class struct{
		Title string
		Students []*Student
	}
	c:=&Class{
		Title:"59",Students:make([]*Student,0,200),
	}
	for i:=1;i<=10;i++{
		stu:=&Student{
			Name:fmt.Sprintf("stu%02d",i),
			Gender:"男",
			Id:i,
		}
		c.Students=append(c.Students,stu)
	}
	//JSON序列号：结构体-->Json格式的字符串
	data,err:=json.Marshal(c)
	if err!=nil{
		fmt.Println("json Marshal failed")
		return
	}
	fmt.Printf("data is %s\n",data)
	//JSONA反序列化，JSON格式的字符创-->结构体
	c1:=&Class{}
	err=json.Unmarshal([]byte(data),c1)
	if err!=nil{
		fmt.Println("json Unmarshal failed")
		return
	}
	fmt.Printf("c1 is %#v\n",c1)
}
/*
结构体标签

 */
func test17(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	type student struct{
		Id int `json:"id"` //通过指定tag实现json序列化该字段时的key
		Gender string //json序列化是默认使用字段名作为key
		name  string //私有属性不能带json包访问到
	}
	s1:=student{
		Id:1,Gender:"男",name:"私有属性",
	}
	data,err:=json.Marshal(s1)
	if err!=nil{
		fmt.Println("json Marshal failed")
		return
	}
	fmt.Printf("data is %s\n",data)	
}
/*
因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意，在方法中使用传入的slice的拷贝进行结构体赋值

不过版本的go修复了这个问题
 */
type person1 struct{
	name string
	age int8
	dreams []string
}
func (p *person1) SetDreams(dreams []string){
	p.dreams = dreams
	copy(p.dreams,dreams)
}
func test18(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	p1:=person1{name:"张三",age:20}
	data:=[]string{"吃饭","睡觉","打豆豆"}
	p1.SetDreams(data)
	//您真的想要修改p1.dreams吗？
	data[1]="不睡觉"
	fmt.Println(p1.dreams)
}