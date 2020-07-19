### 一、annotation注解
### 二、aspect切入
### 三、aop切面
#### 1、协程间的数据混淆
> 通过代理类的方法去解决，在php中代理类可以是一个interface,如果requestInterface                                                                                                            绑定的实际是一个代理类，代理的所有方法都是从Context里面取出对应的对象进行操作。
协程上下文的切换 ConText:set|get

co
### 四、config

### 五、middleware
洋葱圈模型
php bin/hyperf.php gen:middleware BazMiddleware
### 六、consul
```
apt-get install consul
consul agent -dev -bind 19.19.19.11 -client 0.0.0.0 -ui
```
#### 2、发布对应的consul配置文件
php bin/hyperf.php vendor:publish hyperf/consul

### 七、异常处理
继承自\RuntimeException
### 八、事件机制

#### 1、事件解耦
 Event Listener EventDispatcher
#### 2、创建一个监听器
php bin/hyperf.php gen:listener SendSmsListener

#### 3、监听器的优先级可以通过在监听类的priority决定
```
**
 * @Listener(priority=9)
 *
```
### 九、命令行