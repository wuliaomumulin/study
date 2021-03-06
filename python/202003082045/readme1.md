## 网络层篇

### 前言、网络层概述:
        OSI参考模型        TCP/IP参考模型
		应用层             应用层
		表示层            
		会话层
		传输层             传输层
		**网络层**         **网络层**
		数据链路层         网络接口层
		物理层

| 模型名  | 目标 |
| 应用层|为计算机网络提供接口和服务|
|表示层|数据处理(便曾解码、加密解码)|
|会话层|管理(建立、维护、重连)通信会话|
|传输层|管理端到端的通信连接|
|网络层|数据路由(决定数据在网络中的路径)|
|数据链路层|管理相邻节点之间的数据通信|
|物理层|数据通信的光电物理特性|

- 相邻物理节点传输？
- 路由器
- 网络层IP协议相关
 - IP协议
 - 子网划分
 - 简单路由过程
- 网络层其他协议
 - ARP协议和RARP协议
 - ICMP协议
- IP的路由算法
 - 路由的概述
 - 内部网关路由协议 
 - 外部网关路由协议
### 一、IP协议详解
- (1)、虚拟互联网络
 - 实际的互联网络是错综复杂的
 - 物理设备通过使用IP协议，屏蔽了物理网络之间的差异
 - 当网络中的主机使用IP协议连接时，则无需关注网络细节
- (2)、意义:
 - IP协议使得复杂的实际网络变为一个虚拟互联的网络
 - IP协议使得网络层可以屏蔽底层细节而专注网络层的数据妆发
 - IP协议解决了在虚拟网络中数据报传输路径的问题
- (3)、IP协议组成:

		4位版本    4位首部长度    8位服务类型(TOS)            16位总长度字节
	                16位标识                       3位标志    13位片偏移
	    8位生存时间(TTL)       8位协议                      16位首部校验和
	                        32位源IP地址
	                        32位目的IP地址
	                        选项options(若有)

	                        IP数据

**版本**:占4位，指的是IP协议的版本，通信双方的版本必须一致，当前主流版本是4,即IPv4,也有IPv6
**首部长度位**：占4位，最大数值为15，表示的是IP首部长度，单位是32字(4个字节)，也即是IP首部最大长度为60字节
**8位服务类型**：我们不关注
**总长度**:占16位，最大长度为65535，表示的是IP数据报总长度(IP首部+IP数据)

**TTL**:占8位，表明IP数据报文在网络中的寿命，没经过一个设备。TTL减一，当TTL=0时，网络设备必须丢弃报文
**协议**:占8位，表明IP数据所携带的具体数据是什么协议，如(TCP、UDP等)

| 协议名  | ICMP | IGMP | IP | TCP  | UDP | OSPF | ... |
| ---- | ---- |----- | ----- | ----- | --- | -- | -- |
| 字段值 |  1 |  2  |  4 |  6 | 17 |  89 | ... |

### 二、IP协议转发流程
- (1)、hop-by-hop(逐跳)
- (2)、路由表简介:
 - A通过网卡发送数据帧
 - 数据帧到达路由器，路由器取出前6字节
 - 路由器匹配MAC地址表，找到对应的网络接口
 - 路由器往该网络接口发送数据帧 
- (3)、IP协议的转发流程
 - 数据帧的每一跳的MAC地址都在变化
 - IP数据报每一跳的IP地址始终不变

### 三、ARP和RARP协议
#### (1)、ARP(address Resolution Protocol)地址解析协议
`网络层IP32位地址`- ARP  ->`数据链路层MAC48位地址`

- 位置和构成
| 目的地址 | 源地址 | 类型 | 帧数据 | CRC |
| ---- | ---- |----- | ----- | ----- |
| 6 | 6 | 2 | 46-1500 | 4 |

|类型 0806| ARP请求/应答 | PAD |
| ---- | ---- |----- |
| 2 | 28 | 18  |

|硬件类型|协议类型|标记|发送端以太网地址|发送端IP地址|目的端以太网地址|目的端IP地址|
| ---- | ---- |----- | ----- | ----- | -- | -- |
|2|2|4|6|4|6|4|

####(2)、RARP(Reverse Address Resolution Protocol)逆地址解析协议
 `数据链路层MAC48位地址` - RARP  -> `网络层IP32位地址`

#### (3)、共性
- (R)ARP协议是TCP/IP协议栈里面基础的协议
- ARP和RARP的操作对程序员是透明的
- 理解(R)ARP协议有助于理解网络分层的细节

### 四、IP地址的子网划分
#### (1)、分类的IP地址
- A类地址:8位网络号，24位主机号,0
- B类地址:16位网络号，16位主机号,10
- A类地址:24位网络号，8位主机号,110

|...|最小网络号|最大网络号|子网数量|最小主机号|最大主机号|主机数量|
|---|--|--|--|--|--|--|
|A|1|127(01111111)|2^7-2|0.0.1|255.255.254|2^24-2|
|B|128.1|191.255|2^14-1|0.1|255.254|2^16-2|
|C|192.0.1|223.255.255|2^21-1|1|254|2^8-2|

