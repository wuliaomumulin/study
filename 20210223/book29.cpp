#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>


/**
	封装Tcpserver

*/
class TcpServer{
public:
	int m_listenfd;//服务端用于监听的socket
	int m_clientfd;//客户端连接上来的socket

	TcpServer();

	bool InitServer(int port);//初始化服务端

	bool Accept();//等待客户端的连接

	//向对端发送报文
	int Send(const void *buf,const int buflen);

	//接受对端的报文
	int Recv(void *buf,const int buflen);

	~TcpServer();
};


int main(int argc,char *argv[]){
	
	if(argc!=2){printf("using:./server port\nExample:./server 5005\n");return -1;}

	//第一步，创建socket
	TcpServer server;
	if(server.InitServer(atoi(argv[1]))==false){printf("服务器初始化失败，端口号:%d,程序退出",atoi(argv[1]));return -1;}
	
	if(server.Accept()==false){printf("服务器accept失败，程序退出");return -1;}

	printf("客户端已连接\n");

	//与客户端通信，接收客户端发过来的报文后，恢复ok
	char buffer[1024];
	while(1){
		memset(buffer,0,sizeof(buffer));
		if(server.Recv(buffer,sizeof(buffer))<=0) break;
		printf("接收:%s\n",buffer);
		
		strcpy(buffer,"ok");

		if(server.Send(buffer,strlen(buffer))<=0) break;

		printf("发送:%s\n",buffer);

	}

	//第六步，关闭socket,释放资源
	printf("客户端已断开连接\n");

}
/*
	封装Tcpserver
*/
TcpServer::TcpServer(){
	//构造函数初始化socket
	m_listenfd=m_clientfd=0;
}
TcpServer::~TcpServer(){
	//析构函数关闭socket
	if(m_listenfd!=0) close(m_listenfd);
	if(m_clientfd!=0) close(m_clientfd);
	
}
//初始化服务端的socket,port为通讯端口
bool TcpServer::InitServer(int port){
	m_listenfd = socket(AF_INET,SOCK_STREAM,0);//创建socket服务
	//把服务端用于通讯的地址和端口绑定到socket上
	struct sockaddr_in servaddr;//服务器地址信息的数据结构

	memset(&servaddr,0,sizeof(servaddr));
	servaddr.sin_family = AF_INET;//协议簇，在socket编程中只能是AF_INET
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY);//本机的任意IP地址
	servaddr.sin_port = htons(port);//指定通信端口

	if(bind(m_listenfd,(struct sockaddr *)&servaddr,sizeof(servaddr))!=0){
		close(m_listenfd);m_listenfd=0;return -false;
	}

	//第三步,把socket设置为监听模式
	if(listen(m_listenfd,5)!=0){close(m_listenfd);m_listenfd=0;return false;}

	return true;

}
bool TcpServer::Accept(){
	if((m_clientfd=accept(m_listenfd,0,0))<=0){perror("accpet");printf("clientfd=%d,listenfd=%d\n",m_clientfd,m_listenfd); return false;}

	return true;
}
int TcpServer::Send(const void *buf,const int buflen){
	return send(m_clientfd,buf,buflen,0);
}

//接受对端的报文
int TcpServer::Recv(void *buf,const int buflen){
	return recv(m_clientfd,buf,buflen,0);
}
//g++ book29.cpp -o book29