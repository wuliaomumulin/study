# -*- encoding=utf-8
import uuid
import threading

# 基本任务对象
class Task:

	# 任务可能是一些函数，并且带有参数
	def __init__(self,func,*args,**kwargs):
		# 任务的具体逻辑，通过函数引用传递进来
		self.callable = func
		self.id = uuid.uuid4()

		#
		self.args = args
		self.kwargs = kwargs

	def __str__(self):
		return '任务ID为' + str(self.id)


# 异步任务对象继承基本任务对象
class AsyncTask(Task):

	def __init__(self,func,*args,**kwargs):
		# 结果
		self.result = None
		# 条件变量
		self.condition = threading.Condition()
		super().__init__(func,*args,**kwargs)

	# 设置运行结果
	def set_result(self,result):
		self.condition.acquire()
		self.result = result
		# 通知可能等待的线程
		self.condition.notify()
		self.condition.release()

	# 获取任务结果
	def get_result(self):
		self.condition.acquire()
		# 如果没有任务，就需要等待
		if not self.result:
			self.condition.wait()
		result = self.result
		self.condition.release()
		return result

if __name__ == '__main__':
	def test_func():
		print('这是一个测试')
	
	task = Task(func=test_func)
	print(task)