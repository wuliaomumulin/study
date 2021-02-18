#include "_ku.h"


/*
	在centos上使用g++来编译


文件操作了类声明

 */
class CFile{
private:
	FILE *m_fp; //文件指针
	bool m_bEnBuffer; //是否启用缓存区,true-启用,false-禁用
public:
	CFile(); //类的构造函数
	CFile(bool bEnBuffer);//类的构造函数

	~CFile();//类的析构函数
	void EnBuffer(bool bEnBuffer=true);//启用、禁用缓存区

	//打开文件、参数与fopen相同，打开成功true,失败返回false
	bool Open(const char *filename,const char *openmode);

	//调用fprintf向文件写入数据
	void Fprintf(const char *fmt,...);

	//调用fgets从文件中读取一行
	bool Fgets(char *strBuffer,const int ReadSize);

	//关闭文件指针
	void Close();
};

int main(int argc,char *argv[]){
	if(argc != 2){printf("请输入带打开的文件名\n");return -1;}

	CFile File;

	if(File.Open(argv[1],"r")==false){printf("文件打开失败\n");return -1;}

	char strLine[301];

	while(true){
		//从文件中读取一行
		if(File.Fgets(strLine,300)==false) break;

		printf("%s\n",strLine);//把从文件中读取到的内容输出到屏幕
	}
}


//定义类的方法

//够的构造函数
CFile::CFile(){
	m_fp=0;
	m_bEnBuffer=true;
}
CFile::CFile(bool bEnBuffer){
	m_fp=0;
	m_bEnBuffer=bEnBuffer;
}
//关闭文件指针
void CFile::Close(){
	if(m_fp!=0) fclose(m_fp);//关闭文件指针
	m_fp=0;
}
//类的析构函数
CFile::~CFile(){
	Close();//调用Close释放资源
}
//启用、禁用缓存区
void CFile::EnBuffer(bool bEnBuffer){
	m_bEnBuffer=bEnBuffer;
}
//打开文件、参数与fopen相同，打开成功true,失败返回false
bool CFile::Open(const char *filename,const char *openmode){
	Close(); //打开新的文件之前、如果已经打开了文件，关闭它

	if((m_fp=fopen(filename,openmode))==0) return false;

	return true;
}
//调用fprintf向文件写入数据
void CFile::Fprintf(const char *fmt,...){
	if(m_fp==0) return;

	va_list arg;
	va_start(arg,fmt);
	vfprintf(m_fp,fmt,arg);
	va_end(arg);

	if(m_bEnBuffer==false) fflush(m_fp);
}

//调用fgets从文件中读取一行
bool CFile::Fgets(char *strBuffer,const int ReadSize){
	if(m_fp==0) return false;
	memset(strBuffer,0,ReadSize);

	if(fgets(strBuffer,ReadSize,m_fp)==0) return false;

	return true;
}