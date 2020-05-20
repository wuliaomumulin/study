<?php
header('content-type:text/html;charset=utf-8');

Swoole\Runtime::enableCoroutine();

class WaitGroup{
	private $count = 0;
	private $chan;

	/**
	* waitgroup constructor.
	* 初始化一个chan
	*/
	public function __construct(){
		$this->chan = new Chan;
	}
	//增加计数
	public function add(){
		$this->count++;
	}
	//任务已完成
	public function done(){
		$this->chan->push(true);
	}
	//等待所有任务完成恢复当前协程的执行
	public function wait(){
		while($this->count--){
			$this->chan->pop();
		}
	}
}

/**
* example
*/
go(function(){
	$wg = new WaitGroup();
	$result = [];

	//第一个chan
	$wg->add();
	go(function() use($wg,&$result){
		$cli = new Swoole\Coroutine\Http\Client('19.19.19.11',443);
		$cli->set(['timeout'=>1]);
		$cli->setHeaders([
			'host' => '19.19.19.11',
			'User-Agent' => 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36',
			'Accept' => 'application/json, text/plain, */*',
			'Accept-Encoding' => 'gzip, deflate',
			'Accept-Language' => 'zh-CN,zh;q=0.9,en;q=0.8',
		]);
		$ret = $cli->get('/');
		var_dump($cli);
		$result['19.19.19.11'] = $cli->statusCode;
		$cli->close();
		$wg->done();
	});
	//第二个chan
	$wg->add();
	go(function() use($wg,&$result){
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
		$result['192.168.1.100'] = $cli->statusCode;
		$cli->close();
		$wg->done();
	});	
	//挂起当前协程，等待所有任务完成之后结束
	$wg->wait();
	var_dump($result);
});

