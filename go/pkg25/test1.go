package pkg25

import(
	"fmt"
	"runtime"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"strings"
)
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
GET请求示例
*/
func Test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	resp,err:=http.Get("http://19.19.19.70")
	if err!=nil{
		fmt.Println("get failed,err:",err)
		return
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("read from resp.Body,err:",err)
		return
	}
	fmt.Println(string(body))
}
/*
带参数GET请求示例
*/
func Test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	//apiUrl:="http://19.19.19.70/api.php/unlogin/test1"
	apiUrl:="http://127.0.0.1:8000/get"
	//Url Param
	data:=url.Values{}
	data.Set("name","liyb")
	data.Set("passwd","123456")
	u,err:=url.ParseRequestURI(apiUrl)
	if err!=nil{
		fmt.Println("parse url requestUrl failed,err:",err)
		return
	}
	u.RawQuery=data.Encode()//URL Encode
	fmt.Println(u.String())
	resp,err:=http.Get(u.String())
	if err!=nil{
		fmt.Println("post failed,err:",err)
		return
	}
	defer resp.Body.Close()
	b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("get resp failed,err:",err)
		return
	}
	fmt.Println(string(b))
}
/*
 对应Test2的server端如下
*/
func getHandler(w http.ResponseWriter,r *http.Request){
	defer r.Body.Close()
	data:=r.URL.Query()
	fmt.Printf("name:%s,passwd:%s\n",data.Get("name"),data.Get("passwd"))
	answer:=`{"status":"ok"}`
	w.Write([]byte(answer))
}
func Test2Server(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	http.HandleFunc("/get",getHandler)
	http.HandleFunc("/post",postHandler)
	err:= http.ListenAndServe(":8000",nil)
	if err!=nil {
		panic(err)
	}
	log.Println("Test2Server")
}

/*
带参数POST请求示例
*/
func Test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	apiUrl:="http://127.0.0.1:8000/post"
	/*
	表单数据提交
	contentType:="application/x-www-form-urlencoded"
	data:="name=liyb&passwd=123456"
	*/
	//json
	contentType:="application/json"
	data:=`{"name":"liyb","passwd","123456"}`
	resp,err:=http.Post(apiUrl,contentType,strings.NewReader(data))
	if err!=nil{
		fmt.Println("post failed,err:",err)
		return
	}
	defer resp.Body.Close()
	b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("get resp failed,err:",err)
		return
	}
	fmt.Println(string(b))
}
/*
 对应Test2的server端如下
*/
func postHandler(w http.ResponseWriter,r *http.Request){
	defer r.Body.Close()
	//1.请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)//打印form数据
	fmt.Println(r.PostForm.Get("name"),r.PostForm.Get("passwd"))
	//2.请求类型是application/json时从r.Body读取数据
	b,err:=ioutil.ReadAll(r.Body)																	
	if err!=nil{
		fmt.Println("read request.Body failed,err:",err)
		return
	}
	fmt.Println(string(b))
	answer:=`{"status":"ok"}`
	w.Write([]byte(answer))
}

/*
一、自定义client
要管理http的头域、重定向策略和其他设置，创建一个client

client:&http.Client{
	CheckRedirect:redirectPolicyFunc,
}
resp,err:=client.Get("http://example.com")
req.Header.Add("IF-None-Match",`W/"wyzzy"`)
resp,err:=client.Do(req)

二、自定义Transport
要管理代理、TLS配置、keep-alice、压缩和其他设置，创建一个Transport:

tr:=&http.Transport{
	TLSClientConfig:&tls.Config{RootCAs:pool}
	DisableCompression:true
}
client:=&http.Client{Transport:tr}
resp,err:=client.Get("http://example.com")

client和Transport类型都可以安全的被多个goroutine同时使用。处于效率考虑，应该一次建立、尽量复用。

三、自定义server
要管理服务器的行为，额可以创建一个自定义的server:
s:=&http.server{
	Addr:":8080",
	Handler:myHandler,
	ReadTimeout:10*time.Second,
	WriteTimeout:10*time.Second,
	MaxHeaderBytes:1<<20,
}
log.Fatal(s.ListenAndServe())
*/
