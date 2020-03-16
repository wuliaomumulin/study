<?php

class YieldCall{
	protected $callback;

	public function __construct(callable $callback){
		$this->callback = $callback;
	}
	
	public function __invoke(Task $task,Scheduler $scheduler){
		$callback = $this->callback;
		return $callback($task,$scheduler);
	}

	public function getTaskId(){
		//返回一个YieldCall实例
		return new YieldCall(
			//该匿名函数会先去获取任务ID,然后会send给生成器，并且由YieldCall将task_id返回给生成器函数
			function(Task $task,Scheduler $scheduler){
				$task->setSendValue($task->getTaskId());
				$scheduler->schedule($task);
			});
	}
}
/**
 * 任务调度
 */
class Scheduler{
	protected $maxTaskId = 0;//任务ID
	protected $taskMap = [];//taskId = task
	protected $taskQueue;//任务队列
	public function __construct(){
		//SqlQueue is a doubleLinkedList 
		$this->taskQueue = new SplQueue();
	}
	public function add(Generator $coroutine){
		$tid = ++$this->maxTaskId;
		//新增任务
		$task = new Task($tid,$coroutine);
		$this->taskMap[$tid] = $task;
		$this->schedule($task);
		return $tid;	
	}

	/*
	*  任务入列
	*  
	 */
	public function schedule($task){
		$this->taskQueue->enqueue($task);
	}

	public function run(){
		while(!$this->taskQueue->isEmpty()){
			//任务出列进行遍历生成器数据
			$task = $this->taskQueue->dequeue();
			$retval = $task->run();
			
			//如果返回的是YieldCall实例，则先执行
			if($retval instanceof YieldCall){
				$retval($task,$this);
				continue;
			}

			if($task->isFinished()){
				//完成该删除该任务
				unset($this->taskMap[$task->getTaskId()]);
			}else{
				//继续入列
				$this->schedule($task);
			}
		}
	}
}