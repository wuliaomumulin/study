#include "common.h"

#include <sys/shm.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>

#include <iostream>

int main(){
	
	//定义共享内存结构体
	struct ShmEntry *entry;

	//1、申请共享内容,参数为key,size和shmfig[IPC_CREAT|(IPC_CREAT|IPC_EXCL)],返回一个共享内存的id
	int shmid = shmget((key_t)1111,sizeof(struct ShmEntry),0666|IPC_CREAT);
	if(shmid == -1){
		std::cout << "创建共享内存失败" << std::endl;
		return -1;
	}

	//2、连接到当前进程空间,使用共享内存
	//参数共享内存标识符、shmid指定共享内存出现在进程内存地址的什么位置，shmaddr直接指定为NULL让内核自己决定一个合适的地址位置，shmflg读写模式
	entry = (ShmEntry*)shmat(shmid,0,0);
	entry->can_read = 0;
	while(true){
		if(entry->can_read == 1){
			std::cout << "收到信息:" << entry->msg << std::endl;
			entry->can_read = 0;
		}else{
			std::cout << "内存块不可读，休息中1s.." << std::endl;
			sleep(1);
		}
	}

	//3、脱离进程空间
	//是用来断开与共享内存附加点的地址，禁止本进程访问此片共享内存
	shmdt(entry);

	//4、删除共享内存
	shmctl(shmid,IPC_RMID,0);

	return 0;
}
//g++ server.cpp common.h -o server -g -lpthread