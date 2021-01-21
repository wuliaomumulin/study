## win10代理，需要用管理员账号运行cmd
### 输入映射命令
```
netsh interface portproxy add v4tov4 listenport=10001 listenaddress=0.0.0.0 connectport=50070 connectaddress=192.168.163.143
```
### 查看所有监听的映射命令
```
netsh interface portproxy show all
```
### 备注：移除所有映射规则
```
netsh interface portproxy reset
```
## 备注：根据IP和端口移除映射规则
```
netsh interface portproxy delete v4tov4 listenaddress=0.0.0.0 listenport=10001
```

### Question:解决Nginx报错The plain HTTP request was sent to HTTPS port
#### 打开配置文件，查看HTTPS server段的配置：
修改前：
```
server {
        listen       443 ssl;
        server_name  localhost;
        ...
}
```
修改方式，将监听端口后的“ssl”删除，即：
```
server {
        listen       443;
        server_name  localhost;
        ...
}
```
#### 解决办法也很简单，就将配置文中的“ ssl on ; ” 注释掉或者修改成 “ ssl off ;”，这样，Nginx就可以同时处理HTTP请求和HTTPS请求了。
