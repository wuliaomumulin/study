<?php
//一、协程通信
Swoole\Runtime::enableCoroutine();

$chan = new Chan(2);

//chan1
go(function() use($chan){
	$result = [];
	for($i=0;$i<2;$i++){
		$result += $chan->pop();
	}
	var_dump($result);
});
//chan2
go(function() use($chan){
	$cli = new Swoole\Coroutine\Http\Client('19.19.19.11',80);
	$cli->set(['timeout'=>10]);
	$cli->setHeaders([
		'host' => '19.19.19.11',
		'User-Agent' => 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36',
		'Accept' => 'application/json, text/plain, */*',
		'Accept-Encoding' => 'gzip, deflate',
		'Accept-Language' => 'zh-CN,zh;q=0.9,en;q=0.8',
	]);
	$ret = $cli->get('/');
	$chan->push(['19.19.19.11'=>$cli->statusCode]);
});
//chan3
go(function() use($chan){
	$cli = new Swoole\Coroutine\Http\Client('192.168.1.100',80);
	$cli->set(['timeout'=>10]);
	$cli->setHeaders([
		'host' => '192.168.1.100',
		'User-Agent' => 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36',
		'Accept' => 'application/json, text/plain, */*',
		'Accept-Encoding' => 'gzip, deflate',
		'Accept-Language' => 'zh-CN,zh;q=0.9,en;q=0.8',		
	]);
	$ret = $cli->get('/');
	$chan->push(['192.168.1.100'=>$cli->statusCode]);
});		
?>