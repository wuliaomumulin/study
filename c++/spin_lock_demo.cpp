#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>
#include <vector>
/*
 一、线程同步
 2、自旋锁
 a、自旋锁避免了进程或线程上下文切换的开销；
 b、操作系统有很多地方使用的是自旋锁；
 c、自旋锁不适合在单核CPU使用；
 */
//定义一个自旋锁
pthread_spinlock_t spin_lock;

int num = 0;

//生产者
void *producer(void*){
	int timer = 100000000;
	while(timer--){

		//加锁
		pthread_spin_lock(&spin_lock);

		num+=1;

		//减锁
		pthread_spin_unlock(&spin_lock);
	}
}

//消费者
void *comsumer(void*){
    int timer = 100000000;
	while(timer--){

		//加锁
		pthread_spin_lock(&spin_lock);

		num-=1;

		//减锁
		pthread_spin_unlock(&spin_lock);

	}	
}

int main(){
	printf("Start up main function.\n");
	//初始化一个自旋锁
	pthread_spin_init(&spin_lock,0);

	pthread_t thread1,thread2;
	pthread_create(&thread1,NULL,&producer,NULL);
	pthread_create(&thread2,NULL,&comsumer,NULL);
	pthread_join(thread1,NULL);
	pthread_join(thread2,NULL);
	printf("Print in main function: num = %d\n",num);
	return 0;
}
//g++编译
//g++ spin_lock_demo.cpp -o spin_lock_demo -g -lpthread
//-o 确定文件名 -g gdb调试 -l动态链接库