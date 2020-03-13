#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>
#include <vector>
/*
 一、线程同步
 2、读写锁
 	a、临界资源多读少写；
 	b、读取的时候并不会改变临界资源的值；
 	c、是否存在效率更高的同步方法？
 3、读写锁是一种特殊的自旋锁
 	a、允许多个读者同时访问资源以提高性能；
 	b、对于写操作是互斥的;
 */

int num = 0;
//定义一个读写锁
pthread_rwlock_t rwlock = PTHREAD_RWLOCK_INITIALIZER;
//pthread_mutex_t mutes = PTHREAD_MUTEX_INITIALIZER;


//读者
void *reader(void*){
	int timer = 10000000;
	while(timer--){
		//加读锁
		pthread_rwlock_rdlock(&rwlock);
		//pthread_mutex_lock(&mutes);

		if(timer % 1000 == 0){
			//printf("print num in reader:num=%d\n", num);
			//1s
			//sleep(1);
			//10ms
			usleep(10);
		}

		//减锁
		pthread_rwlock_unlock(&rwlock);
		//pthread_mutex_unlock(&mutes);
		
	} 

}

//写者
void *writer(void*){
    int timer = 10000000;
	while(timer--){
		//加写锁
		pthread_rwlock_wrlock(&rwlock);
		//pthread_mutex_lock(&mutes);

		num+=1;

		//减锁
		pthread_rwlock_unlock(&rwlock);//4.248s
		//pthread_mutex_unlock(&mutes);//10.961s
		
	}	
}

int main(){
	printf("Start up main function.\n");
	pthread_t thread1,thread2,thread3;
	pthread_create(&thread1,NULL,&reader,NULL);
	pthread_create(&thread2,NULL,&reader,NULL);
	pthread_create(&thread3,NULL,&writer,NULL);
	pthread_join(thread1,NULL);
	pthread_join(thread2,NULL);
	pthread_join(thread3,NULL);
	printf("Print in main function: num = %d\n",num);
	return 0;
}
//g++编译
//g++ rwlock_demo.cpp -o rwlock_demo -g -lpthread
//-o 确定文件名 -g gdb调试 -l动态链接库
//time rwlock_demo 运行耗时