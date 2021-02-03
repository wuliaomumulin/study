#!/bin/bash
# liyb
# 批量执行命令
if [ ! -f ip.txt ];then
	echo -e "\033[31mplease create ip.txt,the ip.txt contents as follow: \033[0m"
cat <<EOF
19.19.19.70
192.168.1.86
EOF
	exit
fi

if [ -z "$*" ];then
	echo -e "\033[31mUsage:$0 Command,example {rm /tmp/test.txt|mkdir /tmp/20210202} \033[0m"
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
	ssh -q -l root $IP "$*;echo -e '---------------------\nThe $IP Exec Command:$* success !';sleep 2"
done