- 特殊的主机号
 - 主机号全0表示当前网络段，不可分配为特定主机
 - 主机号全唯1标识广播地址，像当前网络段所有主机发送信息
- 特殊的网络号
  - A类地址网络段全0(00000000)表示特殊网络
  - A类地址网络段后七位全1(01111111:127) 表示回环地址
  - B类地址网络段(10000000.00000000:128.0)是不可使用的
  - C类地址网络段(192.0.0)是不可使用的
  - D类地址:前四位1110...
  - E类地址: 前四位1111..
 
 > 127.0.0.1，通常称为本地回环地址(Loopback Address),不属于任何一个有类别地址类，它代表设备的本地虚拟端口，所以默认被看做是永远不会宕机的接口。在Windows操作系统中也有相应的定义，所以通常在安装网卡前就可以ping通这个本地回环地址。一般都会用来检测本地网络协议、基本数据接口是否正常。 
#### (2)、划分子网
|网络号|子网号|主机号|
|--|--|--|
#### (3)、无分类编址CIDR
- CIRR中没有A、B、C类网络号和子网划分的概念
- CIDR将网络前缀相同的IP地址称为一个“CIDR地址块”
|网络前缀|主机号|
|--|--|

**网络前缀是任意位数的**
- 斜线记法:
 - 193.10.10.129/25

|CIDR前缀|掩码点分十进制| 地指数 |
|--|--|--|
|/13|255.248.0.0|512k|
|/14|255.252.0.0|256k|
|/15|255.254.0.0|128k|
|/16|255.255.0.0|64k|
|/17|255.255.128.0|32k|
|/18|255.255.192.0|16k|
|/19|255.255.224.0|8k|
***相比原来子网划分更加灵活***
### 五、网络地址转换NAT技术
- IPv4最多只有40+亿个IP地址
- 早期IP地址的不合理规划导致IP号浪费
- 内网地址
 - 内部机构使用
 - 避免与外网地址冲突
- 外网地址
 - 全球范围使用
 - 全球公网唯一
#### (1)、三类内网地址
- 10.0.0.0~10.255.255.255 (支持千万数量级设备)
- 172.16.0.0~172.31.255.255 (支持百万数量级别设备)
- 192.168.0.0~192.168.255.255 (支持万数量级设备)
#### (2)、网络地址转换技术(Network Address Translation) 
 - NAT技术用于多个主机通过一个公网IP访问互联网的私有网络中
 - NAT技术减缓了IP地址的消耗，但是增加了网络通信的复杂度
 
### 五、ICMP详解
- 网际控制报文协议(internet Control Message Protocol)
- ICMP协议可以报告错误信息或者异常信息

		IP首部    IP数据报的数据
	         IP数据报的数据
	    帧首部  帧数据    帧尾部
		ICMP报文首部 ICMP报文数据

- 构成

|8位类型|8位代码|16位效验和|
|--|--|--|
|ICMP报文|...|...|

 - 差错报告报文

|ICMP报文种类|类型的值|报文类型|具体代码|
|--|--|--|--|
|差错报告报文|3(终点不可达)|网络不可达|0|
|...|...|主机不可达|1|
|...|5(重定向)|对网络重定向|0|
|...|...|对主机重定向|1|
|...|11|传输超时| - |
|...|12|坏的IP头| 0 |
|...|...|缺少必要参数 | 1 |

 - 询问报文

|ICMP报文种类|类型的值|报文类型|具体代码|
|--|--|--|--|
|询问报文|0或8|回应(Echo)请求或应答 |-|
|询问报文|13或14|时间戳(Timestamp)请求或应答 |-|

### 六、ICMP应用
#### (1)、Ping应用
`向目标IP发送ICMP的32字节(最少20字节头部)询问报文`
- Ping回环地址127.0.0.1
- Ping网关地址
- Ping远端地址
#### (2)、Traceroute应用
> Traceroute可以探测IP数据报在网络中走过的路径。
当数据报文的TTL=0时，ICMP终点不可达差错报文，网络设备丢弃该报文。
初始封装一个TTL=1的数据报文，当终点跳的目标主机是不可达状态时，就会循环或再次封装一个TTL+=1的数据报文，直到找到目标主机的确切位置。
tracert github.com

### 七、网络层路由概述
#### 一、算法本质:
> 下一跳地址是怎么来的？
  下一跳地址是唯一的吗？
  下一跳地址是最佳的吗?
  路由器这么多，他们是怎么协同工作的？
	**路由算法实际上是图论算法，但是由于网络环境复杂，路由算法要比图论算法复杂**

- 要求:
 - 算法是正确的，完整的；
 - 算法在计算上应该尽可能的简单;
 - 算法可以适应网络中的变化
 - 算法是稳定和公平的
