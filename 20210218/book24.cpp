#include <stdio.h>
#include <string.h>

/*
	C++动态内存
	在C++中，内存需求都是在编写程序的时候声明变量来确定，但是程序在运行过程中需要动态分配内存的情况，c++语言将运算符new和delete合并在一起

	datatype *pointer = new datatype
	delete pointer;
	datatype可以是基本数据类型，也可以是结构体，还可以是类

	int *pi=new int;//动态分配一个整数大小的内存
	(*pi)=10;
	delete pi;//释放pi指向的内存


*/
#include <string>
#include <vector>



class CGirl{
public:
	char name[50];
	int age;
	int height;
	char sc[30];
	char yz[30];
	int show();
};


int main(int argc,char *argv[]){
	int *pi=new int;//动态分配一个整数大小的内存
	(*pi)=10;
	printf("pi=%d\n",*pi);
	delete pi;//释放pi指向的内存	
	
	CGirl *pgirl = new CGirl;//动态分配CGirl类
	strcpy(pgirl->name,"西施");
	pgirl->age=25;
	pgirl->height=168;
	strcpy(pgirl->sc,"良好");
	strcpy(pgirl->yz,"漂亮");
	pgirl->show();

	delete pgirl;//释放pgirl指向的内存

}
//函数成员函数的定义
int CGirl::show(){
	printf("name=%s,age=%d,height=%d,sc=%s,yz=%s\n",name,age,height,sc,yz);
}
//g++ book24.cpp -o book24