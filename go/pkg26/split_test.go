package pkg26

import(
	"reflect"
	"testing"
	"time"
	"os"
	"fmt"
)
/*
一、测试函数
	func TestName(t *testing.T){}
	*必须以Test开头，后缀名首字符必须大写,必须接收一个*testing.T的参数
在pkg26包路径下，执行go test命令，可以看到输出结果如下:

PASS
ok      _/F_/www/go/20210224/pkg26      0.054s

*/
func TestSplit(t *testing.T){
	got:=Split("a:b:c",":")//程序输出的结果
	want:=[]string{"a","b","c"}//期望的结果
	if !reflect.DeepEqual(want,got){//因为slice不能直接比较，借助反射包中的方法比较
		t.Errorf("execpted:%v,got:%v",want,got)//测试失败输出结果提示
	}
}
/*
一个测试用例有点单薄，我们再编写一个测试使用多个字符串切割字符串的例子。
结果如下

--- FAIL: TestMoreSplit (0.00s)
    split_test.go:31: execpted:[a b],got:[a cd]
FAIL
exit status 1
FAIL    _/F_/www/go/20210224/pkg26      0.063s

*/
func TestMoreSplit(t *testing.T){
	got:=Split("abcd","bc")
	want:=[]string{"a","d"}//期望的结果
	if !reflect.DeepEqual(want,got){//因为slice不能直接比较，借助反射包中的方法比较
		t.Errorf("execpted:%v,got:%v",want,got)//测试失败输出结果提示
	}
}

