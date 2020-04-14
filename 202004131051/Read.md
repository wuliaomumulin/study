## 一、mysql
### （1）、日志落盘
```
show variables like "%general_log%";
set global general_log=off
tail -n 50  /var/lib/mysql/andi.log
```
### （2）、sql关联分组查询统计 
```
sql_mode = STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
```

## 二、library调用与被调用
yaf-library管理
library方面
Test\Test1.php
```
class Test_Test1{
	public function test(){
		echo __FUNCTION__;
	}
}
```
controller方面
```
$Test = new Test_Test1();
$Test->test();
```
### 三、Ubuntu开机启动文件编辑
```
/etc/profile
```

### 四、主机性能
```
top -bn1|grep load|awk '{printf "CPU Load:%2f\n",$(NF-2)}
free -m|awk 'nr==2{printf "Memory Usage:%s%sMB(%.2f%%)\n"}'
```
### 五、监控网卡流量
```
watch ifconfig ens33
watch cat /proc/net/dev
```
### 六、snmp协议
```
yum list all|grep net-snmp* # 列出可以使用的包
rpm -qa|grep net-snmp* # 已安装的包
yum install --skip-broken -y net-snmp net-snmp-utils # 安装
rpm -ql net-snmp-5.7.2-38.el7_6.2.x86_64 # 查看单个安装包具体情况
```

#### 1、使用
```
snmpget -v协议版本 -c 指定密码 oid
snmpget -v2c -c public 192.168.1.86 .1.3.6.1.2.1.1.1.0 
oid

snmpwalk 127.0.0.1 -c public -v 2c # 抓取本机全部
snmpwalk 127.0.0.1 -c public -v 2c 1.3.6.1.2.1.1.1 # 抓取操作系统
snmpwalk 127.0.0.1 -c public -v 2c 1.3.6.1.2.1.1.3 # 抓取开机时间
snmpwalk 127.0.0.1 -c public -v 2c 1.3.6.1.2.1.1.5 # 抓取主机名称
```
### 七、vi与bom
```
:set nobomb # 去掉BOM标记
:set bomb # 设置bom
:set bomb? #查询当前UTF-8编码的文件是否有BOM标记
:%!xxd # 以16进制模式打开文件
:%!xxd -r # 将以16进制格式打开的文件返回文本模式编辑,所以，先用第一个命令将文件以16进制打开，删除文件开头的EF BB BF，然后再用第二个命令返回文本模式。
```
