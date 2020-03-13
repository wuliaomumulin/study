#include <iostream>
#include <cstring>
#include <stdio.h>
#include <unistd.h>

using namespace std;

int main(){
	int num = 888;
	pid_t pid;
	pid = fork();
	//一次执行，fork会返回两次
	if(pid == 0){
		//表示输出一个回车，并刷新输出流，和printf("\n")显示一样
		cout << "这是一个子进程." << endl;
		cout << "num is son process:" <<  num << endl;
		while(true){
			num += 1;
			cout << "num is son process:" << num << endl;
			sleep(1);
		}
	}else if(pid > 0){
		cout << "这是一个父进程." << endl;		
		cout << "子进程id:" << pid << endl;
		cout << "num is father process:" <<  num << endl;
		while(true){
			num -= 1;
			cout << "num is father process:" << num << endl;
			sleep(1);
		}
	}else if(pid < 0){
		//进程没有创建成功
		cout << "创建进程失败." << endl; 
	}
	return 0;
}
// g++ fork_demo.cpp -o fork_demo -g