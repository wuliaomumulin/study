<?php 
require_once 'node/Node.php';


# LRU最近最少使用页面置换算法
# 剔除最少使用的字块，将最近使用的字块提前到链表头部
class LRU extends DoubleLinkedList{
	public $capacity;
	public $size = 0;
	public $map = [];
	public $list = [];

	public function __construct($capacity){
		$this->capacity = $capacity;
		$this->list = new DoubleLinkedList($this->capacity);
	}

	public function get($key){
		if(!array_key_exists($key,$this->map)) return -1;
		$node = $this->map[$key];
		$this->list->remove($node);
		$this->list->append_front($node);
		return $node->value;
	}

	public function put($key,$value){
		if(array_key_exists($key,$this->map)){
			$node = $this->map[$key];
			$this->list->remove($node);
			$node->value = $value;
			$this->list->append_front($node);
		}else{
			$node = new Node($key,$value);
			if($this->size >= $this->capacity){
				$old_node = $this->list->remove();
				unset($this->map[$old_node->key]);
				$this->size -= 1;
			}

			$this->list->append_front($node);
			$this->map[$key] = $node;
			$this->size += 1;
		}
	}
	public function console(){
		$this->list->console();
	}
}

//逻辑测试
$cache = new LRU(2);
$cache->put(2,2);
$cache->console();
$cache->put(1,1);
$cache->console();
$cache->put(3,3);
$cache->console();
print($cache->get(1));
$cache->console();
print($cache->get(2));
$cache->console();
print($cache->get(3));
$cache->console();