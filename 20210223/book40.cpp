#include "book40public.h"
/*
linux静态库和动态库
	1、静态库：我们通常把公用的自定义函数和类从主程序中分离出来，函数和类的申明在头文件中，定义在程序文件中，主城需要包含头文件，编译时要和程序文件一起编译.
		编译成静态库指令
		1)、g++ -c -o libbook40public.a book40public.c
		2)、g++ -o book40 book40.cpp libbook40public.a
			或者g++ -o book40 book40.cpp -L/root/20210127 -L/home/liyb/demo -lbook40public lpublic1 //-L指定静态库目录,-l指定静态库名,

		3)、./book40
	2、动态库:动态库在编译时并不会连接到目标代码中，而是在程序运行时才会载入。因此在程序运行时还需要指定动态库的目录，动态库的命名方式与静态库相似，前缀相同，为lib,后缀变为.so
		1)、g++ -fPIC -shared book40public.c -o libbook40public.so
		2)、g++ -o book40 book40.cpp -L/home/liyb/demo -lpublic//如果动态库与静态库同时存在，优先编译动态库
			如果出现：error while loading shared libraries: libbook40public.so: cannot open shared object file: No such file or directory
		这是因为采用了动态链接库的可执行文件程序在运行时需要指定动态库文件的目录，linux采用LD_LIBRARY_PATH环境变量指定动态库文件的目录
		如果要制定多个动态库
			export LD_LIBRARY_PATH=/root/20210127:/home/liyb/demo:.

	3、优缺点
	   1)、静态库
			优点:静态库相当于复制一份库文件到可执行程序中，不需要像动态库文件那样有动态加载和内存开销，也就是说采用静态编译的可执行程序运行更快
	   		缺点:
	   			a、静态链接生成的可执行文件比动态链接生成的大很多，运行时占用的内存也更多;
				b、库文件的更新不会反映到可执行文件中，可执行程序需要重新编译;
	   2)、动态库
	   		优点:
	   			a、相对于静态库，动态库在时候更新(修复bug、增加新功能)不需要重新编译;
	   			b、全部的可执行程序共享动态库代码，运行时占用的内存更少;
	   		缺点:
	   			a、使可执行程序在不同平台上移植变得复杂，因为它需要为每个不同的平台提供相应平台的共享库;
	   			b、增加了可执行程序运行时的时间和空间开销，因为应用程序需要在运行过程中查找依赖的库文件，并加载到内存中;
	 4、编译的优先级
	 	动态库在程序运行时被链接，故程序的运行速度和链接静态库的版本相比必然大打折扣，然而瑕不掩瑜，动态库的不足相对于它带来的好处在当今硬件下简直微不足道，所以链接程序在链接时一般是优先链接动态库的，除非用-static参数指定链接静态库。

*/



void jtk();
void dtk();

int main(){

	func();

}

void jtk(){ 
	printf("静态库\n");

}
void dtk(){
	printf("动态库\n");
	
	
}