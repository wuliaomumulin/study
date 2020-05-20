<?php
//延迟任务
Swoole\Runtime::enableCoroutine();

go(function(){
	echo "a";
	defer(function(){
		echo "~a";
	});
	echo "b";
	defer(function(){
		echo "~b";
	});
	sleep(1);
	echo 'c';
	//依次输出abc~b~a,综训先进后出原则
});
?>