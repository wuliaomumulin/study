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
};

class CKCon:public CGirl{ //继承该类
public:
	char m_ch[50];//称号
	char m_palace[50];//居住的宫殿
	int m_sal;//俸禄
	int show(){
		printf("姓名:%s,称号:%s,居住的宫殿:%s,俸禄:%d两银子\n",name,m_ch,m_palace,m_sal);
	}
};


int main(int argc,char *argv[]){
	
	CKCon KCon;//实例化一个KCon对象
	strcpy(KCon.name,"张三");
	KCon.age=31;
	KCon.height=178;
	strcpy(KCon.sc,"一般");
	strcpy(KCon.yz,"拉跨");
	strcpy(KCon.m_ch,"法外狂徒");
	strcpy(KCon.m_palace,"法院");
	KCon.m_sal=10000;

	KCon.show();//如果子类没有该方法，会调用父类的该方法
}
//g++ book25.cpp -o book25