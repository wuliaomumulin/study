package main
import(
	"fmt"
	"runtime"
	"time"
	"sync"
	"sync/atomic"
	"strconv"
)
/*
并发
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
}
//返回调用者函数名称
func Getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。使用goeoutine非常简单，只需要在调用函数的时候加上go关键字,就可以为函数创建一个goroutine.
*/
func test1(){
	fmt.Println("---------",Getfunctionname(),"-------------")//打印函数名称
	go func(){
		fmt.Println("Hello,anonymous.")
	}()
	fmt.Println("Hello,test1.")
/*
这一个的执行结果只打印了hello test1,没有打印Hello,anonymous.为什么呢？
在程序启动时，Go程序就会为main()函数创建一个默认的goroutine.
当main()返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束。main函数所在的goroutine就类似于Boss,其他的就类似于follow,Boss一结束，follow全部结束。
所以我们要想办法让main等一等hello函数，最简单的方式就是time.Sleep()
*/
}
func test2(){
	fmt.Println("---------",Getfunctionname(),"-------------")//打印函数名称
	go func(){
		fmt.Println("Hello,anonymous.")
	}()
	fmt.Println("Hello,test1.")
	time.Sleep(time.Second)
}
/*
执行以上代码会发现，这一次先打印Hello,test1.然后接着打印Hello,anonymous.
首先为什么会先打印Hello,test1.是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。

一、启动多个goroutine的同步
多次执行下面的代码，会发现每次打印的数字的顺序都不一致，这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
*/
var wg3 sync.WaitGroup
func hello3(i int){
	defer wg3.Done()//goroutine结束就登记-1
	fmt.Println("hello",i)
}
func test3(){
	fmt.Println("---------",Getfunctionname(),"-------------")//打印函数名称
	for i:=0;i<10;i++{
		wg3.Add(1)//启动一个goroutine就登记+1
		go hello3(i)
	}
	wg3.Wait()//等待所有的goroutine都结束
}
/*
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码，默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器就会把go代码同时调度到8个OS线程上(GOMAXPROCS是m:n调度中n)
*/

