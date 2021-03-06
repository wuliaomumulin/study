## 应用层
- 传输层以及以下的层提供完整的信息服务
- 应用层是面向用户的一层
 - 已有的应用层软件
 - 面向传输层的编程

|FTP|HTTP|HTTPS|DNS|TELNET|
|--|--|--|--|--|
|21|80|443|53|23|

> UDP:提供多媒体分发,如视频、语音、实时信息;
TCP:可靠性传输,如金融交易、可靠通讯、MQ(消息队列);

- 定义应用间通信的规则
 - 应用进程的报文类型(请求报文、应答报文);
 - 报文的语法和格式;
 - 应用进程发送数据的时机、规则;
### 一、DNS详解
#### (1)、DNS(Domain Name System:域名系统)
- 意义:
 - 使用域名帮助记忆
- 域名由点、字母和数字组成
- 点分割不同的域
- 域名可以分为顶级域、二级域、三级域
#### (2)、域名详解
#### i.顶级域
- 国家
 - cn(中国)
 - uk(英国)
 - us(美国)
 - ca(加拿大)
- 通用
 - com(公司机构)
 - net(网络机构)
 - gov(政府机构)
 - org(组织机构)
##### ii.二级域
- aliyun
- amazon
- taobao
- qq
- baidu
#### (4)、域名服务器
***
### 二、DHCP详解
> DHCP(Dynamic Host Configuration Protocol:动态主机设置协议)
#### (1)、DHCP是一个局域网协议
- 即插即用联网
- 临时IP
- 租期
#### (2)、DHCP是应用UDP协议的应用层协议
 > DHCP服务器默认监听端口:67

- 主机使用UDP协议广播DHCP发现报文
- DHCP服务器发出DHCP提供报文
- 主机向DHCP服务器发出DHCP请求报文
- DHCP服务器回应并提供IP地址
***
### 三、HTTP协议详解
#### (1)、HTTP(HyperText Transfer Protocol:超文本传输协议)
> 超文本:超级文本，带超链接文本，就是一个富文本，可以有图片、文本、文件、动图、音频、视频等。
http(s)://<主机>:<端口>/<路径>

#### (2)、HTTP协议是可靠的数据传输协议
- client-server工作模式
- 硬件部分和软件部分
##### i.web服务器的工作过程
- 接受客户端连接
- 接受请求报文
- 处理请求
- 范围Web资源
- 构造应答
- 发送应答
#### (3)、HTTP的请求方法和指定资源
##### i.HTTP请求动作
|HTTP请求方式|含义|
|--|--|
|GET|获取指定的服务端资源|
|POST|提交数据到服务端|
|DELETE|删除指定的服务端资源|
|UPDATE|更新指定的服务端资源|
|PUT|--|
|OPTIONS|--|
|PATCH|--|
|HEAD|--|
|TRACE|--|

##### ii.在地址中指定和在请求数据中指定
- 请求报文
<table style="text-align:center;">
	<tr><td>请求方法</td><td>请求地址</td><td>HTTP版本</td></tr>
	<tr><td colspan="3">请求头</td></tr>
	<tr><td colspan="3">请求内容</td></tr>
</table>
- 请求报文example:
		POST https://cn.bing.com/
		Accept-Encoding:gzip
		Accept-Language:zh-CN
		{
			"sort":0,
			"unlearn":0,
			"page":2,
		}

- 应答报文
<table style="text-align:center;">
	<tr><td>HTTP版本</td><td>状态码</td><td>状态解释</td></tr>
	<tr><td colspan="3">应答头</td></tr>
	<tr><td colspan="3">应答内容</td></tr>
</table>
- 应答报文状态码:
|状态码|含义|
|--|--|
|200~299|成功状态码|
|300~399|重定向状态码|
|400~499|客户端错误状态码|
|500~599|服务端错误状态码|

***
### 四、HTTP工作的结构
#### (1)、Web缓存
- 20%/80%(二八原则)
- 存储器层次结构(缓存量小，主存中等，辅存最大)
#### (2)、Web代理

图1

- i.正向代理

图2

- ii.反向代理

图3

- iii.成功案例
 - Nginx
 - HAPorxy
***
#### (3)、CDN
> CDN(Content Delivery Network:内容分发网络)
基本上是承载多媒体内容为主，内容分发节点分布在各个区域。
***
#### (3)爬虫
- 增加网络拥塞
- 损耗服务器资源
***
### 五、HTTPS协议详解
> HTTP是明文传输的
账号密码-->个人信息,如账户金额、交易信息、敏感信息

#### (1)、HTTPS(Secure)是安全的HTTP协议
> http(s)://<主机>:<端口>/<路径>
默认端口443

#### (2)、加密模型
##### i.对称加密
> 使用密钥对数据加密之后的数据，再将它使用密钥解密，得到的数据应该没有改变。
##### ii.非对称加密
> 两个密钥不一致
- A密钥、B密钥是拥有一定数学关系的一组密钥
- 公钥:公钥给大家使用，对外公开
- 私钥:私钥自己使用，不对外公开
- 数字证书是可信任组织颁发给特定对象的认证
		证书格式、版本号
		整数序列号
		签名算法
		有效期
		对象名称
		对象公开密钥
		...

- SSL(Secure Sockets Layer:安全套接层)
 - 位于传输层和应用层之间
 - 提供数据安全和数据完整服务
 - 对传输层数据进行加密后传输
- 
#### (3)、HTTPS流程
图4
##### i.SSL安全参数握手
图5

图6
- 综合使用对称加密、非对称加密;
- 双方分别生成密钥，没有经过传输;
***
