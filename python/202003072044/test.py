# -*- encoding=utf-8

from operate_system import task,pool 
import time,random


class SimpleTask(task.Task):
	def __init__(self,callable):
		# 调用父类构造参数
		super(SimpleTask,self).__init__(callable)

def process():
	print('这是标号为%d调用方法.' % random.randint(0,100))
	time.sleep(1)

# 任务投递
def test():
	# 1、初始化一个线程池
	test_pool = pool.ThreadPool()
	test_pool.start()
	# 2、生成一系列的任务
	for i in range(10):
		simple_task = SimpleTask(process)
		# 3、往线程池提交任务执行
		test_pool.put(simple_task)

# 异步任务投递
def test_async_task():

	# 内部函数不允许外部调用
	def async_process():
		num = 1
		for i in range(100):
			num += i
		return num

	# 初始化一个线程池
	test_pool = pool.ThreadPool()
	test_pool.start()
	# 生成一系列的异步任务
	for i in range(10):
		simple_task = task.AsyncTask(func=async_process)
		# 3、往线程池提交任务执行
		test_pool.put(simple_task)
		# 4、获取异步任务执行的结果
		result = simple_task.get_result()
		# 5、显示
		print('异步任务处理的结果为:%d' % result)		

# 测试等待是否真正执行
def test_async_task2():
 	# 内部函数不允许外部调用
	def async_process():
		num = 1
		for i in range(100):
			num += i

		# 模拟阻塞
		time.sleep(10)
		
		return num

	# 初始化一个线程池
	test_pool = pool.ThreadPool()
	test_pool.start()
	# 生成一系列的异步任务
	for i in range(1):
		simple_task = task.AsyncTask(func=async_process)
		# 3、往线程池提交任务执行
		test_pool.put(simple_task)
		# 打印调用之前的时间
		print('投递之前的时间:%d' % time.time())
		# 打印调用之后的时间
		# 4、获取异步任务执行的结果
		result = simple_task.get_result()
		# 5、显示
		print('异步任务处理之后的时间:%d,处理的结果为:%d' % (time.time(),result))


# 测试没有等待是否可以正常获取结果
def test_async_task3():
 	# 内部函数不允许外部调用
	def async_process():
		num = 1
		for i in range(100):
			num += i		
		return num

	# 初始化一个线程池
	test_pool = pool.ThreadPool()
	test_pool.start()
	# 生成一系列的异步任务
	for i in range(1):
		simple_task = task.AsyncTask(func=async_process)
		# 3、往线程池提交任务执行
		test_pool.put(simple_task)
		# 打印调用之前的时间
		print('投递之前的时间:%d' % time.time())
		# 非任务处理所用时间,而是在获取结果之前新增的一个时间片
		# 这里可以不是sleep，可以转而处理别的逻辑
		time.sleep(5)

		# 打印调用之后的时间
		# 4、获取异步任务执行的结果
		result = simple_task.get_result()
		# 5、显示
		print('异步任务处理之后的时间:%d,处理的结果为:%d' % (time.time(),result))



if __name__ == '__main__':
	test()
	test_async_task()
	test_async_task2()
	test_async_task3()