func test4(){
	fmt.Println("---------",Getfunctionname(),"-------------")//打印函数名称
	runtime.GOMAXPROCS(2)//将逻辑核心设为2并行执行
	go func(){
		for i:=0;i<10;i++{
			fmt.Printf(" 一:%d ",i)
		}
	}()
	go func(){
		for i:=0;i<10;i++{
			fmt.Printf(" 二:%d ",i)
		}
	}()
	time.Sleep(time.Second)
/*
操作系统线程和goroutine的关系
	1、一个操作系统线程对应用户态多个goroutien;
	2、go程序可以同时使用多个操作系统线程
	3、goroutine和OS线程是多对多的关系，即m:n
*/
}
/*
channel是一种特殊的引用类型。通道像一个传送带或队列，总是遵循先入先出(First In First Out)规则，保证收发数据的准确性。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
	var 变量 chan 元素类型
	examlple
		var ch1 chan int   //声明一个传递整形的通道
		var ch2 chan bool  //声明一个传递布尔值的通道
		var ch3 chan []int //声明一个传递int切片的通道
创建channel
	通道是引用类型，类型空值为nil
初始化channel
	make(chan 元素类型.[缓存大小])
	channel的缓存大小是是可选的
		ch4 :=make(chan int)
		ch4 :=make(chan bool)
		ch4 :=make(chan []int)
操作channel
	通道有发送(send),接收(receive)和关闭(close)三种操作
	发送和接收都是用<-符号
	ch <- 10 //把10放到ch中
	x:= <- ch //从ch中接受值并赋值给变量x
	<-ch //从ch中接收值，忽略结果
	close(ch) //关闭通道
		关闭后的通道有以下特点:
			1、对一个关闭的通道发送值就会导致panic;
			2、对一个关闭的通道进行接收会一直获取值直到通道为空;
			3、对一个关闭的兵器没有值的通道执行接收操作会得到对应类型的零值;
			4、关闭一个已经关闭的通道会导致panic;
	*创建无缓存的通道，也叫同步通道.无缓存的通道只有在有goroutine接收值的时候才会发送值，否则会形成死锁。
*/
func recv5(c chan int){
	fmt.Println("接收成功",<-c)
}
func test5(){
	fmt.Println("---------",Getfunctionname(),"-------------")//打印函数名称
	ch:=make(chan int)
	go recv5(ch)//启用goroutine从通道接收值
	ch<-10
	fmt.Println("发送成功")
/*
无缓存通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时候才能发送成功，两个goroutine将相反执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到发送方goroutine在该通道上发送一个值.
*/
}
/*
有缓存的通道
解决上面的问题，我们还可以使用make函数初始化通道的时候为其指定通道的容量.

判断一个通道是否被关闭了？
*/
func test6(){
	fmt.Println("---------",Getfunctionname(),"-------------")//打印函数名称
	ch1:=make(chan int)
	ch2:=make(chan int)
	//开启goroutine将0~100的数发送到ch1中
	go func(){
		for i:=0;i<100;i++{
			ch1<-i
		}
		close(ch1)
	}()
	//开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func(){
		for {
			i,ok:=<-ch1
			if !ok{
				break
			}
			ch2<-i*i
		}
		close(ch2)
	}()
	//在主goroutine中打印从ch2中接收值
	//通道关闭后会退出for range循环
	for i:=range ch2{
		fmt.Printf("%d ",i)
	}
}
/*
单向通道
chan<- int是一个只写单向通道(只能对其写入int类型值)，可以对其执行发送操作但是不能执行接收操作
<-chan int是一个只读单向通道(只能对其读取int类型值),可以对其执行接收操作但是不能执行发送操作
*/
func counter7(out chan<- int){
	for i:=0;i<100;i++{
		out<-i
	}
	close(out)
}
func squarer7(out chan<- int,in <-chan int){
	for i:= range in{
		out<-i*i
	}
	close(out)
}
func printer7(in <-chan int){
	for i:=range in{
		fmt.Printf("%d ",i)
	}
}
func test7(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	ch1:=make(chan int)
	ch2:=make(chan int)
	go counter7(ch1)
	go squarer7(ch2,ch1)
	printer7(ch2)
}
/*
channel异常情况总结
channel 	nil 	非空			空的			满了			没满
接收 		阻塞	  	接收值		阻塞			接收者		接收者
发送	 		阻塞 	发送值		发送值 		阻塞			发送者
				 	关闭成功后，	关闭成功,	关闭成功后， 	关闭成功后，
关闭			panic	读取数据后，	返回零值。	读取数据后， 	读取数据后，
					返回零值。 				返回零值。 	返回零值.

关闭已经关闭的channel也会触发channel

二、worker pool(goroutine池)
一个简易的work pool示例代码
*/
func worker8(id int,jobs <-chan int,results chan<- int){
	for j:=range jobs{
		fmt.Printf("worker:%d start job:%d\n",id,j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n",id,j)
		results<-j*2
	}
}
func test8(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	jobs:=make(chan int,100)
	results:=make(chan int,100)
	//开启三个goroutine
	for w:=1;w<=3;w++{
		go worker8(w,jobs,results)
	}
	//五个任务
	for j:=1;j<=5;j++{
		jobs<-j
	}
	close(jobs)
	//输出结果
	for a:=1;a<=5;a++{
		fmt.Printf("results[%d] is %d \n",a,<-results)
	}
}
/*
三、select多路复用的使用
使用select语句能提高代码的可读性:
	可处理一个或多个channel的发送/接收操作;
	如果多个case同时满足，select会随机选择一个;
	对于没有case的select{}会一直等待，可用于阻塞main函数;
*/
func test9(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	ch:=make(chan int,1)
	for i:=0;i<10;i++{
		select {
		case x:=<-ch:
			fmt.Println("x",x)
		case ch<-i:
			fmt.Println("i",i)
		}
	}
}
/*
	并发安全和锁
*/
var x10 int64
var wg10 sync.WaitGroup
func add10(){
	for i:=0;i<5000;i++{
		x10+=int64(i)
	}
	wg10.Done()
}
func test10(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	wg10.Add(2)
	go add10()
	go add10()
	wg10.Wait()
	fmt.Println(x10)
	/*
	上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果和期待的不符。
	 */
}
/*
互斥锁
*/
var x11 int64
var wg11 sync.WaitGroup
var lock11 sync.Mutex
func add11(){
	for i:=0;i<5000;i++{
		lock11.Lock()
		x11+=int64(i)
		lock11.Unlock()
	}
	wg11.Done()
}
func test11(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	wg11.Add(2)
	go add11()
	go add11()
	wg11.Wait()
	fmt.Println(x11)
	/*
	上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果和期待的不符。
	 */
}
/*
读写互斥所
 */
var (
	x12 int64
	wg12 sync.WaitGroup
	lock12 sync.Mutex
	rwlock12 sync.RWMutex
)

func write12(){
	rwlock12.Lock()//加写锁
	x12+=1
	time.Sleep(10*time.Millisecond)//假设读操作时10ms
	rwlock12.Unlock()//减写锁
	//lock12.Unlock()//减互斥锁
	wg12.Done()
}
func read12(){
	rwlock12.RLock()//加读锁
	time.Sleep(time.Millisecond)//假设读操作耗时1ms
	rwlock12.RUnlock()//减读锁
	wg12.Done()
}
func test12(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	start:=time.Now()
	for i:=0;i<10;i++{
		wg12.Add(1)
		go write12()
	}
	for i:=0;i<1000;i++{
		wg12.Add(1)
		go read12()
	}
	wg12.Wait()
	end:=time.Now()
	fmt.Println(end.Sub(start))
	/*
	需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
	 */
}
/*
sync.WaitGroup
在代码中生硬的使用time.Sleep肯定是不合适的，go语言中可以使用sync.WaitGroup,来实现并发任务的同步。
需要注意sync.WaitGroup是一个结构体，传递的时候要传递指针
 */
var wg13 sync.WaitGroup
func hello13(){
	defer wg13.Done()
	fmt.Println("你好,goroutine")
}
func test13(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	wg13.Add(1)
	go hello13()
	fmt.Println("main goroutine done")
	wg13.Wait()
}
/*
sync.Once
在很多的场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等
*/
var icons14 map[string]string
var loadIconsOnce14 sync.Once
func loadIcons14(){
	icons14=map[string]string{
		"left":"left.png",
		"up":"up.png",
		"right":"right.png",
		"down":"down.png",
	}
}
//Icon是并发安全的
func Icon14(name string) string{
	loadIconsOnce14.Do(loadIcons14)
	return icons14[name]
}
func test14(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	fmt.Println(Icon14("up"))
}
/*
并发安全的单例模式singleton
sync.Once内部其实包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
 */
type singleton15 struct{}
var instance15 *singleton15
var once15 sync.Once
func GetInstance15() *singleton15{
	once15.Do(func(){instance15=&singleton15{}})
	return instance15
}
func test15(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	fmt.Println(GetInstance15())
}
/*
sync.Map
内置的map并不是并发安全的，在少量或许没问题，在大量的goroutine中修改map值可能会报错fatal error:concurrent map writes错误。
所以需要对map加锁，也可以直接使用内置的sync.Map
 */
func test16(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	var m=sync.Map{}
	wg:=sync.WaitGroup{}
	for i:=0;i<20;i++{
		wg.Add(1)
		go func(n int){
			key:=strconv.Itoa(n)
			m.Store(key,n)
			value,_:=m.Load(key)
			fmt.Printf("k=%v,v=%v\n",key,value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
/*
原子操作
代码中的加锁操作因为涉及到内核态的上下文切换会比较耗时，代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是go语言提供的方法他在用户态就可以完成，因此性能比加锁操作更好。go语言中原子操作有内置的标准库sync/atomic提供。

我们看一个实例来比较互斥锁和原子锁的性能
*/
type Counter17 interface{
	Inc()
	Load() int64
}
//普通版
type CommonCounter17 struct{
	counter17 int64
}
func (c CommonCounter17) Inc(){
	c.counter17++
}
func (c CommonCounter17) Load()int64{
	return c.counter17
}
//互斥锁版
type MutexCounter17 struct{
	counter17 int64
	lock sync.Mutex
}
func (m *MutexCounter17) Inc(){
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter17++
}
func (m *MutexCounter17) Load()int64{
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter17
}
//原子操作版
type AtomicCounter17 struct{
	counter17 int64
}
func (a *AtomicCounter17) Inc(){
	atomic.AddInt64(&a.counter17,1)//修改操作
}
func (a *AtomicCounter17) Load()int64{
	return atomic.LoadInt64(&a.counter17)//读取操作
}

func run17(c Counter17){
	var wg sync.WaitGroup
	start:=time.Now()
	for i:=0;i<1000;i++{
		wg.Add(1)
		go func(){
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end:=time.Now()
	fmt.Println(c.Load(),end.Sub(start))
}
func test17(){
	fmt.Println("\n---------",Getfunctionname(),"-------------")//打印函数名称
	c1:=CommonCounter17{}//非并发安全的
	run17(c1)
	c2:=MutexCounter17{} //使用互斥锁实现并发安全的
	run17(&c2)
	c3:=AtomicCounter17{} //并发安全且比互斥锁效率更高
	run17(&c3)
}
/*
atomic包提供了低层的原子级别操作，对于同步算法的实现很有用。这些函数必须谨慎的保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现更好。
 */