#include <stdio.h>
#include <string.h>

/*
	继承

  class <派生类名>:<继承方式[private|protected|public]> <基类名1>,<继承方式[private|protected|public]> <基类名2>

*/




class CGirl{
public:
	char name[50];
	int age;
	int height;
	char sc[30];
	char yz[30];
	int show(){
		printf("name=%s,age=%d,height=%d,sc=%s,yz=%s\n",name,age,height,sc,yz);
	}
	virtual int virtualshow(){
		printf("name=%s,age=%d,height=%d,sc=%s,yz=%s\n",name,age,height,sc,yz);
	}
	virtual int virtualshow1()=0;//在基类中声明函数而不定义叫做纯虚函数
};

class CKCon:public CGirl{ //继承该类
public:
	char m_ch[50];//称号
	char m_palace[50];//居住的宫殿
	int m_sal;//俸禄
	int show(){
		printf("姓名:%s,称号:%s,居住的宫殿:%s,俸禄:%d两银子\n",name,m_ch,m_palace,m_sal);
	}
	int virtualshow(){
		printf("姓名:%s,称号:%s,居住的宫殿:%s,俸禄:%d两银子\n",name,m_ch,m_palace,m_sal);
	}
	/*
	1、接口描述了类的行为和功能，是标准和规范，而不需要完成类的功能实现;
	2、C++接口是用抽象类来实现的，如果类中至少有一个函数被声明为纯虚函数，则这个类就是抽象类；
	3、设计抽象类的目的，是为了给其他类提供一个可以继承的基类。它只能作为接口使用。如果试图实例化一个抽象类的对象，会导致编译错误;
	4、如果一个基类的派生类需要实例化，则必须实现每个虚函数的定义，如果没有在派生类中有纯虚函数的定义会导致编译错误。
	5、可用于实例化对象的类被称为具体类
	*/
	virtual int virtualshow1(){
		printf("姓名:%s,称号:%s,居住的宫殿:%s,俸禄:%d两银子\n",name,m_ch,m_palace,m_sal);
	}
};


int main(int argc,char *argv[]){
	
	/*
		静态多态
	*/
	printf("-----------------静态多态-------------\n");
	CKCon KCon;//实例化一个KCon对象
	strcpy(KCon.name,"张三");
	KCon.age=31;
	KCon.height=178;
	strcpy(KCon.sc,"一般");
	strcpy(KCon.yz,"拉跨");
	strcpy(KCon.m_ch,"法外狂徒");
	strcpy(KCon.m_palace,"法院");
	KCon.m_sal=10000;

	CGirl *pGirl;//基类的指针
	CKCon *pCon;//派生类的指针
	
	pGirl=pCon=&KCon;//基类的指针和派生类的指针都指向派生类

	pGirl->show();//将调用基类的方法
	pCon->show();//将调用父类的方法

	/**
		动态多态
		是在程序运行时根据基类的引用(指针)指向的对象来确定自己具体该调用哪一类的虚函数
		让我们对程序稍作修改，在CGirl类中，Show方法的声明前放置关键字virtual,如下所示:
		virtual int show();

	*/
	printf("-----------------动态多态-------------\n");
	pGirl->virtualshow();//将调用父类的方法
	pCon->virtualshow();//将调用父类的方法

}
//g++ book26.cpp -o book26