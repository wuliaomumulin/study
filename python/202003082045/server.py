# -*- encoding=utf-8

import datetime,socket

def server():
	# 创建套接字
	s = socket.socket()
	host = "127.0.0.1"
	port =6666
	# 绑定套接字
	s.bind((host,port))
	# 监听
	s.listen(5)
	
	while True:
		# 接受请求
		c,addr = s.accept()
		
		print('连接请求:',addr)
		# 发送信息
		c.send('欢迎登陆服务器,当前时间为:%s' % datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S'))

		# 关闭连接
		c.close()


if __name__ == '__main__':
	server()