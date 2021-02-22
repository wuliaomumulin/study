#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <signal.h>
/*
一、多线程socket服务端
	1、创建线程
	int pthread_create
	2、线程的终止
		1)、线程的start_routine函数代码结束，自然消亡;
		2)、线程调用start_routine函数调用pthread_exit(void *retval)，参数默认填0即可结束;
		3)、被主进程或其他线程终止;

	多线程的应用
		封装Tcpserver
		g++ book37.cpp -o book37 -lpthread
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

	//指定fd向对端发送报文
	int Send(const int clientfd,const void *buf,const int buflen);

	//指定fd接受对端的报文
	int Recv(const int clientfd,void *buf,const int buflen);

	//关闭客户端的socket
	void CloseClient();

	//关闭用于监听的socket
	void CloseListen();

	~TcpServer();
};

//与客户端通信线程的主函数
void *pth_main(void *arg);
//SIGNINT和SIGTERM的处理函数
void Exit(int sig);

//声明一个全部类
TcpServer server;

int main(int argc,char *argv[]){
	


	if(argc!=2){printf("using:./server port\nExample:./server 5005\n");return -1;}

	//第一步，创建socket

	for (int i = 0; i < 50; ++i) signal(i,SIG_IGN);//忽略全部信号
	//设置SIGNINT和SIGNTERM的处理函数
	signal(SIGINT,Exit);signal(SIGTERM,Exit);



	if(server.InitServer(atoi(argv[1]))==false){printf("服务器初始化失败，端口号:%d,程序退出",atoi(argv[1]));return -1;}
	


	while(1){
		if(server.Accept()==false) continue;

		pthread_t pthid;//创建一条线程，与新连接上来的客户端通信
		/**
		int pthread_create(pthread_t *thread, const pthread_attr_t *attr,
        void *(*start_routine) (void *), void *arg)
		参数thread为指向线程标识符的地址
		参数attr用户设置线程属性，一般为空，表示使用默认属性
		参数start_rounine是线程运行函数的地址，填函数名就行
		蚕食arg是线程运行函数的参数，新创建的线程从start_rountine函数的地址开始运行，该函数只有一个无类型指针参数arg。若要向start_routine传递多个参数，可以将多个参数放在一个结构体中，然后把结构体的地址作为arg参数传入。
		*/
		if(pthread_create(&pthid,NULL,pth_main,(void*)((long)server.m_clientfd))!=0){printf("创建线程失败，程序退出。\n");return -1;}

		printf("与客户端通信的线程已建立。\n");
	}

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
//指定fd向对端发送报文
int TcpServer::Send(const int clientfd,const void *buf,const int buflen){
	return send(clientfd,buf,buflen,0);
}

//指定fd接受对端的报文
int TcpServer::Recv(const int clientfd,void *buf,const int buflen){
	return recv(clientfd,buf,buflen,0);
}
//关闭客户端的socket
void TcpServer::CloseClient(){
	if(m_clientfd!=0){close(m_clientfd);m_clientfd=0;}
}

//关闭用于监听的socket
void TcpServer::CloseListen(){
	if(m_listenfd!=0){close(m_listenfd);m_listenfd=0;}
}


//SIGNINT和SIGTERM的处理函数
void Exit(int sig){
	printf("程序退出，信号值:%d\n",sig);
	//在类外使用close(TcpServer.m_listenfd)
	close(server.m_listenfd);
	exit(0);
}
//与客户端通信线程的主函数
void *pth_main(void *arg){

	int clientfd=(long) arg;//arg参数为新客户端的socket
	//与客户端通信之后，接收客户端发过来的报文后，回复ok
	char buffer[1024];

	while(1){
		memset(buffer,0,sizeof(buffer));
		if(server.Recv(clientfd,buffer,sizeof(buffer))<=0) break;
		printf("接收:%s\n",buffer);
		
		strcpy(buffer,"ok");

		if(server.Send(clientfd,buffer,strlen(buffer))<=0) break;

		printf("发送:%s\n",buffer);

	}

	printf("客户端已断开连接\n");
	close(clientfd);//关闭你客户端连接
	//线程主函数中不能使用return语句，可以使用下方返回
	pthread_exit(0);

}
/**
查看线程方式:
	top -H
	ps xH|grep book37

*/