/*
这一次我们测失败了，我们可以为go test命令添加参数-v,查看测试函数名和运行时间
go test -v

=== RUN   TestSplit
--- PASS: TestSplit (0.00s)
=== RUN   TestMoreSplit
    split_test.go:39: execpted:[a b],got:[a cd]
--- FAIL: TestMoreSplit (0.00s)
FAIL
exit status 1
FAIL    _/F_/www/go/20210224/pkg26      0.059s

1、正则函数名匹配测试函数:
go test -v -run="More"

二、测试组
我们现在还想要测试以下split函数对中文字符串的支持，这个时候我们可以再编写一个TestClineseSplit测试函数，但是我们也可以使用如下更友好的一种方式来添加更多的测试用例。
*/
func TestChineseSplit(t *testing.T){
	//定义一个测试用例类型
	type test struct{
		input,sep string
		want []string
	}
	//定义一个储存测试用例的切片
	tests:=[]test{
		{input:"a:b:c",sep:":",want:[]string{"a","b","c"}},
		{input:"a,b,c",sep:",",want:[]string{"a","b","c"}},
		{input:"abcd",sep:"bc",want:[]string{"a","d"}},		
		{input:"都完成中...",sep:"成",want:[]string{"都完","中..."}},		
	}
	//遍历切片，逐一执行测试用例
	for _,tc:=range tests{
		got:=Split(tc.input,tc.sep)
		if !reflect.DeepEqual(got,tc.want) {
			//t.Errorf("execpted:%v,got:%v",tc.want,got)
			t.Errorf("execpted:%#v,got:%#v",tc.want,got)//更明显的提示
		}
	}
}
/*
测试结果(有错误)：
=== RUN   TestChineseSplit
    split_test.go:79: execpted:[都完 中...],got:[ 都完 中...]
--- FAIL: TestChineseSplit (0.00s)
FAIL
exit status 1
FAIL    _/F_/www/go/20210224/pkg26      0.058s

我们的测试出现了问题，仔细看打印失败提示信息:你会发现参照集中有个不明显的空串，这种情况下十分推荐使用%#v的格式化方法。

三、子测试
看起来都挺不错的，但是如果测试用例比较多的时候，我们是没有办法一眼看出来具体是哪个测试用例失败，我们可能会想到以下办法:
*/
func TestChinese1Split(t *testing.T){
	//定义一个测试用例类型
	type test struct{
		input,sep string
		want []string
	}
	//定义一个储存测试用例的切片
	tests:=map[string]test{
		"simple":{input:"a:b:c",sep:":",want:[]string{"a","b","c"}},
		"wrong sep":{input:"a,b,c",sep:",",want:[]string{"a","b","c"}},
		"more sep":{input:"abcd",sep:"bc",want:[]string{"a","d"}},		
		//"leading sep":{input:"成都完成中...",sep:"成",want:[]string{"都完","中..."}},		
		"leading sep":{input:"成都完成中...",sep:"成",want:[]string{"","都完","中..."}},		
	}
	//遍历切片，逐一执行测试用例
	for name,tc:=range tests{
		got:=Split(tc.input,tc.sep)
		if !reflect.DeepEqual(got,tc.want) {
			//t.Errorf("execpted:%v,got:%v",tc.want,got)
			t.Errorf("name:%s,execpted:%#v,got:%#v",name,tc.want,got)//更明显的提示
		}
	}
}
/*
上面的做法可以解决该问题，同样的go1.7+中新增了子测试，我们可以按照如下方式使用t.Run()执行子测试

测试结果:
=== RUN   TestChineseSplit
--- PASS: TestChineseSplit (0.00s)
=== RUN   TestChinese1Split
    split_test.go:116: name:leading sep,execpted:[]string{"都完", "中..."},got:[]string{"", "都完", "中..."}
--- FAIL: TestChinese1Split (0.00s)
=== RUN   TestChinese2Split
=== RUN   TestChinese2Split/leading_sep
    split_test.go:141: name:leading sep,execpted:[]string{"都完", "中..."},got:[]string{"", "都完", "中..."}
=== RUN   TestChinese2Split/simple
=== RUN   TestChinese2Split/wrong_sep
=== RUN   TestChinese2Split/more_sep
--- FAIL: TestChinese2Split (0.00s)
    --- FAIL: TestChinese2Split/leading_sep (0.00s)
    --- PASS: TestChinese2Split/simple (0.00s)
    --- PASS: TestChinese2Split/wrong_sep (0.00s)
    --- PASS: TestChinese2Split/more_sep (0.00s)
FAIL
exit status 1
FAIL    _/F_/www/go/20210224/pkg26      0.054s

1、运行指定的子测试项
	go test -v -run="TestChinese2Split/more_sep"

*/
func TestChinese2Split(t *testing.T){
	//定义一个测试用例类型
	type test struct{
		input,sep string
		want []string
	}
	//定义一个储存测试用例的切片
	tests:=map[string]test{
		"simple":{input:"a:b:c",sep:":",want:[]string{"a","b","c"}},
		"wrong sep":{input:"a,b,c",sep:",",want:[]string{"a","b","c"}},
		"more sep":{input:"abcd",sep:"bc",want:[]string{"a","d"}},		
		"leading sep":{input:"成都完成中...",sep:"成",want:[]string{"","都完","中..."}},		
	}
	//遍历切片，逐一执行测试用例
	for name,tc:=range tests{
		t.Run(name,func(t *testing.T){//使用t.Run执行子测试
			got:=Split(tc.input,tc.sep)
			if !reflect.DeepEqual(got,tc.want) {
				t.Errorf("name:%s,execpted:%#v,got:%#v",name,tc.want,got)//更明显的提示
			}
		})
	}
}
/*
四、覆盖率测试
测试覆盖率是你的代码被测试套件覆盖的百分比，通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次占总代码的比例。
go test -cover
go test -cover -coverprofile=split.out //覆盖率输出到指定文件
go tool cover -html=split.out //使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个html报告，绿色代码块表示已覆盖，红色快代表未覆盖。

五、基准测试
	func BenchmarkName(b testing.B){}
基准测试默认不会执行
	go test -bench=Split //执行基准测试

测试结果:
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkSplit-8         5845838               217.0 ns/op
PASS
ok      _/F_/www/go/20210224/pkg26      1.538s

其中
BenchmarkSplit-8表示对Split函数进行基准测试，数据8表示GOMAXPROCS的值，这个对于并发基准测试很重要。
5845838和1.538s表示每次调用Split函数耗时1.538s,这个结果5845838次调用的平均值。


	go test -bench=Split -benchmem //获取内存分配的统计数据

测试结果:
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkSplit-8         6023971               195.4 ns/op           112 B/op          3 allocs/op
PASS
ok      _/F_/www/go/20210224/pkg26      1.437s

112 B/op表示每次操作内存分配112字节
3 allocs/op则表示每次操作进行了3次内存分配

*/
func BenchmarkSplit(b *testing.B){
	for i:=0;i<b.N;i++{
		Split("成都完成中","成")
	}
}
/*
优化Split函数
这一次我们提前使用make()将result初始化为一个容量足够大的切片，而不是像之前一样通过调用append函数来追加。我们来看一下这个改进性能会得到多大的提升：

测试结果:
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkSplit-8        10823467               108.6 ns/op            48 B/op          1 allocs/op
PASS
ok      _/F_/www/go/20210224/pkg26      1.345s

这个函数make提前分配内存的举动，减少了2/3的内存分配次数，并且减少了一半的内存分配。

1、性能比较函数
上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理一千个元素和耗时和处理一百万个元素的耗时的差别是多少？
再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。
性能比较函数通常带有一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用:

运行基准测试
	go test -bench=.
*/
func benchmarkFib(b *testing.B,n int){
	for i:=0;i<b.N;i++{
		Fib(n)
	}
}
func BenchmarkFib1(b *testing.B){ benchmarkFib(b,1) }
func BenchmarkFib2(b *testing.B){ benchmarkFib(b,2) }
func BenchmarkFib3(b *testing.B){ benchmarkFib(b,3) }
func BenchmarkFib10(b *testing.B){ benchmarkFib(b,10) }
func BenchmarkFib20(b *testing.B){ benchmarkFib(b,20) }
func BenchmarkFib40(b *testing.B){ benchmarkFib(b,40) }
/*
测试结果:

goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkSplit-8        10200663               109.7 ns/op
BenchmarkFib1-8         756791254                1.608 ns/op
BenchmarkFib2-8         248334554                4.785 ns/op
BenchmarkFib3-8         141808200                8.030 ns/op
BenchmarkFib10-8         3905816               304.4 ns/op
BenchmarkFib20-8           30724             38039 ns/op
BenchmarkFib40-8               2         574046200 ns/op
PASS
ok      _/F_/www/go/20210224/pkg26      11.143s

这里需要注意的是，默认情况下，每个基准测试至少运行1s，如果在Benchmark函数返回没有到达1s,则b.N的值会按1,2,5,10,20,50...增加，并且函数再次运行。
最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到1s.像这样的情况下我们应该可以使用-benchtime标志增加最小基准时间，已产生更准确的结果。例如:
	go test -bench=Fib40 -benchtime=20s

测试结果:

goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkFib40-8              40         577470758 ns/op
PASS
ok      _/F_/www/go/20210224/pkg26      23.746s

这一次的BenchmarkFib40函数运行了40次，结果就会更准确一些。

*使用性能比较函数做测试的时候一个容易犯的错误就是把B,N作为输入参数。

2、重置时间
b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。例如:
*/
func BenchmarkSplit1(b *testing.B){
	time.Sleep(5*time.Second)//假设需要做一些耗时的操作
	b.ResetTimer()//重置计时器
	for i:=0;i<b.N;i++{
		Split("沙河有沙又有河","沙")
	}
}
/*
3、并行测试
func (b *B) RunParallel(body func(*PB))会以并行的方式执行给定的基准测试。
RunParallel会创建出多个goroutine,并将b.N分配给这些foroutine执行，其中goroutine数量的默认值为GOMAXPROCS.用户如果想要增加非CPU受限(non-CPU-bound)基准测试的并行性，那么可以在RunParallel之前调用SetParallelism.RunParallel通常会与-cpu标志一同使用。

go test -bench=Parallel -cpu=1
可以通过命令行参数来设置CPU核心数
*/
func BenchmarkSplitParallel(b *testing.B){
	b.SetParallelism(1)//设置CPU的核心数
	b.RunParallel(func(pb *testing.PB){
		for pb.Next(){
			Split("沙河有沙子又有河","沙")
		}
	})
}
/*
4、TestMain
通过在*_test.go文件定义TestMain函数来可以在测试之前进行额外的设置(setup)或在测试之后进行拆卸(teardown)操作。
如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用TestMain(m),然后再运行具体测试。TestMain运行在主goroutine中，可以在调用m.Run前后做任何设置(setup)和(teardown).退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit().

例如：
需要注意的是，在调用TestMain时，flag.Parse并没有被调用,所以如果TestMain依赖于command-line标志(包括testing包的标记),则应该显示的调用flag.Parse.
 */
