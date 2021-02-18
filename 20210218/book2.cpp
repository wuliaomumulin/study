#include <stdio.h>
#include <string.h>

/*
	C++结构体中可以有函数，strcut可以不书写,而C中不可以
    
*/

struct st_girl{
	char name[50];
	int age;
	int height;
	char sc[30];
	char yz[30];
	int show();
};
/**
	C++把结构体看成了类(class),类的成员可以是变量和函数，通过类定义的变量也有特定的称呼，叫做对象
	
*/
class CGirl{
public:
	char name[50];
	int age;
	int height;
	char sc[30];
	char yz[30];
	int show();
};

int main(){

	st_girl stgirl;//struct关键字可以不书写
	memset(&stgirl,0,sizeof(stgirl));
	strcpy(stgirl.name,"张三");
	stgirl.age=28;
	stgirl.height=187;
	strcpy(stgirl.sc,"火辣");
	strcpy(stgirl.yz,"漂亮");
	stgirl.show();//调用结构体的成员函数

	CGirl stgirl1;//实例化一个girl对象

	strcpy(stgirl1.name,"李四");
	stgirl1.age=26;
	stgirl1.height=164;
	strcpy(stgirl1.sc,"普通");
	strcpy(stgirl1.yz,"一般");
	stgirl1.show();//调用结构体的成员函数


}
//结构体成员函数的定义
int st_girl::show(){
	printf("name=%s,age=%d,height=%d,sc=%s,yz=%s\n",name,age,height,sc,yz);
}
//函数成员函数的定义
int CGirl::show(){
	printf("name=%s,age=%d,height=%d,sc=%s,yz=%s\n",name,age,height,sc,yz);
}