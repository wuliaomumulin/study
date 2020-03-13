#include <stdio.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <strings.h>
#include <string.h>
#include <netinet/in.h>
#include <stdlib.h>
#include <unistd.h>

#include <iostream>

//域套接字定义
#define SOCKET_PATH "./domainsocket"
#define MSG_SIZE 2048

int main(){
	int socket_fd,accept_fd;
	int ret = 0;
	socklen_t addr_len;
	char msg[MSG_SIZE];
	struct sockaddr_un server_addr;

	//1、创建域套接字
	socket_fd = socket(PF_UNIX,SOCK_STREAM,0);
	if(-1 == socket_fd){
		std::cout << "套接字创建失败" << std::endl;
		return -1;
	}

	//内存区域置零
	bzero(&server_addr,sizeof(server_addr));
	server_addr.sun_family = PF_UNIX;
	strcpy(server_addr.sun_path,SOCKET_PATH);

	//2、连接域套接字
	std::cout << "2、连接套接字..." << std::endl;
	ret = connect(socket_fd,(sockaddr *)&server_addr,sizeof(server_addr));

	if(0 > ret){
		std::cout << "连接套接字失败" << std::endl;
		return -1;
	}

	while(true){
		//键入信息
		std::cout << "键入信息>>>";
		fgets(msg,MSG_SIZE,stdin);
		//3、发送消息
		ret = send(socket_fd,msg,MSG_SIZE,0);	
	}

	close(socket_fd);
	return 0;
}
//g++ client.cpp -o client -g -lpthread
