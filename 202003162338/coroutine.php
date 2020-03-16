<?php
require_once 'Task.php';
require_once 'Scheduler.php';

function task1(){
	$task_id = (yield getTaskId()); 
	for($i=0;$i<=300;$i++){
		//写入文件,大概要3000微妙
		usleep(3000);
		echo "写入文件{$i}\n";
		yield $i;
	}
}

function task2(){
	$task_id = (yield getTaskId());
	for($i=0;$i<=500;$i++){
		//给500个会员发送邮件,大概要3000微妙
		usleep(3000);
		echo "发送邮件{$i}\n";
		yield $i;
	}
}
function task3(){
	$task_id = (yield getTaskId());
	for($i=0;$i<=100;$i++){
		//模拟写入一百条数据,大概要3000微妙
		usleep(3000);
		echo "插入数据{$i}\n";
		yield $i;
	}
}
$scheduler = new Scheduler();
$scheduler->add(task1());
$scheduler->add(task2());
$scheduler->add(task3());

$scheduler->run();
?>