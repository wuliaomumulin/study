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

	//移除已有套接字路径
	remove(SOCKET_PATH);
	//内存区域置零
	bzero(&server_addr,sizeof(server_addr));
	server_addr.sun_family = PF_UNIX;
	strcpy(server_addr.sun_path,SOCKET_PATH);

	//2、绑定域套接字
	std::cout << "绑定套接字..." << std::endl;
	ret = bind(socket_fd,(sockaddr *)&server_addr,sizeof(server_addr));

	if(0 > ret){
		std::cout << "绑定套接字失败" << std::endl;
		return -1;
	}

	//3、监听套接字
	std::cout << "监听套接字..." << std::endl;
	ret = listen(socket_fd,10);
	if(-1 == ret){
		std::cout << "监听套接字失败" << std::endl;
		return ret;
	}
	std::cout << "等待新的请求." << std::endl;
	accept_fd = accept(socket_fd,NULL,NULL);

	bzero(msg,MSG_SIZE);

	//true必须小写
	while(true){
		//4、接受&处理信息
		recv(accept_fd,msg,MSG_SIZE,0);
		std::cout << "接受信息：" << msg << std::endl;
	}

	close(accept_fd);
	close(socket_fd);
	return 0;
}
//g++ server.cpp -o server -g -lpthread