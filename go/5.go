package main
import(
	"fmt"
	"sort"
)


func Sort(){
	var a = [...]int{3,7,8,9,1}
	//sort包内实现了内部数据类型的排序
	sort.Ints(a[:])//将数组切片
	fmt.Println(a)
	//降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.println(a)
}
/*
* Slice切片操作
* var name []T
*/
func main(){
	var a []string
	var a1 = []int{}
	var a2 = []bool{false,true}
	//var a3 = []bool{false,true}
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a==nil)
	fmt.Println(a1==nil)
	fmt.Println(a2==nil)
	/*
	fmt.Println(a2==a3)
	切片是引用类型，不支持直接比较，只能和nil比较
	*/
	//基于数组定义切片
	b := [5]int{55,56,57,58,59}
	b1 := b[1:4]//基于数组a创建切片，包含元素a[1],a[2],a[3]
	fmt.Println(b1) //[56 57 58]
	fmt.Printf("type of b:%T\n",b1)//type of b:[]int
	//还支持如下格式
	fmt.Println(b[1:])
	fmt.Println(b[:4])
	fmt.Println(b[:])
	//切片再切片,cap求容量
	var b2 = [...]string{"北京","上海","广州","深圳","成都","重庆"}
	fmt.Printf("b2:%v type:%T len:%d cap:%d\n",b2,b2,len(b2),cap(b2))
	b3:=b2[1:3]//>=1 and <3 所以有两个
	fmt.Printf("b3:%v type:%T len:%d cap:%d\n",b3,b3,len(b3),cap(b3))
	b4:=b3[1:]
	fmt.Printf("b4:%v type:%T len:%d cap:%d\n",b4,b4,len(b4),cap(b4))
	fmt.Println("使用make函数构造切片")
	var b5 = make([]int,2,10)
	fmt.Println(b5)
	fmt.Println(len(b5))
	fmt.Println(cap(b5))
	//
	var b6 []int
	var b7 =[]int{}
	var b8 =make([]int,0)
	fmt.Println("要判断一个切片是否为空，要是用len(s)==0来判断，不应该使用s==nil来判断,因为他会发生一下错误，导致不准确")
	fmt.Printf("b6:%v type:%T len:%d cap:%d,isnil:%v\n",b6,b6,len(b6),cap(b6),b6==nil)
	fmt.Printf("b7:%v type:%T len:%d cap:%d,isnil:%v\n",b7,b7,len(b7),cap(b7),b7==nil)
	fmt.Printf("b8:%v type:%T len:%d cap:%d,isnil:%v\n",b8,b8,len(b8),cap(b8),b8==nil)
	fmt.Println("切片的拷贝赋值")
	b9 := make([]int, 3)//[0 0 0]
	b10 := b9[1:] //将b9直接赋值给b10,b9和b10公用一个底层数组
	b10[0] = 100
	fmt.Println(b9)
	fmt.Println(b10)
	fmt.Println("切片的遍历--第一种")
	for i:=0;i<len(b9);i++{
		fmt.Printf("%d\t",b9[i])
	}
	fmt.Println("切片的遍历--第二种")
	for index,value := range b10{
		fmt.Printf("%d=%d\t",index,value)
	}
	fmt.Println("append为切片添加元素")
	var b11 []int
	for i:=0;i<10;i++{
		b11 = append(b11,i,i+1)
		fmt.Printf("%v len:%d,cap:%d,ptr:%p\n",b11,len(b11),cap(b11),b11)
	}
	fmt.Println("使用copy赋值切片")
	var b12 = make([]int,9,9)
	copy(b12,b11)
	b12[0] = 100
	fmt.Println(b11,b12)//截取9个
	fmt.Println("切片原色删除,我们可以使用切片本身的特性来删除元素")
	b12 = append(b12[:2],b12[3:]...)
	fmt.Println("b12",b12)//截取9个
	var index = 3;
	fmt.Println("去掉index=2的元素:",append(b12[:index],b12[index+1:]...))//<少于index and 大于等于index+1

	var b13 =make([]string,5,10)
	for i:=0;i<10;i++{
		b13 = append(b13,fmt.Sprintf("%v",i))
	}
	fmt.Println(b13)

	//sort()
}