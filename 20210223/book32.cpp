#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <signal.h>

/**
 多进程的应用
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

	//关闭客户端的socket
	void CloseClient();

	//关闭用于监听的socket
	void CloseListen();

	~TcpServer();
};


int main(int argc,char *argv[]){
	
	signal(SIGCHLD,SIG_IGN);//忽略子进程退出的信号，避免出现僵尸进程

	if(argc!=2){printf("using:./server port\nExample:./server 5005\n");return -1;}

	//第一步，创建socket
	TcpServer server;
	if(server.InitServer(atoi(argv[1]))==false){printf("服务器初始化失败，端口号:%d,程序退出",atoi(argv[1]));return -1;}
	
	while(1){
		if(server.Accept()==false) continue;

		//父进程回到while,继续accept
		if(fork()>0){server.CloseClient();continue;}

		//子进程负责与客户端进行通信，直到客户端断开连接
		server.CloseListen();

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

		return 0;//或者exit(0),子进程退出

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
//关闭客户端的socket
void TcpServer::CloseClient(){
	if(m_clientfd!=0){close(m_clientfd);m_clientfd=0;}
}

//关闭用于监听的socket
void TcpServer::CloseListen(){
	if(m_listenfd!=0){close(m_listenfd);m_listenfd=0;}
}
/*

g++ book32.c -o book32

1、注意，服务端的while主程序是一个死循环，没有退出机制，可以按ctrl+c强制终止它，这不是正确的办法，后面介绍正确的办法

2、僵尸进程
	ps -ef|grep book32|grep -v grep
	(1)、带有<defunct>标识的就是僵尸进程，如果按ctrl+c终止程序，父进程退出，僵尸进程随之消失
	(2)、僵尸进程的危害
		a、僵尸进程是子进程结束后，父进程没有回收子进程占用的资源;
		b、僵尸进程在消失之前会继续占用系统资源;
		c、如果父进程先退出，子进程被系统接管，子进程退出后系统会回收其占用的相关资源，不会成为僵尸进程。父进和先退出的应用场景在以后章节介绍;
	(3)、解决僵尸进程
		a、子进程退出之前，回向父进程发送一个信号，父进程调用wait函数等待这个信号，直到等到了，就不会产生僵尸进程。这话说的容易，在并发的服务程序中这是不可能的，因为父进程要做其他的事，例如等待客户端的新连接，不可能去等待子进程的退出信号，这个方法就不介绍了;
		b、另一种方法就是父进程忽略子进程的退出信号
			signal(SIGCHLD,SIG_IGN);//忽略子进程退出的信号，避免出现僵尸进程

*/