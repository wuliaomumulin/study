#include "_ku.h"


/*
	为了习惯使用C进行开发的程序员来说，&符号是取地址符，但是在c++中，他除了取地址符，还有其他的用途，叫做引用(reference)，引用是c++的新特性
	引用的申明方法 
	数据类型 &引用名=目标变量名;
	int ii;
	int &rii=ii;//定义引用rii，他是变量ii的引用，即别名
	rii=1;等价于ii=1;

 */
//用于函数的参数
void func(int *a){
	//a是一个指针
	*a=20;
}
void func(int &ra){
	//ra是一个引用
	ra=30;
}
//用户函数的返回值
//
int val; //定义一个全部变量val;
int func1(int ii);//函数的返回值是int
int &func2(int ii);//函数的返回值是int的引用

int main(int argc,char *argv[]){
	//用于函数的参数
	printf("用于函数的参数\n");
	int ii=0;
	func(&ii);//传递变量的地址
	printf("1 = %d\n",ii);

	func(ii);//引用
	printf("2 = %d\n",ii);

	//用户函数的返回值
	printf("用户函数的返回值\n");
	int aa=func1(10);
	printf("aa=%d\n",aa);//20

	//int &bb=func1(10);
	//printf("bb=%d\n",bb);//类型为‘int’的右值初始化类型为‘int&’的非常量引用无效

	int cc=func2(20);
	printf("cc=%d\n",cc);//40
	int &dd=func2(20);
	printf("dd=%d\n",dd);//40

}
//函数的返回值是int
int func1(int ii){
	val=10+ii;
	return val;
}
//函数的返回值是int的引用
int &func2(int ii){
	val=20+ii;
	return val;
}
