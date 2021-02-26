<?php
namespace App\Websocket\Controller;

use EasySwoole\EasySwoole\Task\TaskManager;
use EasySwoole\EasySwoole\ServerManager;


class Index extends Base
{
	public function index(){
		$fd = $this->caller()->getClient()->getFd();//用户请求的fd
		$data = $this->caller()->getArgs();//获取请求参数
		var_dump($data);
		$this->response()->setMessage("您的用户ID:{$fd}");
		echo "接收到客户端的连接\n";
	}
	/**
	 * decode 
	 * @param string $raw 客户端原始消息
	 * @param WebSocket $client Websocket client对象
	 * @return Caller Socket 调用对象
	 */
	public function hello()
	{
		$this->response()->setMessage('调用hello方法携带参数为:'.json_encode($this->caller()->getArgs()));
	}

	public function who(){
		$this->response()->setMessage('你的fd为'.$this->caller()->getClient()->getFd());
	}
	/**
	* 异步任务推送，可以做列表
	*/
	public function delay()
	{
		$this->response()->setMessage('这是一个延迟动作');
		$client = $this->caller()->getClient();

		//异步推送，这里直接use fd也是可以的
		TaskManager::getInstance()->async(function() use($client){
			$server = ServerManager::getInstance()->getSwooleServer();
			$i = 0;
			while($i < 5){
				sleep(1);
				$server->push($client->getFd(),'在http请求中推送时间'.date('H:i:s'));
				$i++;
			}
		});
	}
}
?>