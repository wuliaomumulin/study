<?php
$http = new Swoole\Http\Server('19.19.19.11',9501);

$http->on('request',function($request,$response){
	var_dump($request->get,$request->post);
	$response->header("Content-Type","text/html; charset=utf-8");
	$response->end("<h1>Hello,#".rand(1000,9999)."</h1>");
});

$http->start();
?>