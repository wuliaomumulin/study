#! -*- encoding=utf-8 -*-
# FIFO先进先出页面缓存置换算法
# 1、当需要淘汰缓存时，把最先进入列表的字块淘汰
from node.Node import Node,DoubleLinkedList

class FIFO(object):
	def __init__(self,capacity):
		self.capacity = capacity
		self.size = 0
		self.map = {}
		self.list = DoubleLinkedList(self.capacity)

	def get(self,key):
		if key not in self.map:
			return -1
		else:
			return self.map.get(key).value 

	def put(self,key,value):
		if self.capacity == 0:
			return
		if key in self.map:
			node = self.map.get(key)
			self.list.remove(node)
			node.value = value
			self.list.append(node)
		else:
			if self.size == self.capacity:
				node = self.list.pop()
				del self.map[node.key]
				self.size -= 1
			node = Node(key,value)
			self.list.append(node)
			self.map[key] = node
			self.size += 1

	def print(self):
		self.list.print()


# 测试逻辑
if __name__ == '__main__':
	cache = FIFO(2)
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
	cache.put(4,4)
	cache.print()
	print(cache.get(1))
# exit()
# print(__name__)
# print(__name__)