#include "_ku.h"


/*
	获取系统错误

 */
int main(int argc,char *argv[]){
 //打印全部错误码和内容
 int i;
 for (i = 0; i < 132; ++i)
 {
 	printf("%d:%s\n",i,strerror(i));
 }

 //打开文件
 FILE *fp;
 if((fp=fopen("/tmp/2021-01-27.txt","r")) == 0){
 	printf("%d:%s\n",errno,strerror(errno));
 	printf("文件打开失败\n");
 	perror("错误前缀");
 	return -1;
 }

 return 0;
}