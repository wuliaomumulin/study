<?php
require_once 'node/Node.php';

/**
 * LFU最不常用算法
 * 1、淘汰魂村时，把使用频率最少的Node淘汰
 * 2、可能出现相同频率的一组Node节点，我们应该淘汰哪个呢？
 * 这时我们就用到了频率组的概念，就是相同频率的Node节点我们称之为一组频率，可以同时出现多组频率组成的频率地图，我们将它存在map元素中。
 * 3、同频率下我们按照FIFO算法淘汰
 */
class LFUNode extends Node{
	//记录频率，0
	public $freq = 0;
	public function __construct($key,$value){
		//parent是一个父类的指针，我们一般用parent::__construct调用父类的构造参数
		parent::__construct($key,$value);
	}
}

class LFU extends DoubleLinkedList{
	public $capacity;
	public $size = 0;
	public $map = [];
	//保存每一个频率,key:频率 value:频率对应的双向链表
	public $freq_map = [];

	public function __construct($capacity){
		$this->capacity = $capacity;
	}

	//更新节点频率的操作
	private function __update_freq($node){
		$freq = $node->freq;
		//删除
		$node = $this->freq_map[$freq]->remove($node);
		if($this->freq_map[$freq]->size == 0) unset($this->freq_map[$freq]);

		//更新
		$freq += 1;
		$node->freq = $freq;

		if(!array_key_exists($freq,$this->freq_map)){
			$this->freq_map[$freq] = new DoubleLinkedList();
		}
		$this->freq_map[$freq]->append($node);
	}

	public function get($key){
		if(!array_key_exists($key,$this->map)) return -1;
		$node = $this->map[$key];
		self::__update_freq($node);
		return $node->value;
	}

	public function put($key,$value){
		if($this->capacity == 0) return;
		if(array_key_exists($key,$this->map)){
			$node = $this->map[$key];
			$node->value = $value;
			self::__update_freq($node);
		}else{
			if($this->capacity == $this->size){
				$min_freq = min(array_keys($this->freq_map));
				$node = $this->freq_map[$min_freq]->pop();
				unset($this->map[$node->key]);
				$this->size -= 1;
			}
			$node = new LFUNode($key,$value);
			$node->freq = 1;
			$this->map[$key] = $node;
			if(!array_key_exists($node->freq,$this->freq_map)){
				$this->freq_map[$node->freq] = new DoubleLinkedList();
			}
			$node = $this->freq_map[$node->freq]->append($node);
			$this->size += 1;
		}
	}

	//打印测试
	public function console(){
		print('****************'.PHP_EOL);
		foreach ($this->freq_map as $k => $v) {
			printf("freq = %d",$k);
			$this->freq_map[$k]->console();
		}
		print('****************'.PHP_EOL);
	}
}

//测试逻辑
$cache = new LFU(3);
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
print($cache->get(3));
$cache->console();
$cache->put(4,4);
$cache->console();
print($cache->get(1));
$cache->console();
print($cache->get(3));
$cache->console();
print($cache->get(4));
$cache->console();
?>