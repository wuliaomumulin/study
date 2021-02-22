#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>




int main(int argc,char *argv[]){
	
	if(argc!=3){printf("using:./client ip port\nExample:./client 19.19.19.70 5005\n");return -1;}

	//第一步，创建socket
	int sockfd;
	if((sockfd = socket(AF_INET,SOCK_STREAM,0))==-1){perror("socket");return -1;}
	
	//第二步，向服务端发起连接请求
	struct hostent* h;
	//指定服务端的ip地址
	if((h=gethostbyname(argv[1]))==0){printf("gethostbyname failed.\n");close(sockfd);return -1;}
	//服务端地址信息的数据结构
	struct sockaddr_in servaddr;
	memset(&servaddr,0,sizeof(servaddr));
	servaddr.sin_family = AF_INET;//协议簇，在socket编程中只能是AF_INET
	memcpy(&servaddr.sin_addr,h->h_addr,h->h_length);
	servaddr.sin_port = htons(atoi(argv[2]));//指定通信端口
	if(connect(sockfd,(struct sockaddr *)&servaddr,sizeof(servaddr))!=0){
		perror("connect");close(sockfd);return -1;
	}
	//第三步,与服务端通信，发送一个报文后等待回复，然后再发下一个报文
	char buffer[1024];
	for (int i = 0; i < 3; ++i)
	{
		int iret;
		memset(buffer,0,sizeof(buffer));
		sprintf(buffer,"这是第%d次握手,编号%03d",i+1,i+1);

		if((iret=send(sockfd,buffer,strlen(buffer),0))<=0){
			//向服务端发送报文
			perror("send");break;
		}
		printf("发送:%s\n",buffer);

		memset(buffer,0,sizeof(buffer));

		if((iret=recv(sockfd,buffer,sizeof(buffer),0))<=0){
			//接收服务端的回应报文
			printf("iret=%d\n",iret);break;
		}
		printf("接收:%s\n",buffer);
		
	}

	//第四步，关闭socket,释放资源
	close(sockfd);

}
//g++ book28.cpp -o book28