#!/bin/bash
# liyb
# 批量操作,移动文件、拷贝目录
# 	免密钥登录
# 		1、ssh-keygen
# 		ssh-copy-id -i /root/.ssh/id_rsa.pub 192.168.1.86
# 		ssh-copy-id -i /root/.ssh/id_rsa.pub 192.168.1.190

if [ ! -f ip.txt ];then
	echo -e "\033[31mplease create ip.txt,the ip.txt contents as follow: \033[0m"
cat <<EOF
19.19.19.70
192.168.1.86
EOF
	exit
fi

if [ -z "$1" ];then
	echo -e "\033[31mUsage:$0 Command,example {Src_Files|Src_Dir Des_Dir} \033[0m"
	exit
fi

count=`cat ip.txt|wc -l`
rm -rf ip.txt.swp
i=0
while((i<$count))
do
	i=`expr $i + 1`
	sed "${i}s/^/&${i} /g" ip.txt >> ip.txt.swp
	IP=`awk -v I="$i" '{if(I==$1)print $2}' ip.txt.swp`

	# 第一种方式
	#scp -r $1 root@${IP}:$2
	# 第二种方式
	# 目录同步完全一致
	rsync -aP --delete $1 root@${IP}:$2

done


