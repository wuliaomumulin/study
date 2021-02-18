#include <stdio.h>

/*
	从c到c++
    本人是在ubuntu下模拟的数据
    g++ book18.c -o book18
*/
int main(void){
 /*
	C语言没有彻底从语法上支持真或假，只是用0或非0来代表，这点在C++中得到了改善，c++新增了bool类型，它占用一个字节长度，bool类型只有两个取值，true和flase：true代表真，false代表假
 */
 bool flag=true;//定义bool类型，赋值为true

 if(flag) printf("flag is true\n");
 else printf("flag is false\n");

 flag=false;
 if(flag) printf("flag is true\n");
 else printf("flag is false\n");

 /*
   c++变量定义的位置
 */
 int num=0;
 for (int i = 0; i < 100; ++i)
 {
 	num = num + i;
 }
 printf("你好,总数为%d\n",num);


 return 0;
}


