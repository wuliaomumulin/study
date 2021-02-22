#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>


/**
	封装socket通讯类

	server
*/

int initserver(int port);

int main(int argc,char *argv[]){
	
	if(argc!=2){printf("using:./server port\nExample:./server 5005\n");return -1;}

	//第一步，创建socket
	int listenfd = initserver(atoi(argv[1]));
	if(listenfd <= 0){printf("服务器初始化失败，程序退出");return -1;}
	
	//第四步，接收客户端的连接
	int clientfd;//客户端的socket
	if((clientfd=accept(listenfd,0,0))<=0){printf("服务器accept失败，程序退出");return -1;}

	printf("客户端已连接\n");

	//与客户端通信，接收客户端发过来的报文后，恢复ok
	char buffer[1024];
	while(1){
		memset(buffer,0,sizeof(buffer));
		if(recv(clientfd,buffer,sizeof(buffer),0)<=0) break;
		printf("接收:%s\n",buffer);
		
		strcpy(buffer,"ok");

		if(send(clientfd,buffer,strlen(buffer),0)<=0) break;

		printf("发送:%s\n",buffer);

	}

	//第六步，关闭socket,释放资源
	printf("客户端已断开连接\n");

	close(listenfd);close(clientfd);

}
/*
	初始化服务端socket，port为端口
	返回值，成功返回socket,失败返回-1
*/
int initserver(int port){
	int listenfd = socket(AF_INET,SOCK_STREAM,0);//创建socket服务
	//把服务端用于通讯的地址和端口绑定到socket上
	struct sockaddr_in servaddr;//服务器地址信息的数据结构
	memset(&servaddr,0,sizeof(servaddr));
	servaddr.sin_family = AF_INET;//协议簇，在socket编程中只能是AF_INET
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY);//本机的任意IP地址
	servaddr.sin_port = htons(port);//指定通信端口

	if(bind(listenfd,(struct sockaddr *)&servaddr,sizeof(servaddr))!=0){
		perror("bind");close(listenfd);return -1;
	}

	//第三步,把socket设置为监听模式
	if(listen(listenfd,5)!=0){perror("listen");close(listenfd);return -1;}

	return listenfd;
}
//gcc book29.c -o book29