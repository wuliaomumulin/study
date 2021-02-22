#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>




/**
	封装线程池TcpClient

*/


/*
一、线程同步-----锁
	g++ book37.cpp -o book37 -lpthread

*	可以尝试把线程锁关闭，可以发现客户端的两个线程的报文收发出现了混乱

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


//int pthread_mutex_t mutex;//声明一个互斥锁
pthread_mutex_t mutex;//声明一个互斥锁


TcpClient client;

void *pth_main(void *arg){
	int pno=(long)arg;//线程编号

	pthread_detach(pthread_self());

	char buffer[1024];
	for (int i = 0; i < 3; ++i)
	{
		//int pthread_mutex_lock(&mutex);//加锁
		pthread_mutex_lock(&mutex);//加锁

		memset(buffer,0,sizeof(buffer));
		sprintf(buffer,"这是第%d次握手,编号%03d",i+1,i+1);

		if(client.Send(buffer,strlen(buffer))<=0) break;
		printf("线程%d发送:%s\n",pno,buffer);

		memset(buffer,0,sizeof(buffer));

		if(client.Recv(buffer,sizeof(buffer))<=0) break;
		printf("线程%d接收:%s\n",pno,buffer);

		//int pthread_mutex_unlock(&mutex);//释放锁
		//usleep(100);//否则其他的线程无法获取锁
		pthread_mutex_unlock(&mutex);//释放锁
		usleep(100);//否则其他的线程无法获取锁
	}

	pthread_exit(0);
}



int main(int argc,char *argv[]){
	
	if(argc!=3){printf("using:./client ip port\nExample:./client 19.19.19.70 5005\n");return -1;}

	if(client.connecttoserver(argv[1],atoi(argv[2]))==0){printf("服务器连接失败\n");return -1;}

	//int pthread_mutex_init(&mutex,0);//创建锁
	pthread_mutex_init(&mutex,0);//创建锁
	
	pthread_t pthid1,pthid2;
	pthread_create(&pthid1,NULL,pth_main,(void*)1);//创建第一个线程
	pthread_create(&pthid2,NULL,pth_main,(void*)2);//创建第二个线程

	pthread_join(pthid1,NULL);//等待线程1退出
	pthread_join(pthid2,NULL);//等待线程2退出

	//int pthread_mutex_lock(&mutex);//销毁锁
	pthread_mutex_lock(&mutex);//销毁锁
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
