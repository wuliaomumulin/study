#! -*- encoding=utf-8 -*-

# LRU最近最少使用页面置换算法
# 剔除最少使用的字块，将最近使用的字块提前到链表头部
from node.Node import Node,DoubleLinkedList

class LRU(object):
	def __init__(self,capacity):
		self.capacity = capacity
		self.map = {}
		self.list = DoubleLinkedList(self.capacity)
		self.size = 0

	def get(self,key):
		if key in self.map:
			node = self.map[key]
			self.list.remove(node)
			self.list.append_front(node)
			return node.value 
		else:
			return -1 

	def put(self,key,value):
		if key in self.map:
			node = self.map.get(key)
			self.list.remove(node)
			node.value = value
			self.list.append_front(node)
			# pass
		else:
			node = Node(key,value)
			if self.size >= self.capacity:
				old_node = self.list.remove()
				self.map.pop(old_node.key)
				self.size -= 1
			self.list.append_front(node)
			self.map[key] = node
			self.size += 1
			# pass
	
	def print(self):
		self.list.print()

# 测试逻辑
if __name__ == '__main__':
	cache = LRU(2)
	cache.put(2,2)
	cache.print()
	cache.put(1,1)
	cache.print()
	cache.put(3,3)
	cache.print()
	print(cache.get(1))
	cache.print()
	print(cache.get(2))
	cache.print()
	print(cache.get(3))
	cache.print()