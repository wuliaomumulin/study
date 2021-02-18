#include <stdio.h>
#include <string.h>

/*
	函数重载
    
*/

//重载strcpy函数，对dest初始化，防止内存溢出
extern char *strcpy(char *dest, const char *src, size_t n);

int main(){

	char dest[10];

	strcpy(dest,"www.baidu.com",sizeof(dest));//新strcpy
	printf("%s\n",dest);

	strcpy(dest,"www.baidu.com");
	printf("%s\n",dest);



}

//重载strcpy函数，对dest初始化，防止内存溢出
char *strcpy(char *dest, const char *src, size_t n){
	memset(dest,0,n);//对dest初始化

	// 计算需要复制的字符数，不能超过dest-1
	size_t len=0;
	if(strlen(src)<=n-1) len=strlen(src);
	else len=n-1;

	strncpy(dest,src,len);//复制字符串，超出的将丢掉
	dest[len]=0;

	return dest;

}
