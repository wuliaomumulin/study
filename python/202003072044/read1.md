## 网络实践之网络抓包
***
### 一、网卡模式
- 混杂模式:接受经过网卡设备的数据
- 直接模式:只接受目的地址指向自己的数据
### 二、操作字节序
- 大端字节序(应用领域网络)
- 小端字节序(应用领域主机)
- 计算机电路先处理低位字节效率更高
- 人类习惯读写大端字节序
#### i.格式字符
|格式字符|类型|
|--|--|
|%s|字符串|
|%d|整数|
|%x|十六进制|
|%f|浮点数|

|格式字符|C++Python类型|标准大小(字节)|
|--|--|--|
|B|unsigned char/整数|1|
|H|unsigned short/整数|2|
|L|unsigned long/整数|4|
|s|char[]/字节串|~|

***
### 三、实现IP报文解析器
#### 1、前言:
- i、IP头部
图1
- ii、IP头部八位协议号说明
图2
- iii、Udp协议结构
图3
- iv、TCP协议头部结构
图4

#### 2、基础:
- i.新建包名computer_network,在包里创建一个struct_demo.py,然后运行分析结果。

...

#### 3、实现:
- i、在包里创建一个server.py作为程序入口

...

- ii.在computer_network包里创建一个processor子包，在processor子包里创建一个IP协议解析器文件parser.py

...

- iii.在processor包里创建一个协议TCP/UDP解析器文件trans.py

...
