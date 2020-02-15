<?php

class Node{
	protected $key = NULL;//键
	protected $value = NULL;//值
	protected $prev = NULL;//前驱
	protected $next = NULL;//后驱
	public function __construct($key,$value){
		$this->key = $key;
		$this->value = $value;	
	}
	public function __toString(){
		return sprintf('%s => %s',$this->key,$this->value);
	}
	/*
	php获取内存地址：
	因为传统PHP是单线程的，从设计思想上来说是不支持获取内存地址的，这样可以使程序员更注重于业务开发，而且获取到地址用处也有限，如果非要获取内存地址可以使用&符号代替。
	 */
}
class DoubleLinkedList extends Node{
	/*一个简单的双向列表*/
	private $head;
	private $tail;
	protected $size = 0;
	private $capacity = 0xffff;

	public function __construct($capacity=0xffff){
		$this->capacity = $capacity;
		$this->head = $this->tail = new Node(NULL,NULL);
	}

	//头部添加
	private function __add_head($node){
		
		if(self::is_empty($this->head->key)){
			$this->head = $node;
			$this->tail = $node;
			$this->head->prev = NULL;
			$this->head->next = NULL;
		}else{
			$node->next = $this->head;
			$this->head->prev = $node;
			$this->head = $node;
			$this->head->prev = NULL;
		}
		$this->size += 1;
		return $node;
	}

	//尾部追加
	private function __add_tail($node){
		if(@self::is_empty($this->tail->key)){
			$this->tail = $node;
			$this->head = $node;
			$this->tail->next = NULL;
			$this->tail->prev = NULL;
		}else{
			$this->tail->next = $node;
			$node->prev = $this->tail;
			$this->tail = $node;
			$this->tail->next = NULL;
		}
		$this->size += 1;
		return $node;
	}

	//尾部删除
	private function __del_tail(){
		if(self::is_empty($this->tail->key)) return;
		$node = $this->tail;
		if(!self::is_empty($node->prev)){
			$this->tail = $node->prev;
			$this->tail->next = NULL;
		}else{
			$this->tail = $this->head = NULL;
		}
		$this->size -= 1;
		return $node;
	}

	//头部删除
	private function __del_head(){
		if(self::is_empty($this->head->key)) return;
		$node = $this->head;
		if(@!self::is_empty($node->next->key)){
			$this->head = $node->next;
			$this->head->prev = NULL;
		}else{
			$this->tail = $this->head = NULL;
		}
		$this->size -= 1;
		return $node;
	}

	//删除任意节点
	private function __remove($node){
		//如果node==NULL删除尾部节点
		if($this->is_empty($node)){
			$node = $this->tail;
		}

		if($node == $this->tail){
			$this->__del_tail();
		}elseif($node == $this->head){
			$this->__del_head();
		}else{
			$node->prev->next = $node->next;
			$node->next->prev = $node->prev;
			$this->size -= 1;
		}
		return $node;
	}

	//弹出头部节点
	public function pop(){
		return $this->__del_head();
	}

	//追加节点
	public function append($node){
		return $this->__add_tail($node);
	}

	//头部添加节点
	public function append_front($node){
		return $this->__add_head($node);
	}
	
	//删除节点
	public function remove($node=NULL){
		return $this->__remove($node);
	}
	/*
	 判断当前节点是否为空,为空返回TRUE,不为空返回FALSE
	*/
	public static function is_empty($key){
		return (isset($key) and !is_null($key) and !empty($key)) ? FALSE : TRUE;
	}

	//输出
	public function console(){
		$p = $this->head;		
		$line = '';
		while(!self::is_empty($p)){
			$line .= sprintf("{%s: %s}",$p->key,$p->value);
			$p = $p->next;
			if(!self::is_empty($p)){
				$line .= '=>';
			}
		}
		print($line.PHP_EOL);
	}
}
//测试双向链表
/*$l = new DoubleLinkedList(10);
$nodes = [];
for($i=1;$i<=10;$i++){
 	array_push($nodes,new Node($i,$i));
}
$l->append($nodes[0]);
$l->console();
$l->append($nodes[1]);
$l->console();
$l->pop();
$l->console();
$l->append($nodes[2]);
$l->console();
$l->append_front($nodes[3]);
$l->console();
$l->append($nodes[4]);
$l->console();
$l->remove($nodes[2]);
$l->console();
$l->remove();
$l->console();
*/