### 一、免密登录
- 编辑文件(项目目录/.git/config),修改修改remote中的url那行如下
`url = https://账号:密码@github.com/地址`
***
### 二、ssh配置github
- 1、ssh-keygen -t rsa -C "761071654@qq.com"回车
- 2、等待出现Enter passphrase(empty for no passphrase):让你输入两次密码
- 3、去c:/Users/administratorn/.ssh/id_rsa.pub拷贝公钥内容到github-->Personal settings-->SSH keys的页面
- 4、:ssh -T git@github.com然后输入密码测试主机联通性