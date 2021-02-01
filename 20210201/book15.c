#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <utime.h>
#include <stdio.h>	

/**
* 一、access库函数 判断问文件或者目录的访问权限
* stat 获取指定文件或者目录的信息，并将文件保存到结构体buf中，执行成功返回0，失败返回-1

 二、stat库函数
struct stat{
		unsigned long   st_dev;         文件的设备编号.  
        unsigned long   st_ino;         文件的inode.  
        unsigned int    st_mode;        文件的类型和存储的权限.  
        unsigned int    st_nlink;       连到该文件的硬链接数目.  
        unsigned int    st_uid;         文件所有者的用户识别码.  
        unsigned int    st_gid;         文件所有者的组识别码. 
        unsigned long   st_rdev;        若此文件为设备文件，则为其设备编号.  

        long            st_size;        文件大小，以字节计算 Size of file, in bytes.  
        int             st_blksize;     Optimal block size for I/O. 文件系统的IO 缓存区的大小 
        int             __pad2;
        long            st_blocks;      Number 512-byte blocks allocated. 占用文件区块的个数，每一区块大小为512个字节 
        long            st_atime;       Time of last access.文件最近一次被存取或被执行的时间，一般只有在用mknod、utime、read、write与truncate时改变  
        unsigned long   st_atime_nsec;
        long            st_mtime;       Time of last modification. 文件最后一次被修改的时间，一般只有在用mknod、utime、write时才会改变  
        unsigned long   st_mtime_nsec;
        long            st_ctime;       Time of last status change. i-node 最近一次被修改的时间，此参数会在文件所有者、组、权限被更改时更新 
};

 三、utime库函数
	用于更改文件的修改时间和更改时间
struct utimbuf {
               time_t actime;        access time 
               time_t modtime;       modification time 
           };
  四、rename库函数
  #include <stdio.h>
  五、remove库函数
  成功返回0，失败返回-1
  	
*/




int main(int argc,char *argv[],char *envp[]){
	
	if(argc != 2){ printf("参数个数%d!=2\n",argc);return -1; }

	//判断是否文件是否存在
	if(access(argv[1],F_OK)!=0){return -1;}

	struct stat ststat;
	//获取文件结构体
	if(stat(argv[1],&ststat)!=0) return -1;

	if(S_ISREG(ststat.st_mode)){printf("%s是一个文件\n",argv[1]);}
	if(S_ISDIR(ststat.st_mode)){printf("%s是一个目录\n",argv[1]);}
}