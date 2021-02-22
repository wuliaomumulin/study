#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>



/**
	进程的概念
	进程就是内存中运行的程序，linux下一个进程在内存中有代码段、堆栈段、数据段三部分组成。
	代码段：存放代码程序
	堆栈段：程序的返回地址、程序的参数以及程序的局部变量
	数据段：存放数据的全局变量、常数以及动态数据分配的数据空间(比如new函数分配的空间)

	系统如果同时运行多个相同的程序，他们的代码是相同的，堆栈段和数据段是不同的(相同的程序，处理的程序不同)

	一、进程的编号
	ps -ef|more //分页显示系统全部进程
	ps -ef|grep book|grep -v grep//查看是否启动

	UID:启动进程的操作系统用户
	PID:进程编号
	PPID:进程的父进程编号
	C:CPU使用的资源百分比
	STIME:进程的启动时间
	TTY:进程所属的终端
	TIME:使用掉的CPU时间
	CMD:执行的是什么命令

	可以用 pid_t getpid();//获取程序运行的系统编号

*/

void fatchfunc(){
	printf("我是爸爸\n");
}
void childfunc(){
	printf("我是儿子\n");
}


//父子进程是相互隔离的
int i=10;
void separate(){
	int j=20;
	if(fork()>0){
		i=11;j=21;sleep(1);printf("父进程:i=%d,j=%d\n",i,j);
	}else{
		i=12;j=22;sleep(1);printf("子进程:i=%d,j=%d\n",i,j);
	}
}

int main(int argc,char *argv[]){
	
	printf("本程序的进程编号:%d\n",getpid());

	int ipid=fork();

	sleep(1);

	printf("pid=%d\n",ipid);

	if(ipid!=0) printf("父进程的编号:%d\n",getpid());
	else printf("子进程的编号:%d\n",getpid());

	sleep(5);

	//第二段
	printf("------------第二段-------------\n");
	if(fork()>0){printf("父进程\n");fatchfunc();}
	else{printf("子进程\n");childfunc();}

	sleep(1);
	printf("父子进程执行完自己的函数后都护来到这里\n");
	sleep(1);

	printf("------------第三段-------------\n");
	separate();

}

//g++ book31.c -o book31