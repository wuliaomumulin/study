# -*- encoding=utf-8

import datetime,socket

def client(i):
	# 创建套接字
	s = socket.socket()
	host = "127.0.0.1"
	port =6666
	# 连接套接字
	s.connect((host,port))
	
	# 接受请求
	print('接受信息:%s,客户端:%d' %(s.recv(1024),i))
		
	# 关闭连接
	s.close()


if __name__ == '__main__':
	for i in range(10):
		client(i)