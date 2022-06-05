## docker使用基础(三)
> 在Window10系统下使用docker构建easyswoole3开发环境的补充

### 一、目录映射
```
docker run -it -p 9501:9501 easyswoole/easyswoole3
# upload
docker run -it -p 9501:9501 --network devel -v E:/www/db.easyswoole.cn/:/easyswoole/ --name test easyswoole/easyswoole3
# download
docker run --rm --volumes-from test -v E:/www/db.easyswoole.cn:/home easyswoole/easyswoole3 bash -c "tar cvf /home/easyswoole.tar /easyswoole/ "
## volume seting
docker run -it -p 9501:9501 --name test -v test:/easyswoole easyswoole/easyswoole3 
```
### 二、网络通讯
```
docker network create devel
docker network ls
docker run -d --name redis01 --network devel --network-alias redis01 redis
```

### 三、镜像组成
docker-compose.yml
```
version: "3"

services:
  app:
    build: ./
    ports:
      - 9501:9501
    volumes:
      - ./:/easyswoole
    environment:
      - TZ=Asia/Shanghai
    networks:
      - test
  redis01:
    image: redis:7.0.0
    ports:
      - 6379:6379
    volumes:
      - redis01:/data
    environment:
      - TZ=Asia/Shanghai
    networks:
      - test

volumes:
  redis01:

networks:
  test: {}
```

### 四、镜像构建
Dockerfile

```
FROM centos:8

#version defined
ENV SWOOLE_VERSION 4.4.26
ENV EASYSWOOLE_VERSION 3.4.x

#seting yum environment
Run cd /etc/yum.repos.d/ \
    && sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* \
    && sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-* \
    && yum makecache \
    && yum update -y

#install libs
RUN yum install -y curl zip unzip  wget openssl-devel gcc-c++ make autoconf git epel-release
RUN dnf -y install https://rpms.remirepo.net/enterprise/remi-release-8.rpm
#install php
RUN yum --enablerepo=remi install -y php74-php php74-php-devel php74-php-mbstring php74-php-json php74-php-simplexml php74-php-gd

RUN ln -s /opt/remi/php74/root/usr/bin/php /usr/bin/php \
    && ln -s /opt/remi/php74/root/usr/bin/phpize /usr/bin/phpize \
    && ln -s /opt/remi/php74/root/usr/bin/php-config /usr/bin/php-config

# composer
RUN curl -sS https://getcomposer.org/installer | php \
    && mv composer.phar /usr/bin/composer && chmod +x /usr/bin/composer
# use aliyun composer 由于最近阿里云镜像不稳定，废弃使用
# RUN composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/

# swoole ext
RUN wget https://github.com/swoole/swoole-src/archive/v${SWOOLE_VERSION}.tar.gz -O swoole.tar.gz \
    && mkdir -p swoole \
    && tar -xf swoole.tar.gz -C swoole --strip-components=1 \
    && rm swoole.tar.gz \
    && ( \
    cd swoole \
    && phpize \
    && ./configure --enable-openssl \
    && make \
    && make install \
    ) \
    && sed -i "2i extension=swoole.so" /etc/opt/remi/php74/php.ini \
    && rm -r swoole

# Dir
WORKDIR /easyswoole
# install easyswoole
RUN cd /easyswoole \
    && composer require easyswoole/easyswoole=${EASYSWOOLE_VERSION} \
    && php vendor/easyswoole/easyswoole/bin/easyswoole install \
    && php easyswoole install \
    && composer dump-autoload \
    && php easyswoole server start -d


EXPOSE 9501
```

#### 4.1 内容排除
.dockerignore

```
#Ignore the logs directory
Log/
Temp/

#Ignoring the password file
passwords.txt
*.txt

#Ignoring git and cache folders
.git
.cache
```
