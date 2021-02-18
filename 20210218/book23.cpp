#include <stdio.h>
#include <string.h>

/*
	与string一样，向量vector同属于STL(Standard Template Library,标准模板库)中的定义的类，vector是一个封装了动态数组的顺序容器(Sequence Container),它能够存放各种类型的数据和对象
	
	首先，如果要在程序中使用容器，必须包含头文件<vector>,如下
	#include <vector>
	vector类是一个模板类，位于std命名空间下，为方便使用还需增加:
	using namespace std;
	声明一个容器很简单
	vector<int> vi;//定义存放整数的容器
	vector<double> vd;//定义存放浮点数容器
	vector<string> vs;//定义用于存放string字符串的容器
	vector<struct str_girl> vgirl; //定义用于存放超女结构体的容器
	vector<CGirl> vGirl;//定义用于存放超女类的容器


*/
#include <string>
#include <vector>
//sort函数需要的头文件
#include <algorithm>

//存放结构体
struct st_girl{
	char name[50];
	int age;
};

//容器的排序
//自定义排序函数，按姓名排序
bool sortbyname(const st_girl &p1,const st_girl &p2){
	if(strcmp(p1.name,p2.name)<=0) return true;
	return false;
}
//自定义排序，按照年龄排序
bool sortbyage(const st_girl &p1,const st_girl &p2){
	if(p1.age<p2.age) return true;
	return false;
}

int main(int argc,char *argv[]){
	
	int height=0;//存放从键盘输入的超女身高
	char name_tmp[50];//存放姓名

	std::vector<int> vheight;
	std::vector<std::string> vname;//存放超女姓名的容器

	while(true){
		printf("请输入\n");

		printf("姓名:");
		scanf("%s",name_tmp);
		if(strcmp(name_tmp,"0")==0) break;
		vname.push_back(name_tmp);//把数据加入容器

		printf("身高:");
		scanf("%d",&height);
		if(height==0) break;
		vheight.push_back(height);//把数据加入容器

		printf("\n");

	}

	for(int i=0;i<vheight.size();i++){
		//显示排序前容器的记录
		printf("vname[%d]=%s,vheight[%d]=%d\n",i,vname[i].c_str(),i,vheight[i]);
	}

	/*
		数组的起始地址，结束地址、排序方式，如果不是基本数据类型，则必须指定排序方法函数
	*/
	sort(vheight.begin(),vheight.end());//容器中的记录排序
	printf("排序完成\n");

	for(int i=0;i<vheight.size();i++){
		//显示排序前容器的记录

		printf("vname[%d]=%s,vheight[%d]=%d\n",i,vname[i].c_str(),i,vheight[i]);
	}

	vheight.clear();//清空容器，可以不写

	printf("-----------------存放结构体-------------\n");
	struct st_girl stgirl;//数据结构
	std::vector<struct st_girl> vgirl;//存放结构体的容器
	strcpy(stgirl.name,"张三");stgirl.age=25;
	vgirl.push_back(stgirl);
	strcpy(stgirl.name,"李四");stgirl.age=28;
	vgirl.push_back(stgirl);
	strcpy(stgirl.name,"王五");stgirl.age=22;
	vgirl.push_back(stgirl);

	//采用数组下标访问容器中的元素
	for (int i = 0; i < vgirl.size(); ++i)
	{
		printf("vgirl[%d].name=%s,vgirl[%d].age=%d\n",\
			i,vgirl[i].name,i,vgirl[i].age);
	}

	//把容器中的记录复制到结构体中
	for (int i = 0; i < vgirl.size(); ++i)
	{
		memcpy(&stgirl,&vgirl[i],sizeof(struct st_girl));
		printf("stgirl.name=%s,stgirl.age=%d\n",stgirl.name,stgirl.age);
	}

	printf("-----------------容器的排序-----------------\n");
	printf("-----------------按照名称排序-----------------\n");

	sort(vgirl.begin(),vgirl.end(),sortbyname);
	for (int i = 0; i < vgirl.size(); ++i)
	{
		printf("vgirl[%d].name=%s,vgirl[%d].age=%d\n",\
			i,vgirl[i].name,i,vgirl[i].age);
	}
	printf("-----------------按照年龄排序-----------------\n");
	sort(vgirl.begin(),vgirl.end(),sortbyage);
	for (int i = 0; i < vgirl.size(); ++i)
	{
		printf("vgirl[%d].name=%s,vgirl[%d].age=%d\n",\
			i,vgirl[i].name,i,vgirl[i].age);
	}

}
