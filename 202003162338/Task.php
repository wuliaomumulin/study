<?php
	/**
	 * 任务类
	 */
	Class Task{
		protected $taskId;//任务ID
		protected $coroutine;//生成器
		protected $sendValue = null;//生成器send值
		protected $beforeFirstYield = True;//迭代生成器是否是第一个

		public function __construct($taskId,Generator $coroutine){
			$this->taskId = $taskId;
			$this->coroutine = $coroutine;
		}

		public function getTaskId(){
			return $this->taskId;
		}
		/**
		 * 设置插入数据
		 */
		public function setSendValue($sendValue){
			$this->sendValue = $sendValue;
		}
		/**
		 * 对数据进行迭代
		 */
		public function run(){
			//如果是
			if($this->beforeFirstYield){
				$this->beforeFirstYield = False;
				var_dump($this->coroutine->current());
				return $this->coroutine->current();
			}else{
				$retval = $this->coroutine->send($this->sendValue);
				$this->sendValue = null;
				return $retval;
			}
		}
		/**
		 * 是否完成
		 */
		public function isFinished(){
			return !$this->coroutine->valid();
		}
	}
?>