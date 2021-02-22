#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <unistd.h>
#include <stdlib.h>
/*
调用可执行程序
	1、execl函数族
	2、system函数
	 int system(const char * string);
 	 底层是:
 	 execl("/bin/sh", "sh", "-c", command, (char *) 0);

*/


void Execl();
void System();

int main(){

	System();


	Execl();


}

void Execl(){ 
	/**

	int execl(const char *path, const char *arg, ...
                        (char  *) NULL *);
       int execlp(const char *file, const char *arg, ...
                        (char  *) NULL );
       int execle(const char *path, const char *arg, ...
                       , (char *) NULL, char * const envp[] );
       int execv(const char *path, char *const argv[]);
       int execvp(const char *file, char *const argv[]);
       int execvpe(const char *file, char *const argv[],
                       char *const envp[]);


	* path 要执行的程序路劲，可以是绝对路径和相对路劲
	* file 要执行的程序名称，如果该参数中包含'/'字符，则视为路径名直接执行，否则是为单独的文件名，系统将根据PATH环境变量指定的路径顺序搜索指定的文件
	* argv 命令行参数的数组
	* envp 带有该参数的exec函数可以在调用时指定一个环境变量数组，其他不带该参数的exe函数则使用调用者进程的环境变量
	* arg 程序的第0个参数，即程序名本身
	* ... 命令行参数列表，调用相应程序时有多少个命令行参数，就需要有多少个输入参数项。注意：在使用此类函数时，在所有命令行参数的最后一个参数增加一个空的参数项，表明命令行参数结束

	执行成功不返回，执行失败返回-1，失败原因存在errno中	
	*/
	printf("Execl\n");


	int iret=execl("/bin/lss","/bin/ls","-1","/usr/include/stdio.h",0);// /bin/lss不存在，执行不成功
	int iret1=execl("/bin/ls","/bin/ls","-1","/usr/include/stdio.h",0);// 执行成功的代码

	printf("iret=%d,iret1=%d\n",iret,iret1);
	if(iret==-1) printf("%d:%s\n",errno,strerror(errno));

}
void System(){
	
	int iret;
	printf("System\n");

	//调用不成功的代码
	iret=system("/bin/lss -l /usr/include/stdio.h");
	printf("iret=%d\n",iret);
	if(iret==-1) printf("%d:%s\n",errno,strerror(errno));

	//调用成功的代码
	iret=system("/bin/ls -l /usr/include/stdio.h");
	printf("iret=%d\n",iret);
	if(iret==-1) printf("%d:%s\n",errno,strerror(errno));
}