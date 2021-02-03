#!/bin/bash
# auto_deny_ip.sh
# 自动拒绝ssh failed ip address
# liyb
SEC_FILE=/var/log/secure
#以下为截取secure文件恶意ip远程登录22端口,大于等于4次就写入防火墙，禁止以后再登录该机器的22端口
IP_ADDR=`tail -n 1000 /var/log/secure | grep "Failed password" | egrep -o "([0-9]{1,3}\.){3}[0-9]{1,3}" | sort -nr | uniq -c | awk ' $1>=4 {print $2}'`
IPTABLE_CONF=/etc/sysconfig/iptables
echo
cat <<EOF
---------------欢迎使用ssh login drop failed-----------------
------------------------------------------------------------
------------------------------------------------------------
EOF
for i in `echo $IP_ADDR`
do
  #查看iptables配置文件是否含有提取的IP信息
  cat $IPTABLE_CONF | grep $i > /dev/null

  if [ $? -ne 0 ];then
    # 判断IPtables配置文件里面是否存在已拒绝的IP，如果不存在，就添加相应条目
    # sed a参数是匹配之后加入行的意思
    sed -i "/lo/a -A INPUT -s $i -m state --state NEW -m tcp -p tcp --dport 22 -j DROP" $IPTABLE_CONF
  else
    echo "this is $i exists in iptables,please exit..."
  fi
done
#最后重启iptables生效

