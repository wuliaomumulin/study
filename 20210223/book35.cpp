#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <sys/ipc.h>
#include <sys/shm.h>
/*
	一、共享内存
	
	1、查看系统的共享内存
		ipcs -m
			key            shmid owner perms bytes nattch status
			十六进制的键值  标识   拥有者 权限   大小          状态
	2、手工删除共享内存
		ipcrm -m 标识

*/


int mem(){
	int shmid;//共享内存标识
	//创建共享内存，键值为0x5005,共1024字节,权限为0640,IPC_CREAT表示全部用户对他可读写，如果共享内存不存在，就创建一块共享内存
	if((shmid=shmget((key_t)0x5005,1024,0640|IPC_CREAT))==-1){printf("shmget(0x5005) failed\n");return -1;}

	char *ptext=0;//用户指向共享内存的指针

	/*
		将共享内存连接到当前进程的地址空间，有ptext指针指向它
		shmid：共享内存标识
		shm_addr:指定共享内存连接到当前进程中的地址位置，通常为空，表示让系统来选择共享内存的地址
		shm_flg:一组标识位，通常为0;
	*/
	ptext = (char *)shmat(shmid,0,0);

	//操作本程序的ptext指针就是从左共享内存
	printf("写入前:%s\n",ptext);
	sprintf(ptext,"本程序的进程号是:%d",getpid());
	printf("写入后:%s\n",ptext);


	/*
		把共享内存从当前进程中分离
		shmaddr:是shmat函数返回的地址
		调用成功返回0，失败返回-1
	*   用root创建的共享内存，不管创建的权限是什么，普通用户无法删除;
	*/
	shmdt(ptext);

	sleep(5);

	/*
		删除共享内存
		shmid:是shmget函数返回的共享内存标识符
		command:填IPC_RMID
		buf:填0;
	*/
	if(shmctl(shmid,IPC_RMID,0)==-1){
		printf("shmctl(0x5005) failed\n");return -1;
	}

	printf("ok\n");

	return 0;
}

int main(){

	mem();
	

}

