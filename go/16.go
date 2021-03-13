package main
import(
	"fmt"
	"runtime"
	"html/template"
	"net/http"
	"io/ioutil"
)
/*
	html/template包实现了数据驱动的模版，用于生成可防止代码注入的安全html内容。他提供了和text/template包相同的接口，go语言中输出HTML的场景都应使用html/template这个包

模板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码。
模板文件中使用{{和}}包裹和标识需要传入的数据。
传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
除{{和}}包裹的内容外，其他内容均不做修改原样输出。
*/
func main(){
	test1()

}
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
	sayhello
*/
func sayhello(w http.ResponseWriter,r *http.Request){
	//解析指定文件，生成模版对象
	tmpl,err:=template.ParseFiles("./tmpl/sayhello.tmpl")
	if err!=nil {
		fmt.Println("创建模版失败,err:",err)
		return
	}
	//利用给定数据渲染模版，并将结果写入w
	tmpl.Execute(w,"法外狂徒张三")	
}
/*
访问http://127.0.0.1:9090，这是一个最简单的实例
 */

func test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	http.HandleFunc("/",sayhello)
	http.HandleFunc("/sayhello1",sayhello1)
	http.HandleFunc("/sayhello2",sayhello2)
	http.HandleFunc("/sayhello3",sayhello3)
	http.HandleFunc("/sayhello4",sayhello4)
	http.HandleFunc("/sayhello5",sayhello5)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		fmt.Println("http server failed,err:",err)
		return	
	}
}
/*
	模版语法都包含在{{和}}之间，其中{{.}}中的点表示当前对象
	当我们传入一个结构体数对象时，我们可以根据.来访问结构体的对应字段。例如
*/
type UserInfo struct{
	Name string
	Gender string
	Age int
}
func sayhello1(w http.ResponseWriter,r *http.Request){
	//解析指定文件，生成模版对象
	tmpl,err:=template.ParseFiles("./tmpl/sayhello_1.tmpl")
	if err!=nil {
		fmt.Println("创建模版失败,err:",err)
		return
	}
	//利用给定数据渲染模版，并将结果写入w
	user := UserInfo{
		Name:"法外狂徒张三",
		Gender:"男",
		Age:26,
	}
	tmpl.Execute(w,user)	
}
//自定义函数
func sayhello2(w http.ResponseWriter,r *http.Request){
	htmlByte,err:=ioutil.ReadFile("./tmpl/sayhello_2.tmpl")
	if err!=nil {
		fmt.Println("文法读取sayhello_2.tmpl,err:",err)
		return
	}
	kua:=func(arg string)(string,error){
		return arg+",安帝科技",nil
	}
	//采用链式操作在Parse之前调用Func添加自定义的函数
	tmpl,err:=template.New("sayhello_2").Funcs(template.FuncMap{"kua":kua}).Parse(string(htmlByte))
	if err!=nil {
		fmt.Println("创建模版失败,err:",err)
		return
	}
	//利用给定数据渲染模版，并将结果写入w
	user := UserInfo{
		Name:"李四",
		Gender:"女",
		Age:28,
	}
	tmpl.Execute(w,user)	
}
/*
 模版嵌套
 */
func sayhello3(w http.ResponseWriter,r *http.Request){
	//解析指定文件，生成模版对象
	tmpl,err:=template.ParseFiles("./tmpl/t.tmpl","./tmpl/ul.tmpl")
	if err!=nil {
		fmt.Println("创建模版失败,err:",err)
		return
	}
	//利用给定数据渲染模版，并将结果写入w
	user := UserInfo{
		Name:"法外狂徒张三",
		Gender:"男",
		Age:26,
	}
	tmpl.Execute(w,user)	
}
/*
	模版继承
 */
func sayhello4(w http.ResponseWriter,r *http.Request){
	//解析指定文件，生成模版对象
	tmpl,err:=template.ParseGlob("./tmpl/*block.tmpl")
	if err!=nil {
		fmt.Println("创建模版失败,err:",err)
		return
	}
	err=tmpl.ExecuteTemplate(w,"block.tmpl",nil)
	if err!=nil {
		fmt.Println("render继承模版失败,err:",err)
		return
	}
}
/*
	xss
 */
func sayhello5(w http.ResponseWriter,r *http.Request){
	//解析指定文件，生成模版对象
	tmpl,err:=template.New("sayhello_5.tmpl").Funcs(template.FuncMap{"safe":func(s string)template.HTML {return template.HTML(s)},}).ParseFiles("./tmpl/sayhello_5.tmpl")
	if err!=nil {
		fmt.Println("创建模版失败,err:",err)
		return
	}
	jsStr:=`<script>alert("江畔何人初见月")</script>`
	err=tmpl.Execute(w,jsStr)
	if err!=nil {
		fmt.Println("Execute",err)
	}
}