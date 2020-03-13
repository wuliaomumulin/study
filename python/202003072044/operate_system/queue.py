# -*- encoding=utf-8

import threading
import time
import random

class ThreadSafeQueueException(Exception):
	pass

# 线程安全的队列
class ThreadSafeQueue(object):
	
	def __init__(self,max_size=0):
		self.queue = []
		self.max_size = max_size
		self.lock = threading.Lock()
		self.condition = threading.Condition()

	# 当前队列元素的数量
	def size(self):
		self.lock.acquire()
		size = len(self.queue)
		self.lock.release()
		return size

	# 向队列放入元素
	def put(self,item):
	 	if self.size != 0 and self.size() > self.max_size:
	 		return ThreadSafeQueueException()
	 	self.lock.acquire()
	 	self.queue.append(item)
	 	self.lock.release()

	 	self.condition.acquire()
	 	self.condition.notify()
	 	self.condition.release()

	# 向队列放入多个元素
	def batch_put(self,item_list):
		# 判断变量是不是个链表
		if not isinstance(item_list,list):
			item_list = list(item_list)
		for item in item_list:
			self.put(item)

	# 从队列头部取出元素
	# block 否，是否阻塞
	# timeout 最大等待时间
	def pop(self,block=False,timeout=0):
		if self.size() == 0:
			# 如果需要阻塞等待
			if block:
				self.condition.acquire()
				self.condition.wait(timeout=timeout)
				self.condition.release()
			else:
				return None

		# 先加锁
		self.lock.acquire()

		item = None

		# 即使阻塞等待了timeout时间段还是没有数据的判断
		if len(self.queue) > 0:
			item = self.queue.pop()
		self.lock.release()
		return item

	# 从队列中取出元素
	def get(self,index):
 		self.lock.acquire()
	 	item = self.queue[index]
	 	self.lock.release()
	 	return item

if __name__ == "__main__":
	queue = ThreadSafeQueue(max_size=10)
	
	def producer():
		while True:
			queue.put(random.randint(0,100))
			# time.sleep(1)
			# 将生产时间放久一点
			time.sleep(3)
	
	def consumer():
		while True:
			# 打印一些输出
			# item = queue.pop()
			item = queue.pop(block=True,timeout=2)
			print('获得队列项目为:%s' % item)
			time.sleep(1)

	thread1 = threading.Thread(target=producer)
	thread2 = threading.Thread(target=consumer)
	thread1.start()
	thread2.start()
	thread1.join()
	thread2.join()
