> 前提生产环境是7.2

#### 扩展版本
```
swoole
Swoole => enabled
Author => Swoole Team <team@swoole.com>
Version => 4.4.16

redis
Redis Support => enabled
Redis Version => 5.2.0
Redis Sentinel Version => 0.1
Available serializers => php, json
```
#### 安装位置(php.ini)
- 举例
```
[redis]
extension="swoole.so"
extension = /www/server/php/72/lib/php/extensions/no-debug-non-zts-20170718/redis.so

; 关闭短标签
swoole.use_shortname = off

```