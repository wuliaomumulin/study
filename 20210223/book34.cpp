#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
/*
	一、进程通信
    	1、管道:包括无名管道(pipe)和命名管道(named pipe),无名管道用于具有父进程和子进程之间通信，命名管道克服了管道没有名字的限制，因此，除具有管道所具有的功能外，它还允许无亲缘关系进程间的通信;
    	2、消息队列(message):进程可以向队列中添加消息，其他的进程可以读取队列里的消息;
    	3、信号(signal):信号用于通知其他进程有某种事件发生;
    	4、共享内存(shared memory):多个进程可以访问同一片内存空间;
    	5、信号量(semaphore):也叫信号灯，用于进程之间对共享资源进行加锁;
    	6、套接字(socket):可用于不同计算机之间的进程间的通信;

	二、如何让程序在后台运行
		1、加&符号;
		2、采用fork一个子进程，同时父进程
			if(fork()>0) return 0;
	
	三、终止后台运行中的程序
		1、killall book34 
		2、kill -9 pid


	四、signal信号
		详情搜索linux信号列表
	五、发送信号
		函数声明:kill(pid_t pid,int sig);
			参数pid
				1、pid>0将信号传递给进程号为pid的进程
				2、pid=0将信号传给和目前进程相同进程组的所有进程，常用于父进程给子进程传递信号，注意，发送信号者进程也会收到自己发出的信号;
				3、pid=-1将信号广播传递给系统内所有的进程，例如系统关机时，会向所有的登录窗口广播关机信息;
			参数sig
				准备发送的信号代码，假如其值为零则没有任何信号发出，但是系统会执行系统检查，通常会利用sig值为零来检验某个进程是否仍在运行;
			返回值说明：成功返回0，失败返回-1并设置error值为以下某个
				EINVAL:指定的信号码无效(参数sig不合法);
				EPERM:权限不够无法传递信号给指定进程;
				ESRCH:参数pid所指定的进程或进程组不存在;


*/
void Exit(int sign){

	printf("收到了信号%d,程序退出。\n",sign);

	//在这里田间释放资源的代码
	exit(0); //程序退出

}

int main(){

	for (int i = 0; i < 100; ++i) signal(i,SIG_IGN);//屏蔽全部的信号

	/*设置SIGINT和SIGTERM的处理函数为EXIT退出;
		1、ctrl+c
		2、killall book34
	*/
	signal(SIGINT,Exit);signal(SIGTERM,Exit);

	//一个死循环
	while(1){
	
		printf("book34\n");
		sleep(1);

	}

}

