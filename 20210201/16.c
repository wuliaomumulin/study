#include "_ku.h"


/*
	编译预处理
	C源程序->编译预处理->编译->优化程序->汇编程序->链接程序->可执行文件

	预处理指令分三种
	1、include
	2、define
	3、if

*编译预处理不包括错误检测，但是定义宏的时候还是不要带;或其他的错误字符
*宏的定义会在编译预处理阶段对常量名进行替换，但不会替换表达式里面的内容
*

 */
//定义一个宏
#define PI 3.1415926
//定义一个常量，常量在编译预处理阶段不进行常量名的替换
const double PII=3.1415926;
//带参数的宏，宏名和参数之间不能有空格
#define MAX(x,y) ((x)>(y)?(x):(y))
//条件预处理
#define LINUX

int main(int argc,char *argv[]){
 //gcc -E -o 11 11.c
 //只预处理，不编译
 printf("你好%lf,max(20,30)=%d,const PII=%lf\n",PI,max(20,30),PII);
 printf("带参数的宏比大小，MAX(14,21)=%d\n",MAX(14,21));

 //条件预处理宏
 #ifdef LINUX
  printf("linux系统\n");
 #else
  printf("不是linux系统\n");
 #endif

//取消已定义的宏
#undef LINUX

 #ifndef LINUX
  printf("不是linux系统\n");
 #else
  printf("linux系统\n");
 #endif

 return 0;
}