func TestMain(m *testing.M){
	fmt.Println("测试之前做的事情")
	//如果TestMain使用了flags,这里应该加上flag.Parse()
	retCode:=m.Run()//执行测试
	fmt.Println("测试之后做的事情")
	os.Exit(retCode)//退出测试
}

/*


5、子测试的Setup与Teardown
测试程序有时需要在测试之前进行额外的设置(setup)或在测试之后进行拆卸(teardown).
有时候我们可能需要为每个测试集设置Setup和Teardown,也有可能需要为每个子测试设置Setup和Teardown.

go test -v -run="Chinese3"
 */
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要再此执行:Setup前置任务")
	return func(t *testing.T){
		t.Log("如有需要再此执行:Teardown后置任务")
	}
}
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要再此执行:子测试Setup前置任务")
	return func(t *testing.T){
		t.Log("如有需要再此执行:子测试Teardown后置任务")
	}
}
func TestChinese3Split(t *testing.T){
	//定义一个测试用例类型
	type test struct{
		input,sep string
		want []string
	}
	//定义一个储存测试用例的切片
	tests:=map[string]test{
		"simple":{input:"a:b:c",sep:":",want:[]string{"a","b","c"}},
		"wrong sep":{input:"a,b,c",sep:",",want:[]string{"a","b","c"}},
		"more sep":{input:"abcd",sep:"bc",want:[]string{"a","d"}},		
		"leading sep":{input:"成都完成中...",sep:"成",want:[]string{"","都完","中..."}},		
	}
	teardownTestCase:=setupTestCase(t)//测试之前执行setup操作
	defer teardownTestCase(t) //测试周后执行的testdown操作

	//遍历切片，逐一执行测试用例
	for name,tc:=range tests{
		t.Run(name,func(t *testing.T){//使用t.Run执行子测试
			teardownSubTest:= setupSubTest(t)//子测试之前执行setup操作
			defer teardownSubTest(t) //子测试周后执行的testdown操作
			got:=Split(tc.input,tc.sep)
			if !reflect.DeepEqual(got,tc.want) {
				t.Errorf("name:%s,execpted:%#v,got:%#v",name,tc.want,got)//更明显的提示
			}
		})
	}
}
/*
六、示例函数
被go test特殊对待对的第三种就是示例函数，他们的函数名以Example为前缀，他们既没有参数也没有返回值。
func ExampleName(){}

为你的代码编写示例函数有三个用处:
1、示例函数作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联;
2、示例函数只要包含了//Output:也是可以通过go test运行的可执行测试;
	go test -v -run Example 
3、示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用GoPlayground运行示例代码.

func ExampleSplit(){
	fmt.Println(pkg26.Split("a:b:c",":"))
	fmt.Println(pkg26.Split("沙河有沙又有河","沙"))	
}
 */