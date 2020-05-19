<?php
$serv = new Swoole\Server('19.19.19.11',9501);
$serv->set(
	array('task_worker_num'=>4)
);

//worker进程中执行
$serv->on('receive',function($serv,$fd,$from_id,$data){
	$task_id = $serv->task($data);
	echo "分配异步任务:id:{$task_id}\n";
});
//task进程中执行
$serv->on('task',function($serv,$task_id,$from_id,$data){
	echo "新的异步任务，id:{$task_id}\n";
	$serv->finish("{$data} -> OK");
});
//worker进程中执行
$serv->on('finish',function($serv,$task_id,$data){
	echo "异步任务{$task_id}完成:{$data}\n";
});


$serv->start();
?>