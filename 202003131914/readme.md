# 总结-编程必备知识
***
## 一、计算机组成原理
### (1).背景篇
#### A.计算机的发展史
##### a、电子管计算机
##### b、晶体管计算机
##### c、集成电路计算机
##### d、超大规模集成电路计算机
##### e、未来的计算机
#### B.计算机的分类
##### a、CPU算力
##### b、CPU集成核数
#### C.计算机的体系与结构
##### a、冯诺依曼体系
- 五大组成部分
 - 超级计算机
 - 大型机
 - 迷你计算机(普通服务器)
 - 工作站
 - 微型计算机
- 冯诺依曼瓶颈
##### b、现代计算机的结构		 
- 现代计算机的结构
 - 以存储器为核心的冯诺依曼结构计算机
#### D.计算机的层次与编程语言
##### a.程序翻译
- 编译器
- 相关语言
 - C/C++
 - Object-C
 - Golang
##### b.程序解释
- 解释器
- 相关语言
 - Python
 - Php
 - Javascript
##### c.混合性语言
- Java
- C#
##### d.计算机的层次和编程语言
###### I.七个层次 
###### II.编程语言
- 物理机器
 - 系统软件
 - 应用软件
- 虚拟机器
#### E.计算机的计算单位
##### a.容量单位
- 名称
- 进制&运算
##### b.速度单位
- 网络速度
- 计算速度
#### F.计算机的字符和编码集
##### a.发展历史
- 早起英美的ASCII码
- Externet ASCII码
- 国际化
##### b.中文编码集
- GB2312
- GBK
- Unicode
 - UTF-8
### (2).组成篇
#### A.计算机的总线
##### a.总线的概览
- 是什么
- 怎么用
- 总线的分类
 - 片内总线
 - 系统总线
##### b.总线的仲裁
- 链式查询
- 计时器定时查询
- 独立请求
#### B.计算机的输入输出设备
##### a.常见输入输出设备
- 输入设备
 - 字符输入设备
 - 图像输入设备
- 输出设备
##### b.输入输出接口的通用设计
- 数据线
- 状态线
- 命令线
- 设备选择线
##### c.CPU和IO设备的通信
- 程序中断
- DMA
#### C.计算机存储器规范
##### a.存储器的分类
- 指标
 - 读写速度
 - 存储容量
 - 价格
##### b.层次结构
- 高速缓存
- 主存
- 辅存
##### c.局部性原理
#### D.计算机的主存储器和辅助存储器
##### a.主存
- 简介&构造&原理
##### b.辅存
- 简介&构造&原理
- 寻道算法
 - 先来先服务算法
 - 最短寻道时间优先
 - 扫描算法
 - 循环扫描算法
#### E.计算机的高速缓存
##### a.工作原理
- 字&字块
- 运算
 - 字数、字块数
 - 命中率
 - 访问效率
##### b.替换策略
- 随机算法
 - 先进先出算法
 - 最不经常使用算法
 - 最近最少使用算法
#### F.计算机的指令系统
##### a.形式
- 操作码&地址码
- 一级指令
- 二级指令
- 三级指令
##### b.操作类型
- 数据传输
- 算法逻辑
- 移位操作
- 控制指令
##### c.寻址方式
- 指令寻址
 - 顺序寻址
 - 跳跃寻址
- 数据寻址
 - 立即寻址
 - 直接寻址
 - 间接寻址
#### G.计算机的控制器
- 程序计数器 
- 时序发生器
- 指令译码器
- 寄存器
#### H.计算机的运算器
- 数据缓冲器
- ALU
- 通用寄存器
- 状态字寄存器
- 总线
#### I.计算机指令的执行过程
***
### (3).计算篇
#### A.进制运算的基础
- 什么是进制
- 二进制转十进制的方法
- 十进制转二进制的方法
#### B.有符号数和无符号数
- 有符号数的存储方式
- 原码表示法的定义 
- 原码表示法的不足 
#### C.二进制的补码表示法
- 补码表示法定义
- 求补码的运算
- 直接求补码运算的不足
#### D.二进制的反码表示法
- 反码表示法的定义
- 反码标识法的作用
#### E.小数的二进制补码表示法
#### F.定点数与浮点数
- 定点数的表示方法
- 浮点数的表示方法
 - 格式
 - 规范
 - 规范化
- 两者的对比
 - 范围
 - 精度
 - 运算的复杂性
#### G.定点数的加减运算
- 数值位与符号位一起运算
- 双符号位判断溢出
#### H.浮点数的加减运算
- 对阶
- 尾数求和
- 尾数规范化
- 舍入
- 溢出判断
#### I.浮点数的乘除法运算
- 阶码运算
- 尾数运算
- 尾数规范化
- 舍入
- 判断溢出
## 二、操作系统
### (1).基础篇
#### A.操作系统概览
- 什么是操作系统
- 为什么需要操作系统
- 操作系统的基本功能
 - 管理资源
 - 抽象资源
 - 提供用户操作接口
