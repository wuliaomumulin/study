#! -*- encoding=utf-8 -*-

# LFU最不常使用算法
# 1、淘汰缓存时，把使用频率最少的淘汰
# 2、可能存在相同频率的情况，这时应该淘汰哪个节点呢?
# 3、同频率下按照FIFO算法淘汰
# 
from node.Node import Node,DoubleLinkedList

class LFUNode(Node):
	def __init__(self,key,value):
		# 记录频率
		self.freq = 0
		# super是一个调用父类的方法，下句为调用父类的构造方法
		super(LFUNode,self).__init__(key,value)

class LFU(object):
	def __init__(self,capacity):
		self.capacity = capacity
		self.size = 0
		self.map = {}
		# 保存每一个频率,key:频率 value:频率对应的双向链表
		self.freq_map = {}

	# 更新节点频率的操作
	def __update_freq(self,node):
		freq = node.freq
		# 删除
		node = self.freq_map[freq].remove(node)
		if self.freq_map[freq].size == 0:
			del self.freq_map[freq]

		# # 更新
		freq += 1
		node.freq = freq

		if freq not in self.freq_map:
			self.freq_map[freq]= DoubleLinkedList()
		self.freq_map[freq].append(node)

	def get(self,key):
		if key not in self.map:
			return -1
		node = self.map.get(key)
		# print('key:',node)
		self.__update_freq(node)
		return node.value


	def put(self,key,value):
		if self.capacity == 0:
			return

		if key in self.map:
			node = self.map.get(key)
			node.value = value
			self.__update_freq(node)

		else:
			if self.capacity == self.size:
				min_freq = min(self.freq_map)
				node = self.freq_map[min_freq].pop()
				del self.map[node.key]
				self.size -= 1
			node = LFUNode(key,value)
			node.freq = 1
			self.map[key] = node
			if node.freq not in self.freq_map:
				self.freq_map[node.freq] = DoubleLinkedList()
			node = self.freq_map[node.freq].append(node)
			self.size += 1
	
	def print(self):
		print('*********************')
		for k,v in self.freq_map.items():
			print('freq = %d' % k)
			self.freq_map[k].print()
		print('*********************')
		print()

# 测试逻辑
if __name__ == '__main__':
	cache = LFU(3)
	cache.put(1,1)
	cache.print()
	cache.put(2,2)
	cache.print()
	print(cache.get(1))
	cache.print()
	cache.put(3,3)
	cache.print()
	print(cache.get(2))
	cache.print()
	print(cache.get(3))
	cache.print()
	cache.put(4,4)
	cache.print()
	print(cache.get(1))
	cache.print()
	print(cache.get(3))
	cache.print()
	print(cache.get(4))
	cache.print()
