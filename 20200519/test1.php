<?php
	$serv = new Swoole\Server('0.0.0.0',9502,SWOOLE_PROCESS,SWOOLE_SOCK_UDP);

	$serv->on('Packet',function($serv,$data,$clientInfo){
		$serv->sendto($clientInfo['address'],$clientInfo['port'],"服务 ".$data);
		var_dump($clientInfo);
	});

	$serv->start();
?>