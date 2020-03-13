# -*- encoding=utf-8

import threading
# pip install psutil
import psutil
from operate_system.task import Task,AsyncTask
from operate_system.queue import ThreadSafeQueue

# 任务处理线程的一些功能
# 继承threading扩展库
class ProcessThread(threading.Thread):

	def __init__(self,task_queue,*args,**kwargs):
		threading.Thread.__init__(self,*args,**kwargs)
		# 任务线程停止的标记
		self.dismiss_flag = threading.Event()
		# 任务队列(处理线程不断从队列取出元素处理)
		self.task_queue = task_queue
		self.args = args
		self.kwargs = kwargs

	def run(self):
		while True:
			# 如果线程被标记停止，就跳出
			# 判断event的标志是否为True
			if self.dismiss_flag.is_set():
				break

			task = self.task_queue.pop()
			if not isinstance(task,Task):
				continue
			# 执行task实际逻辑(是通过函数调用引进来的)
			result = task.callable(*task.args,**task.kwargs)
			# 如果是异步处理任务
			if isinstance(task,AsyncTask):
				task.set_result(result)

	
	def dismiss(self):
		# 将event的标志设置为True，调用wait方法的所有线程将被唤醒；
		self.dismiss_flag.set()

	def stop(self):
		self.dismiss()

# 线程池	
class ThreadPool(object):
	def __init__(self,size=0):
		if not size:
			# 约定线程数量一般定义为Cpu核数的两倍(最大实践)
			size = psutil.cpu_count()*2
		# 线程池
		self.pool = ThreadSafeQueue(size)
		# 任务队列
		self.task_queue = ThreadSafeQueue() 

		for i in range(size):
			self.pool.put(ProcessThread(self.task_queue))

	# 启动线程池
	def start(self):
		for i in range(self.pool.size()):
			thread = self.pool.get(i)
			thread.start()

	# 停止线程池
	def join(self):
		# range(start,stop)创建一个证书列表,[0,1,2]
		for i in range(self.pool.size()):
			thread = self.pool.get(i)
			thread.stop()

		# 清空线程池
		while self.pool.size():
			thread = self.pool.pop()
			# join() 方法的功能是在程序指定位置，优先让该方法的调用者使用 CPU 资源
			thread.join()
	# 提交
	def put(self,item):
		if not isinstance(item,Task):
			# 使用raise语句自己触发异常
			raise TaskTypeErrorException()
		self.task_queue.put(item)

	# 批量提交
	def batch_put(self,item_list):
		# 如果不是列表，将数据强制转为列表
		if not isinstance(item_list,list):
			item_list = list(item_list)
		for item in item_list:
			self.put(item)
	# 数量
	def size(self):
		return self.pool.size()

# 线程池标准错误
class TaskTypeErrorException(Exception):
	pass