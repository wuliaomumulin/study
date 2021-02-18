#include <stdio.h>
#include <string.h>

/*
	C++的STL(Standard Template library,标准模板类)中定义的类，他会存放字符的长度自动伸缩，程序员不必担心内存溢出的问题，string类还和C语言的字符串之间可以很方便的切换
	必须包含 #include <string>
	string是一个模板类，位于std命名空间内，为方便使用还需在程序中增加
	using namespace std;//指定省缺的命名空间
	string str;//创建string对象
	如果不指定命名空间，也就是说没有using namespace std,创建string对象的方法如下
	std::string str;

*/
#include <string>

using namespace std;

int main(int argc,char *argv[]){
	char str1[31];//C语言风格写法
	memset(str1,0,sizeof(str1));
	strcpy(str1,"张三");

	string str2;//C++的string字符串
	str2="沉鱼";
	printf("str2=%s\n",str2.c_str());

	//判断str1和str2是否相等
	if(str2!=str1) printf("str2=%s,str1=%s,不相等\n",str2.c_str(),str1);

	strcpy(str1,str2.c_str());//把string的内容赋值给C语言风格的字符串
	printf("str1=%s\n",str1);

	//string是一个类，通过动态分配内存，实现对字符串的存储，我们来看以下代码
	std::string str3;
	str3="大家好";
	printf("%p=>%s\n",str3.c_str(),str3.c_str());//打印指针
	str3+=",我是张三";
	printf("%p=>%s\n",str3.c_str(),str3.c_str());//打印指针
	str3="我是李四";
	printf("%p=>%s\n",str3.c_str(),str3.c_str());//打印指针
	str3="www.baidu.com";
	printf("%p=>%s\n",str3.c_str(),str3.c_str());//打印指针
	str3="www.baidu.com";
	printf("%p=>%s\n",str3.c_str(),str3.c_str());//打印指针

}
