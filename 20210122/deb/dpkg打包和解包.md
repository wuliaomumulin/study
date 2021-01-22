## dpkg打包和解包
> 该工具是一个linux和deb包内容的映射,在包里面的任何目录和内容都会拷贝到外部环境的相同目录之下。

### 一、文件校验
md5sums是以用来验证文件是否损坏的凭证，生成
```
md5sum `find phoronix-test-suite_9.6.1_all -type f` > DEBIAN/md5sums
# 或者
find phoronix-test-suite_9.6.1_all/ -type f -exec md5sum {} + > DEBIAN/md5sums
```

### 二、解压
```
# 解压一个包到指定目录
dpkg -X test.deb test1
# 解压一个包的控制文件到指定目录下
dpkg -e test.deb test1/DEBIAN/
```
DEBIAN下的control文件类似这样的格式
```
Package: test-sh
Version: 1.1.1
Section: Utilities
Installed-Size: 24212
Priority: optional
Architecture: all
Depends: php-cli|php5-cli,php5-cli|php-xml
Recommends: build-essential, php-gd|php5-gd
Maintainer: test benchwork <liyb@163.com>
Description:descr 
 liyb@163.com
```

### 三、构建
```
# 构建一个包
dpkg -b test/ test.deb
# 查看一个包体结构
dpkg -c test.deb
```
### 四、正常使用
```
# 安装
dpkg -i test.deb
# 列出包的版本、名称、类型、描述
dpkg -l test.deb
# 列出包安装之后的位置
dpkg -L test-sh
# 查看包的详细说明
dpkg -s test-sh
# 保留配置移除软件
dpkg -r test-sh
# 不保留配置移除软件
dpkg -P test-sh
```

### 五、其他
```
# php命令行运行某个站点
php -S 0.0.0.0:1234 -t /usr/share/phoronix-test-suite/pts-core/phoromatic/public_html/
```