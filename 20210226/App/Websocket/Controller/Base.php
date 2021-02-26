<?php
namespace App\Websocket\Controller;

use EasySwoole\Socket\AbstractInterface\Controller;
use EasySwoole\Socket\Client\WebSocket as WebSocketClient;
use EasySwoole\EasySwoole\ServerManager;
use EasySwoole\EasySwoole\Task\TaskManager;
use Exception;

class Base extends Controller
{
	//基础控制器
	protected function actionNotFound(?string $action){

		$this->response()->setMessage("您的{$action}请求不存在...");
		
		/**延迟关闭客户端fd*/
		$client = $this->caller()->getClient();
		TaskManager::getInstance()->async(function() use($client){
			//延时2s
			sleep(2);
			//关闭客户端
			ServerManager::getInstance()->getSwooleServer()->disconnect($client->getFd());
		});
		

	}

	protected function afterAction(?string $action){
		echo "请求之后执行\n";
	}

	/**
	* 心跳执行的方法
	*/
	public function heartbeat(){
		$fd = $this->caller()->getClient()->getFd();//用户请求的fd
		$data = $this->caller()->getArgs();//获取请求参数
		//var_dump($data,ServerManager::getInstance()->getSwooleServer()->worker_id);

		$this->response()->setMessage('PONG');

	}

}
?>