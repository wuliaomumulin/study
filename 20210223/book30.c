#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>


int connecttoserver(const char *serverip,const int port);

int main(int argc,char *argv[]){
	
	if(argc!=3){printf("using:./client ip port\nExample:./client 19.19.19.70 5005\n");return -1;}

	//第一步，创建socket
	int sockfd = connecttoserver(argv[1],atoi(argv[2]));
	if(sockfd<=0){printf("服务器连接失败\n");return -1;}

	//与服务端通信，发送一个报文后等待回复，然后再发下一个报文
	char buffer[1024];
	for (int i = 0; i < 10; ++i)
	{

		memset(buffer,0,sizeof(buffer));
		sprintf(buffer,"这是第%d次握手,编号%03d",i+1,i+1);

		if(send(sockfd,buffer,strlen(buffer),0)<=0) break;
		printf("发送:%s\n",buffer);

		memset(buffer,0,sizeof(buffer));

		if(recv(sockfd,buffer,sizeof(buffer),0)<=0) break;
		printf("接收:%s\n",buffer);
		
	}

	//第四步，关闭socket,释放资源
	close(sockfd);

}
/*
	TCP客户端连接服务端的函数，serverip-服务端ip,port-端口号
	返回值，成功返回socket,失败返回-1
*/
int connecttoserver(const char *serverip,const int port){
	int sockfd = socket(AF_INET,SOCK_STREAM,0);//创建socket服务
	struct hostent* h;//ip地址的信息数据结构
	if((h=gethostbyname(serverip))==0){perror("gethostbyname");close(sockfd);return -1;}

	//把服务器的地址和端口转换为数据结构
	struct sockaddr_in servaddr;//服务器地址信息的数据结构
	memset(&servaddr,0,sizeof(servaddr));
	servaddr.sin_family = AF_INET;//协议簇，在socket编程中只能是AF_INET
	servaddr.sin_port = htons(port);//指定通信端口
	memcpy(&servaddr.sin_addr,h->h_addr,h->h_length);

	if(connect(sockfd, (struct sockaddr *)&servaddr,sizeof(servaddr))!=0){
		perror("connect");close(sockfd);return -1;
	}

	return sockfd;
}
//gcc book30.c -o book30