- 基本概念
 - 并发性
 - 共享性
 - 虚拟性
 - 异步性
#### B.进程管理
##### a.进程实体
- 为什么需要进程
- 进程的本质
 - 状态
 - 优先级
 - 程序计数器
 - 上下文数据
 - 内存信息
 - ...
- 进程和线程
 - 关系
 - 区别
##### b.五状态模型
- 就绪状态
- 阻塞状态
- 执行状态
- 创建状态
- 终止状态
##### c.进程同步
- 生产者-消费者问题
- 哲学家进餐问题
- 临界问题
- 进程间同步的方法
- 线程间同步的方法
##### d.Linux的进程管理
- Linux进程相关概念
 - 前台进程
 - 后台进程
 - 守护进程
 - 进程ID
 - 父子进程
- 操作进程的相关指令
 - fg/bg
 - jobs
 - nohup
 - ps
 - kill
#### C.作业管理
##### a.进程调度
- 概述
 - 就绪队列的委派机制
 - 选择运行进程的委派机制
 - 新老进程的上下文切换机制
- 调度算法
 - 先来先服务调度算法
 - 短进程优先调度算法
 - 高优先权优先调度算法
 - 时间片轮转算法
##### b.死锁
- 破坏必要条件
- 银行家算法
#### D.存储管理
##### a.内存分配和回收
- 内存分配的过程
 - 单一连续分配
 - 固定分区分配
 - 动态分区分配
- 内存回收的过程
 - 四种情况
##### b.段页式存储管理
- 页式存储管理
- 段式存储管理
- 段页式存储管理
##### c.虚拟内存
- 概述
 - 是什么
 - 为什么
- 局部性原理
- 置换算法
 - FIFO
 - LRU
 - LFU
##### d.Linux的存储管理
- buddy内存管理算法
 - 伙伴
 - 分配&回收过程
- 交换空间
 - 交换空间VS虚拟空间
#### E.文件管理
##### a.操作系统的文件管理
- 文件的逻辑结构
 - 有结构文件
 - 无结构文件
 - 顺序文件
 - 索引文件
- 辅存的存储空间分配
 - 连续分配
 - 链接分配
 - 索引分配
- 目录管理
 - 唯一路径
 - 文件描述信息
##### b.Linux的基本操作
- Linux目录
- 创建、删除、读取、写入
- 文件类型
 - 套接字
 - 普通文件
 - 目录文件
 - 符号链接
 - 设备文件
 - FIFO
##### c.Linux的文件系统
		- 文件系统概览
		 - FAT
		 - NTFS
		 - Ext2/3/4
		- Ext文件系统
		 - Boot Sector
		 - Block Group
		  - Super Block
		  - Inode Bitmap
		  - Block Bitmap
		  - Inode Table
		  - Data Block  
#### F.设备管理
- IO设备分类
 - 使用特性分类
 - 信息交换的单位
 - 设备共享属性
 - 传输速率
- IO设备的缓存区
- SPOOLing技术

### (2).提升篇
#### A.线程同步实践
##### a.互斥量
##### b.自旋锁
##### c.读写锁
##### d.条件变量
#### B.进程同步实践
##### a.共享内存
##### b.Unix域套接字
#### C.高级概念
##### a.用户态和内核态
##### b.上下文切换
##### c.协程与线程
## 三、计算机网络
### (1)、概述篇
#### A、什么是计算机网络
#### B、计算机网络的分类
#### C、计算机网络的发展历史
##### a、世界互联网的发展历史
- 单个网络
- 三层结构
- 多层结构ISP
##### b、中国互联网的发展历史
- 1980年
- 1989年
- 1994年
- 中国的互联网企业
#### D、计算机网络的层次结构
##### a、层次结构设计的基本原则
- 相互独立
- 灵活性
- 耦合度
##### b、OSI七层模型
##### c、TCP/IP四层模型
#### E、现代互联网的网络拓扑
##### a、边缘部分
##### b、核心部分
##### c、C/S模式
##### d、P2P模式
#### F、计算机网络的性能指标
##### a、速率
##### b、时延
- 发送时延
- 传输时延
- 排队时延
- 处理时延
##### c、往返时间RTT
#### G、物理层概述
##### a、物理层的作用
- 相关设备
##### b、信道的基本概念
- 单工信道
- 半双工信道
- 全双工信道
##### c、分用-复用技术
#### H、数据链路层概述
##### a、主要功能
- 封装成帧
 - 数据帧的结构
