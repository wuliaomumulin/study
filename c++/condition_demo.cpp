#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>
#include <vector>
#include <queue>

/*
 一、线程同步
 2、条件变量
 	a、条件变量是一种相对复杂的线程同步方法；
 	b、条件变量允许线程睡眠，知道满足某种条件；
 	c、当满足条件时，可以向该线程信号，通知唤醒;
 3、生产者和消费者模型：
 	a、缓存区小于等于0时，不允许消费者消费，消费者必须等待；
 	b、缓存区大于等于峰值时，不允消费者往缓存区生产，生产者必须等待;
 */
//设置缓存区峰值
int MAX_BUF = 100;

int num = 0;
//定义一个条件变量
pthread_cond_t cond = PTHREAD_COND_INITIALIZER;
//配合互斥量来使用
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

//生产者
void *producer(void*){
	int timer = 10000000;
	while(timer--){
		pthread_mutex_lock(&mutex);
			while(num >= MAX_BUF){
				//等待
				pthread_cond_wait(&cond,&mutex);
				printf("缓存区已满，等待消费者进行消费.\n");
			}
		num += 1;
		
		printf("生产一个产品，当前产品数量:%d.\n",num);

		pthread_cond_signal(&cond);
		printf("通知可能等待的消费者.\n");
		pthread_mutex_unlock(&mutex);
	
		//为了效果更直接，我们牺牲一个效率
		sleep(1);
	}
} 

 
//消费者
void *consumer(void*){
    int timer = 10000000;
	while(timer--){
		pthread_mutex_lock(&mutex);
		while(num <= 0){
			pthread_cond_wait(&cond,&mutex);
			printf("缓存区已空，等待生产者进行生产.\n");
		}
		num-=1;

		printf("消费一个产品，当前产品数量:%d.\n",num);

		pthread_cond_signal(&cond);
		printf("通知可能等待的生产者.\n");
		pthread_mutex_unlock(&mutex);	
	}	
}

int main(){
	printf("Start up main function.\n");
	pthread_t thread1,thread2;
	pthread_create(&thread1,NULL,&consumer,NULL);
	pthread_create(&thread2,NULL,&producer,NULL);
	pthread_join(thread1,NULL);
	pthread_join(thread2,NULL);
	printf("Print in main function: num = %d\n",num);
	return 0;
}
//g++编译
//g++ condition_demo.cpp -o condition_demo -g -lpthread
//-o 确定文件名 -g gdb调试 -l动态链接库
//time condition_demo 运行耗时