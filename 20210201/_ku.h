#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>
#include <errno.h>
//access
#include <unistd.h>

/*
	比较大小
 */
int max(const int i,const int ii);

int max(const int i,const int ii)
{
	if(i>ii) return i;
	return ii;
}