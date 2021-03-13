package main
import(
	pkg "./pkg25"
)
/*
net/http
# 配置国内镜像源
go env -w GO111MODULE=on
go env -w GOPROXY=http://192.168.1.7:8081/repository/aliyun-go/
go env -w GOSUMDB=off
*/
func main(){
	//pkg.Test1()
	//pkg.Test2()
	
	//pkg.Test2Server()
	pkg.Test3()
	

}



