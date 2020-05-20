<?php
//1、详细约束路由(controller/action)
$ws = new \Swoole\WebSocket\Server('19.19.19.11',9502);

$ws->on('open',function($ws,$request){
	var_dump($request->fd,$request->get,$request->server);
	$ws->push($request->fd,"你好,{$request->fd}");
});
$ws->on('message',function($ws,$frame){
	var_dump($frame);
	echo "消息:{$frame->data}\n";
	$ws->push($frame->fd,"服务:{$frame->data}");
});
$ws->on('close',function($ws,$fd){
	echo "客户端-{$fd}已关闭\n";
});

$ws->start();
?>