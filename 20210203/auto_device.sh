#!/bin/bash
# liyb
# 获取硬件信息
echo -e "\033[34m \033[1m"

cat <<EOF
----------------------welcome to use system Collect----------------
---------------------
EOF
ip=`ifconfig |grep "broadcast"|tail -1|awk '{print $2}'|cut -d: -f 2`
# cpu=`cat /proc/cpuinfo |grep 'model name'|tail -1|awk -F: '{print $2}'|sed 's/^ //g'|awk '{print $1,$2,$3,$(NF-2),$NF}'`
cpu=`cat /proc/cpuinfo |grep 'model name'|tail -1|awk -F: '{print $2}'|sed 's/^ //g'|awk '{print $1,$2,$3,$4,$NF}'`
cpu_physical=`cat /proc/cpuinfo |grep 'physical id'|sort|uniq -c|awk '{print $1}'`
serv=`hostname|tail -1`
disk=`fdisk -l|grep "Disk"|grep -v "identifier"|awk '{print $2,$3,$4}'|sed 's/,//g'|sed ':t;N;s/\n/\;/;b t'`
mem=`free -m |grep "Mem"|awk '{print "Total",$1,$2"M"}'`
load=`uptime |awk '{print "Current Load: "$(NF-2)}'|sed 's/\,//g'`

echo -e "\033[32m---------------------\033[1m"
echo IPADDR:${ip}
echo HOSTNAME:$serv
echo CPU_INFO:${cpu} X${cpu_physical}
echo DISK_INFO:$disk
echo MEM_INFO:$mem
echo LOAD_INFO:$load

echo -e "\033[32m---------------------\033[0m"
echo -n -e "\033[36mYou want to write the databases?\033[1m";read ensure
if [ "$ensure" == "yes" -o "$ensure" == "y" -o "$ensure" == "Y" ];then
	echo "-----------------------------------"
	# 执行sql
	mysql -uroot -p123456 -D test -e "insert into audit_audit_system values('','${ip}','$serv','${cpu} X${cpu_physical}','$disk','$mem','$load','')"
	
	echo -e '\033[31mmysql -uroot -p123456 -D test -e '''"insert into audit_audit_system values('','${ip}','$serv','${cpu} X${cpu_physical}','$disk','$mem','$load','')" ''' \033[0m'

else
	echo "wait exit"
	exit
fi
# 另一张方式获取系统信息
# dmidecode -t system


# 数据表结构

#CREATE TABLE `audit_audit_system` (
#`id` int(11) NOT NULL AUTO_INCREMENT,
# `ip_info` varchar(50) NOT NULL,
# `serv_info` varchar(50) NOT NULL,
# `cpu_info` varchar(50) NOT NULL,
# `disk_info` varchar(50) NOT NULL,
# `mem_info` varchar(50) NOT NULL,
# `load_info` varchar(50) NOT NULL,
# `mark_info` varchar(50) NOT NULL,
# PRIMARY KEY (`id`),
# UNIQUE KEY `ip_info` (`ip_info`),
# UNIQUE KEY `ip_info_2` (`ip_info`)
# ) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

