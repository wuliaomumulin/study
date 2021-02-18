#include <stdio.h>
#include <string.h>

class CGirl{
public:
	char m_name[50]; //姓名
	char m_sc[30]; //身材

	bool operator==(const CGirl &Girl); //重载==运算符
};

int main(int argc,char *argv[]){

	CGirl Girl1,Girl2;
	strcpy(Girl1.m_name,"张三");
	strcpy(Girl2.m_name,"李四");
	if(Girl1 == Girl2) printf("Girl1和Girl2是同一个人\n");
	else printf("Girl1和Girl2不是一个人\n");

}
bool CGirl::operator==(const CGirl &Girl){
	//如果两个类的名称相同就相等
	if(strcmp(m_name,Girl.m_name)== 0) return true;
	
	return false;
}
