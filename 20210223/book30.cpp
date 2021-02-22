#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>


/**
	封装TcpClient

*/
class TcpClient{
public:
	int m_sockfd;
	
	TcpClient();

	//向服务端发起连接,serverip=服务端ip,port=通讯端口
	bool connecttoserver(const char *serverip,const int port);
	//向对端发送报文
	int Send(const void *buf,const int buflen);
	//接受对端的报文
	int Recv(void *buf,const int buflen);

	~TcpClient();
};

int main(int argc,char *argv[]){
	
	if(argc!=3){printf("using:./client ip port\nExample:./client 19.19.19.70 5005\n");return -1;}

	//第一步，创建socket
	TcpClient client;

	if(client.connecttoserver(argv[1],atoi(argv[2]))==0){printf("服务器连接失败\n");return -1;}

	//与服务端通信，发送一个报文后等待回复，然后再发下一个报文
	char buffer[1024];
	for (int i = 0; i < 10; ++i)
	{

		memset(buffer,0,sizeof(buffer));
		sprintf(buffer,"这是第%d次握手,编号%03d",i+1,i+1);

		if(client.Send(buffer,strlen(buffer))<=0) break;
		printf("发送:%s\n",buffer);

		memset(buffer,0,sizeof(buffer));

		if(client.Recv(buffer,sizeof(buffer))<=0) break;
		printf("接收:%s\n",buffer);
		
	}

}
/*
	定义TcpClient
*/
TcpClient::TcpClient(){
	m_sockfd = 0; //构造函数初始化m_sockfd
}
TcpClient::~TcpClient(){
	if(m_sockfd!=0) close(m_sockfd);//析构函数关闭m_sockfd
}
/*
	TCP客户端连接服务端的函数，serverip-服务端ip,port-端口号
	返回值，成功返回true,失败返回false
*/
bool TcpClient::connecttoserver(const char *serverip,const int port){
	m_sockfd = socket(AF_INET,SOCK_STREAM,0);//创建socket服务
	struct hostent* h;//ip地址的信息数据结构
	if((h=gethostbyname(serverip))==0){perror("gethostbyname");close(m_sockfd);return false;}

	//把服务器的地址和端口转换为数据结构
	struct sockaddr_in servaddr;//服务器地址信息的数据结构
	memset(&servaddr,0,sizeof(servaddr));
	servaddr.sin_family = AF_INET;//协议簇，在socket编程中只能是AF_INET
	servaddr.sin_port = htons(port);//指定通信端口
	memcpy(&servaddr.sin_addr,h->h_addr,h->h_length);

	if(connect(m_sockfd, (struct sockaddr *)&servaddr,sizeof(servaddr))!=0){
		perror("connect");close(m_sockfd);return false;
	}

	return true;
}

	//向对端发送报文
int TcpClient::Send(const void *buf,const int buflen){
	return send(m_sockfd,buf,buflen,0);
}
//接受对端的报文
int TcpClient::Recv(void *buf,const int buflen){
	return recv(m_sockfd,buf,buflen,0);
}
//g++ book30.c -o book30