- 透明传输
- 差错监测
#### I、数据链路层的差错监测
##### a、奇偶校验码
##### b、循环冗余校验码
#### J、最大传输单元MTU
##### a、MTU
##### b、路径MTU
#### K、以太网协议衔接
##### a、MAC地址
##### b、以太网协议
### (2)、网络层篇
#### A、网路层的基本功能
##### a、数据路由
#### B、IP协议详解
##### a、虚拟网络技术
##### b、IP协议
- IP地址
- IP报文格式
#### C、IP协议的转发流程
##### a、路由表
##### b、转发流程
#### D、ARP和RARP协议
##### a、ARP协议
##### b、RARP协议
#### E、IP地址的子网划分
##### a、分类的IP地址
- A类
- B类
- C类
- 特殊的网络号
- 特殊的主机号
##### b、划分子网
- 子网掩码
##### c、无分类地址CIDR
#### F、网络地址转换NAT技术
##### a、内网地址
##### b、外网地址
##### c、端口映射
#### G、ICMP协议详解
##### a、功能
##### b、分类
##### c、ICMP报文结构
#### H、ICMP报文的应用
##### a、Ping应用
##### b、Traceroute应用
#### I、网络层的路由概述
##### a、路由表更新问题
##### b、路由算法
##### c、自治系统(AS)
#### J、内部网关路由协议之RIP协议
##### a、距离矢量(DV)算法
##### b、RIP协议的过程
##### c、RIP协议的弊端
#### K、Dijkstra算法
##### a、最短路劲问题
#### L、内部网关路由协议之OSPF协议
##### a、链路状态(LS)协议
##### b、OSPF协议的过程
#### M、外部网关路由协议之BGP协议
##### a、政策
##### b、安全
### (3)、传输层篇
#### A、传输层的主要功能
##### a、进程与进程的通信
##### b、端口的概念
#### B、UDP协议详解
##### a、功能
##### b、特点
##### c、报文结构
#### C、TCP协议详解
##### a、功能
##### b、特点
##### c、报文结构
#### D、可靠传输的基本原理
##### a、停止等待协议
##### b、连续ARQ协议
##### c、超时重传计时器
#### E、TCP协议的可靠传输
##### a、滑动窗口
##### b、累计确认
##### c、选择重传
#### F、TCP协议的流量控制
##### a、窗口
##### b、定时器
#### G、TCP协议的拥塞控制
##### a、慢启动算法
##### b、拥塞避免算法
#### H、TCP协议的建立
##### a、三次握手
#### I、TCP协议的释放
##### a、四次释放
- 2MSL
- 等待计时器
#### J、套接字与套接字编程
##### a、套接字
##### b、服务端编程
##### c、客户端编程
### (4)、应用层篇
#### A、DNS服务详解
##### a、DNS的功能
##### b、域名详解
##### c、域名服务器
#### B、DHCP协议详解
##### a、DHCP是什么
##### b、DHCP的功能
#### C、HTTP协议详解
##### a、HTTP是什么
##### b、Web服务器
##### c、HTTP请求方法
##### d、HTTP指定资源
##### e、HTTP请求报文
##### f、HTTP应答报文
##### g、HTTP应该状态码
#### D、HTTP工作的结构
##### a、Web缓存
##### b、Web代理
##### c、CDN
##### d、爬虫
#### E、HTTPS协议详解
##### a、加密模型
##### b、数字证书
##### c、SSL握手过程
## 四、编程实践
### (1)、计算机组成原理实践
#### A、缓存置换算法
##### a、双向链表
##### b、FIFO缓存置换算法
##### c、LFU缓存置换算法
##### d、LRU缓存置换算法
### (2)、操作系统实践
#### A、线程池/异步任务框架
##### a、线程安全队列Queue
##### b、任务处理线程池Pool
##### c、基本本人对象Task
### (3)、计算机网络实践
#### A、网络嗅探工具
##### a、IP报文解析
##### b、UDP报文解析
##### c、TCP报文解析



## 五、杂项
### (1)、免密登录
- 编辑文件(项目目录/.git/config),修改修改remote中的url那行如下
`url = https://账号:密码@github.com/地址`
***
### (2)、ssh配置github
- 1、ssh-keygen -t rsa -C "761071654@qq.com"回车
- 2、等待出现Enter passphrase(empty for no passphrase):让你输入两次密码
- 3、去c:/Users/administratorn/.ssh/id_rsa.pub拷贝公钥内容到github-->Personal settings-->SSH keys的页面
- 4、:ssh -T git@github.com然后输入密码测试主机联通性
- 5、记忆法
 - 宫殿记忆
 - 思维导图