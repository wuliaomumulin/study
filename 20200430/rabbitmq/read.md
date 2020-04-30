五元组、端口、IP、安全事件列表、威胁情报.

# mysql日志落盘
show variables like "%general_log%";
set global general_log=off
tail -n 50  /var/lib/mysql/andi.log

一、框架的问题
1、PDO事务使用貌似不管用


20200330 
编写资产增加和删除的接口
20200331
编写插件列表、以及增加和删除接口;
20200401
UDP采集器列表、以及增加和删除接口;
SNMP协议的使用和学习;
插件相关数据表新增部分字段和逻辑;
大屏前期的准备
20200402
威胁情报的列表和删除
重新修改数据库调用注释的表示方法;
20200403
系统状态接口开发、以及文档编写;
重新修改威胁情报的表结构;
针对报表上传文件的修改;

态势感知
采集器

2、这两张表
host_properties software_cpe

输出当前数组:key current


1、继承父类;
2、子类写方法;

snmp协议
```
yum list all|grep net-snmp* # 列出可以使用的包
rpm -qa|grep net-snmp* # 已安装的包
yum install --skip-broken -y net-snmp net-snmp-utils # 安装
rpm -ql net-snmp-5.7.2-38.el7_6.2.x86_64 # 查看单个安装包具体情况
```

snmpget -v协议版本 -c 指定密码 oid
snmpget -v2c -c public 192.168.1.86 .1.3.6.1.2.1.1.1.0 
oid

snmpwalk 127.0.0.1 -c public -v 2c # 抓取本机全部
snmpwalk 127.0.0.1 -c public -v 2c 1.3.6.1.2.1.1.1 # 抓取操作系统
snmpwalk 127.0.0.1 -c public -v 2c 1.3.6.1.2.1.1.3 # 抓取开机时间
snmpwalk 127.0.0.1 -c public -v 2c 1.3.6.1.2.1.1.5 # 抓取主机名称


yaf_menu
yaf_user
user_role
system_config
sys_log
user_report
udp_sensor



20200407
告警和安全事件
20200409
大数据平台调试问题
20200410
大屏的网卡数据抓取和展示;	
开会讨论筛选的重新调整

一、library调用与被调用
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
二、Ubuntu开机启动文件编辑
```
/etc/profile
```
20200413
1、确认以及排查大数据平台线上环境问题;
2、解决禅道的一堆bug;
```
set nobomb|bomb|bomb?
```
20200414

1、图片上传的问题

厂级平台权限apache2默认用户:www-data

2、数据发送状态;
kb|MB
3、默认大屏设置;

20200415
websocket的问题
apt-get install php7.2-gd

20200416
资产和采集器详情的接口;
相关功能的前端督促;
学习和bug
	psr4相关规范的学习;
	input的为空判断标准
	
B比特位-->KB->MB->GB


20200417
1、xml数据整理

    //修改管理口IP
    public function editmanagementipAction(){
        $str = APP_PATH.'/outside/modify_configure.py';
        `python {$str}`;
         jsonResult([],'操作成功');
    }


1、udp采集器修改;
2、采集器分类;
20200422
3、目录整合并且权限验证
4、大数据平台默认路由不生效的问题;
5、修改厂级平台表注释;
20200423
1、了解和解析xml
2、大数据平台了解导入功能异常；
20200424
3、代理转发
4、将有关es的ifarme的url地址全部换成基于服务端IP的动态获取


20200426
# 权限配置
chmod +x /work/web/outside/cpu_mem_disk.sh
chmod -R 777 /work/web/log
# 以守护的方式运行任务
nohup php /work/web/application/bin/network.php > /dev/null 2&1 &

大屏字段调整
1、告警列表第一列新增意图和策略,对应字段category
2、安全事件第一列新增事件名称,对应字段plugin_sname
3、将告警列表和安全列表的时间字段放到最后一列;

4、登陆时把查询的theme1字段删去，以保证数据库升级版本没问题;

20200427
1、告警规则;
2、资产修改、显示字段大小写、字段错别字修正;
20200428
1、告警大屏查询慢原因排查；


20200429
# mq使用

# 原料:


# 安装mq客户端
./configure --prefix=/usr/lib/rabbitmq-c
make&&make install

# 安装扩展
./configure --with-php-config=/usr/bin/php-config --with-amqp --with-librabbitmq-dir=/usr/lib/rabbitmq-c/
make&&make install

# 配置添加
vim /etc/php/7.2/cli/conf.d/amqp.ini

[PHP Modules]
extension=amqp.so


安装rabbitmq

20200430
高级搜索、资产字段新增;