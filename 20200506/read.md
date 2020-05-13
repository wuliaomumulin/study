## 一、dmidecode获取设备唯一标识
### System Information设备信息
```
dmidecode -t 1
```
### 设备唯一标识uuid
```
dmidecode -s system-uuid
```
### serial number码
```
dmidecode -s system-serial-number
dmidecode -t system|grep -i 'serial number'
```
## 二、rabbitMQ简单使用
```
rabbitmqctl stop_app #关闭应用
rabbitmqctl reset #清除应用
rabbitmqctl start_app #启动应用
```
### 二、以普通用户执行管理员权限允许执行的命令
#### 1、复制管理员权限，重命名为业务用户,使业务用户可以通过脚本执行sudo命令
```
# User privilege specification
root    ALL=(ALL:ALL) ALL
liyb    ALL=(ALL:ALL) ALL
```
#### 2、业务shell
```
echo "password" | sudo -S cat /etc/sudoers
```
