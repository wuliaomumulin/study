#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/sem.h>
/*
一、信号量
	信号量(信号灯)本质上是一个计数器，用于协调多个进程(包括但不限于父子进程)对共享数据对象的都和写，它不以传送数据为目的，主要是用来保护共享资源(共享内存、消息队列、socket连接池、数据库连接池)，保证共享资源在一个时刻只有一个进程独享

	信号量是一个特殊的变量，只允许进程对它进行等待信号和发送信号操作，最简单的信号量是取值0和1的二元信号量，这是信号量最常见的形式

	通常信号量(可以取值多个正整数值)和信号量集方面的只是比较复杂，遂只介绍二元信号量	

	

*/

class csem{
private:
	union semun //用于信号灯操作的共同体
	{
		int val;
		struct semid_ds *buf;
		unsigned short *arry;
	};

	int sem_id;//信号灯描述符

public:
	//如果信号灯已存在，获取信号灯，如果不存在，则创建信号灯并初始化
	bool init(key_t key);
	//等待信号灯挂出
	bool wait();
	//挂出信号灯
	bool post();
	//销毁信号灯
	bool destroy();

};


int main(int argc,char *argv[]){

	csem sem;

	//初始化信号灯
	if(sem.init(0x5000)==false){
		perror("sem.init failed\n");return -1;
	}
	printf("sem init ok\n");
	
	//等待信号灯挂出，等待成功后，将持有锁
	if(sem.wait()==false){
		perror("sem.wait failed\n");return -1;
	}
	printf("sem wait ok\n");

	sleep(20);//在沉睡的过程中，再次运行当前程序

	//挂出信号灯，释放锁
	if(sem.post()==false){
		perror("sem.post failed\n");return -1;
	}
	printf("sem post ok\n");

	//销毁信号灯
/*	if(sem.destroy()==false){
		perror("sem.destroy failed\n");return -1;
	}
	printf("sem destroy ok\n");*/
}

bool csem::init(key_t key){
	//获取信号灯
	if((sem_id=semget(key,1,0640))==-1){

		//如果信号灯不存在，创建它
		if(errno==2){
			if((sem_id=semget(key,1,0640|IPC_CREAT))==-1){
				perror("init 1 semget()");return false;
			}

			//信号灯初始化成功后，还需要把他初始化为可用状态
			union semun sem_union;
			sem_union.val = 1;
			if(semctl(sem_id,0,SETVAL,sem_union)<0){
				perror("init semctl()");return false;
			} 

		}else{
			perror("init 2 semget()");return false;
		}

	}

	return true;
}

bool csem::destroy(){
	if(semctl(sem_id,0,IPC_RMID)==-1){
		perror("destroy semctl()");return false;
	}
	return true;
}

bool csem::wait(){

	struct sembuf sem_b;
	sem_b.sem_num = 0;
	sem_b.sem_op = -1;
	sem_b.sem_flg = SEM_UNDO;
	if(semop(sem_id,&sem_b,1)==-1){
		perror("wait semop()");return false;
	}

	return true;
}

//挂出信号灯
bool csem::post(){
	struct sembuf sem_b;
	sem_b.sem_num = 0;
	sem_b.sem_op = 1;
	sem_b.sem_flg = SEM_UNDO;
	if(semop(sem_id,&sem_b,1)==-1){
		perror("post semop()");return false;
	}

	return true;
}
/**
信号量查看
ipcs -s
key   semid owner   perms   nsems
键值   编号  创建者   权限  信号量数

信号量删除
ipcrm sem 编号

*/