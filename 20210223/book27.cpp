#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <netdb.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <arpa/inet.h>




int main(int argc,char *argv[]){
	
	if(argc!=2){printf("using:./server port\nExample:./server 5005\n");return -1;}

	//第一步，创建socket
	int listenfd;
	if((listenfd = socket(AF_INET,SOCK_STREAM,0))==-1){perror("socket");return -1;}
	
	//第二步，要把服务端用于通信的地址和端口绑定到socket上
	struct sockaddr_in servaddr;//服务端地址信息的数据结构
	memset(&servaddr,0,sizeof(servaddr));
	servaddr.sin_family = AF_INET;//协议簇，在socket编程中只能是AF_INET
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY);//任意IP地址
	//servaddr.sin_addr.s_addr = inet_addr("19.19.19.70");//指定ip地址
	servaddr.sin_port = htons(atoi(argv[1]));//指定通信端口
	if(bind(listenfd,(struct sockaddr *)&servaddr,sizeof(servaddr))!=0){
		perror("bind");close(listenfd);return -1;
	}
	//第三步,把socket设置为监听模式
	if(listen(listenfd,5)!=0){perror("listen");close(listenfd);return -1;}


	//第四步，接收客户端的连接
	int clientfd;//客户端的socket
	int socklen=sizeof(struct sockaddr_in);
	struct sockaddr_in clientaddr;//客户端的地址信息
	clientfd=accept(listenfd,(struct sockaddr *)&clientaddr,(socklen_t*)&socklen);
	printf("客户端已连接\n",inet_ntoa(clientaddr.sin_addr));

	//第五步，与客户端通信，接收客户端发过来的报文后，恢复ok
	char buffer[1024];
	while(1){
		int iret;
		memset(buffer,0,sizeof(buffer));
		if((iret=recv(clientfd,buffer,sizeof(buffer),0))<=0){
			//接收客户端的请求报文
			printf("iret=%d\n",iret);break;
		}
		printf("接收:%s\n",buffer);
		
		strcpy(buffer,"ok");

		if((iret=send(clientfd,buffer,strlen(buffer),0))<=0){
			//向客户端发送响应结果
			perror("send");break;
		}

		printf("发送:%s\n",buffer);

	}

	//第六步，关闭socket,释放资源
	close(listenfd);close(clientfd);

}
//g++ book27.cpp -o book27