package main
import(
	"fmt"
	"net/http"
)

//上传文件
func FileUploadHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello,world")
	if(r.Method == "GET"){
		//方法获取上传主页

	}else{
		//POST方法获取上传内容
		
	}
}

func IndexHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello,world")
}
/*
* 文件查询接口
* 
*/
func main(){
	http.HandleFunc("/",FileUploadHandler)
	http.ListenAndServe("127.0.0.1:80",nil)
}