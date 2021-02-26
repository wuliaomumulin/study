<?php
namespace App\WebSocket;

use EasySwoole\Socket\AbstractInterface\ParserInterface;
use EasySwoole\Socket\Client\WebSocket;
use EasySwoole\Socket\Bean\Caller;
use EasySwoole\Socket\Bean\Response;

class WebSocketParser implements ParserInterface
{
	/**
	 * decode 
	 * @param string $raw 客户端原始消息
	 * @param WebSocket $client Websocket client对象
	 * @return Caller Socket 调用对象
	 */
	public function decode($raw,$client): ?Caller
	{

		//new 调用者对象
		$caller = new Caller();

		//消息体:{"controller":"index","action":"index","params":{"content":"111"}}
		if($raw !== 'PING'){
			$payload = json_decode($raw,true);
			$class = isset($payload['controller']) ? $payload['controller'] :'index';
			$action = isset($payload['action']) ? $payload['action'] :'actionNotFound';
			$params = isset($payload['params']) ? (array)$payload['params'] : [];
			$controller = "\\App\\Websocket\\Controller\\".ucfirst($class);
			if(!class_exists($controller)) $controller = "\\App\\Websocket\\Controller\\Index";
			$caller->setClient($caller);
			$caller->setControllerClass($controller);
			$caller->setAction($action);
			$caller->setArgs($params);

		}else{
			//设置心跳执行类和方法
			$caller->setControllerClass(\App\Websocket\Controller\Base::class);
			$caller->setAction('heartbeat');
		}

		return $caller;
	}

	public function encode(Response $response,$client): ?string
	{
		return $response->getMessage();
	}
}
?>
