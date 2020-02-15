<?php 
require_once 'node/Node.php';


# FIFO先进先出页面缓存置换算法
# 1、当需要淘汰缓存时，把最先进入列表的字块淘汰
class FIFO extends DoubleLinkedList{
	public $capacity;
	public $size = 0;
	public $map = [];
	public $list = [];

	public function __construct($capacity){
		$this->capacity = $capacity;
		$this->list = new DoubleLinkedList($capacity);
	}

	public function get($key){
		if(!array_key_exists($key,$this->map)) return -1;
		return $this->map[$key]->value;
	}
	public function put($key,$value){
		if($this->capacity == 0) return;
		if(array_key_exists($key,$this->map)){
			$node = $this->map[$key];
			$this->list->remove($node);
			$node->value = $value;
			$this->list->append($node);
		}else{
			if($this->size == $this->capacity){
				$node = $this->list->pop();
				unset($this->map[$node->key]);
				$this->size -= 1;	
			}
			$node = new Node($key,$value);
			$this->list->append($node);
			$this->map[$key] = $node;
			$this->size += 1;
		}
	}
	public function console(){
		$this->list->console();
	}
}


//逻辑测试
$cache = new FIFO(2);
$cache->put(1,1);
$cache->console();
$cache->put(2,2);
$cache->console();
print($cache->get(1));
$cache->console();
$cache->put(3,3);
$cache->console();
print($cache->get(2));
$cache->console();
$cache->put(4,4);
$cache->console();
print($cache->get(1));
$cache->console();
print($cache->get(4));
$cache->console();

?>