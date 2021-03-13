package pkg24

import(
	"fmt"
	"runtime"
	"bufio"
	"net"
	"os"
	"io"
	"strings"
	"bytes"
	"encoding/binary"
)
//返回调用者函数名称
func getfunctionname()string{
	pc:=make([]uintptr,1)
	runtime.Callers(2,pc)
	return runtime.FuncForPC(pc[0]).Name()
}
/*
TCP协议
	TCP-server
	go build -o server.exe 24.go
*/
func process1(conn net.Conn){
	defer conn.Close()//关闭连接
	for {
		reader:=bufio.NewReader(conn)
		var buf [128]byte
		n,err:=reader.Read(buf[:])//读取数据
		if err!=nil{
			fmt.Println("读取客户端失败,err:",err)
			break
		}
		recvStr:=string(buf[:n])
		fmt.Println("收到client端发来的数据:",recvStr)
		conn.Write([]byte(recvStr))//发送数据
	}
}
func Test1(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	listen,err:=net.Listen("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	for{
		conn,err:=listen.Accept()//建立连接
		if err!=nil{
			fmt.Println("accept failed,err:",err)
			continue
		}
		go process1(conn)
	}
}
/*
TCP协议
	TCP-client
	go build -o client.exe 24.go

*/
func Test2(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	conn,err:=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("err:",err)
		return
	}
	defer conn.Close()//关闭连接
	inputReader:=bufio.NewReader(os.Stdin)
	for{
		input,_:=inputReader.ReadString('\n')//读取用户输入
		inputInfo:=strings.Trim(input,"\r\n")
		if strings.ToUpper(inputInfo)=="Q"{
			return //如果用户输入Q就退出
		}
		_,err=conn.Write([]byte(inputInfo))//发送数据
		if err!=nil{
			return
		}
		buf:=[512]byte{}
		n,err:=conn.Read(buf[:])
		if err!=nil{
			fmt.Println("接收失败,err:",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
/*
TCP黏包问题
*/
func process3(conn net.Conn){
	defer conn.Close()//关闭连接
	reader:=bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n,err:=reader.Read(buf[:])//读取数据
		if err==io.EOF{
			break
		}
		if err!=nil{
			fmt.Println("read from client failed,err:",err)
			break
		}
		recvStr:=string(buf[:n])
		fmt.Println("收到client端发来的数据:",recvStr)
	}
}
/*
server
*/
func Test3(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	listen,err:=net.Listen("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	defer listen.Close()
	for{
		conn,err:=listen.Accept()//建立连接
		if err!=nil{
			fmt.Println("accept failed,err:",err)
			continue
		}
		go process3(conn)
	}
}
/*
client
 */
func Test4(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	conn,err:=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("dial failed,err:",err)
		return
	}
	defer conn.Close()//关闭连接
	msg:=`Hello,Hello,How are you?`
	for i:=0;i<20;i++{
		conn.Write([]byte(msg))
	}
}
/*
粘包的解决办法
*/
//Encode将消息编码
func Encode(message string)([]byte,error){
	//读取消息的长度，转换成int32类型(占用4个字节)
	var length=int32(len(message))
	var pkg = new(bytes.Buffer)
	//写入消息头
	err:=binary.Write(pkg,binary.LittleEndian,length)
	if err!=nil{
		return nil,err
	}
	//写入消息实体
	err=binary.Write(pkg,binary.LittleEndian,[]byte(message))
	if err!=nil{
		return nil,err
	}
	return pkg.Bytes(),nil
}
//Decode解码消息
func Decode(reader *bufio.Reader)(string,error){
	//读取消息的长度
	lengthByte,_:=reader.Peek(4)//读取前四个字节
	lengthBuff:=bytes.NewBuffer(lengthByte)
	var length int32
	err:=binary.Read(lengthBuff,binary.LittleEndian,&length)
	if err!=nil{
		return "",err
	}
	//Buffered返回缓存中现有可读取的字节数
	if int32(reader.Buffered())<length+4{
		return "",err
	}
	//读取真正的消息数据
	pack:=make([]byte,int(4+length))
	_,err=reader.Read(pack)
	if err!=nil{
		return "",err
	}
	return string(pack[4:]),nil
}
/*
server
*/
func process5(conn net.Conn){
	defer conn.Close()
	reader:=bufio.NewReader(conn)
	for{
		msg,err:=Decode(reader)
		if err==io.EOF{
			return
		}
		if err!=nil{
			fmt.Println("decode msg failed,err:",err)
			return
		}
		fmt.Println("收到客户端发来的数据:",msg)
	}
}
func Test5(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	listen,err:=net.Listen("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	defer listen.Close()
	for{
		conn,err:=listen.Accept()//建立连接
		if err!=nil{
			fmt.Println("accept failed,err:",err)
			continue
		}
		go process5(conn)
	}
}
//client
func Test6(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	conn,err:=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("dial failed,err:",err)
		return
	}
	defer conn.Close()//关闭连接
	msg:=`Hello,Hello,How are you?`
	data,err:=Encode(msg)
	if err!=nil{
		fmt.Println("Encode msg failed,err:",err)
		return
	}
	for i:=0;i<20;i++{
		conn.Write(data)
	}
}
/*
使用net包实现的UDP通信

server
*/
func Test7(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	listen,err:=net.ListenUDP("udp",&net.UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port:20000,
		})
	if err!=nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	defer listen.Close()
	for{
		var data [1024]byte
		n,addr,err:=listen.ReadFromUDP(data[:])//接收数据
		if err!=nil{
			fmt.Println("read udp failed,err:",err)
			continue
		}
		fmt.Printf("data:%v,addr:%v count:%v\n",string(data[:n]),addr,n)
		_,err=listen.WriteToUDP(data[:n],addr)//发送数据
		if err!=nil{
			fmt.Println("write to udp failed,err:",err)
			continue
		}
	}
}
//client
func Test8(){
	fmt.Println("---------",getfunctionname(),"-------------")//打印函数名称
	socket,err:=net.DialUDP("udp",nil,&net.UDPAddr{
		IP:net.IPv4(0,0,0,0),
		Port:20000,
		})
	if err!=nil{
		fmt.Println("dialUDP failed,err:",err)
		return
	}
	defer socket.Close()//关闭连接
	sendData:=[]byte("Hello server")
	_,err=socket.Write(sendData)//发送数据
	if err!=nil{
		fmt.Println("发送数据失败,err:",err)
		return
	}
	data:=make([]byte,4096)
	n,remoteAddr,err:=socket.ReadFromUDP(data)//接收数据
	if err!=nil{
		fmt.Println("接收数据失败,err:",err)
		return
	}
	fmt.Printf("recv:%v,addr:%v count:%v\n",string(data[:n]),remoteAddr,n)
}