- 规模:
 - 互联网的规模是庞大的
 - 互联网的环境是复杂的
#### 二、自治系统(Autonomous System)
对互联网进行划分
- 一个自愈系统(AS)是处于一个管理机构下的网络设备群
- AS内部网络自行管理，AS对外提供一个或者多个出(入)口
- 自治系统内部路由的协议称为:内部网关协议(RIP、OSPF)
- 自治系统外部路由的协议称为:外部网关协议(BGP)

### 八、内部网关协议之RIP协议
#### 一、距离矢量(DV)算法
- 每一个节点使用两个向量Di和Si
- Di描述的是当前节点到别的节点的距离
- Si描述的是当前节点到别的节点的下一个节点是什么
- 每一个节点与相邻的节点交换向量Di和Si的信息
- 每一个节点都会根据交换的信息更新自己的节点

>  Dij = min(dix+dij)
Di1表示从节点i到节点1的距离
Si1表示从节点i到节点1的下一个节点
n标识节点的数量

#### 二、RIP协议过程
- RIP(Routeing Information Protocol)协议
- RIP是使用DV算法的一种协议
- RIP协议把网络的跳数(hop)作为DV算法的距离
- RIP协议每隔30s交换一次路由信息
- RIP协议认为跳数>15的路由则为不可达路由
- RIP协议缺点:
 - 故障信息传递慢,更新收敛时间过长;
 - 自己不思考、视野不够;
- RIP协议特点:
 - 实现简单，开销很小;
 - 限制了网络的规模;

> 1、路由器初始化路由信息(两个向量Di和Si);
2、对相邻路由器X发过来的信息，对信息的内容经行修改(下一跳地址设置为X,所有距离加1);
 i.检索本地路由，将信息中新的路由插入到路由表里面
 ii.检索本地路由，对于下一跳为X的，更新为修改后的信息
 iii.检索本地路由，对比相同目的地的距离，如果新信息的距离更小，则更新本地路由表信息
3、如果三分钟没有收到相邻的路由信息，则把相邻路由设置为不可达(16跳)
 
### 九、dijkstra(迪杰斯特拉)算法
> Dijkstra是著名的图算法
 Dijkstra算法解决有权图从一个节点到其他节点的最短路径问题
 "以起始点为中心，向外层层扩展"

- 最短路径问题?
 - 1、初始化两个集合(S,U)(S为只有初始顶点点A的集合，U为其他顶点集合)
 - 2、如果U不为空，对U集合顶点进行距离的排序，并取出距离顶点A最小的一个顶点D
i.将顶点D纳入S集合
ii.更新顶点D到达U集合所有点的距离(如果距离更小则更新，否则不更新)
iii.重复ii步骤
 - 3、直到U集合为空，算法完成

### 十、内部网关路由协议之OSPF协议
> 一个为解决RIP协议缺点产生的协议

#### (1)、链路状态(LS)协议
- 向所有路由器发送信息
- 消息描述该路由器与相邻路由器的链路状态(包含距离、时延、带宽...,由网络管理人员决定，连接另一个路由器的代价)
- 只有链路状态发生变化时，才发送更新消息

#### (2)、OSPF协议过程
- OSPF(Open Shortest Path First:开放最短路径优先)
- OSPF协议的核心是dijkstra算法
 - 向所有路由器发送信息
`获得网络中所有的消息`-->`每个路由都可以获得完整的网络拓扑`
**也称为链路状态数据库**
**"链路状态数据库"是全网一致的**
 - 消息描述该路由器与相邻路由器的链路状态
`OSPF协议更加客观、更加先进`
 - 只有链路状态发生变化时，才发送更新消息
 `减少了数据交换，更快收敛`
- OSPF的五种消息类型
 - 问候消息(Hello)
 - 链路状态数据库描述信息
 - 链路状态请求信息
 - 链路状态更新消息
 - 链路状态确认消息
`路由器接入网络`-->`路由器向邻居发送问候消息`-->`与邻居交流链路状态数据库`-->`广播和更新未知路由`

|RIP协议|OSPF协议|
|--|--|
|从邻居看网络|整个网络的拓扑|
|在路由之间累加距离|Dijkstra算法计算最短路径|
|频繁、周期更新，收敛很慢|状态变化更新，收敛很快|
|路邮件拷贝路由信息|路由间传递状态信息，自行计算路径|

### 十一、外部网关路由协议之BGP协议
#### (1)、BGP(Board Gateway Protocol:边际网关协议)
- BGP协议是运行在AS之间的一种协议
 - 互联网的规模很大
 - AS内部使用不同的路由协议
 - AS之间往往需要考虑除网络特征以外的一些因素(政治、安全...)
- BGP协议能够找到一条到达目的地比较好的路由
#### (2)、BGP发言人(speaker)
- BGP并不关心内部网络拓扑
- AS之间通过BGP发言人进行路由信息交换和交流信息
- BGP Speaker可以人为配置策略


