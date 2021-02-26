<?php
namespace App\WebSocket;

use EasySwoole\Socket\AbstractInterface\ParserInterface;
use EasySwoole\Socket\Client\WebSocket;
use EasySwoole\Socket\Bean\Caller;
use EasySwoole\Socket\Bean\Response;

/**
 * 此类是WebSocket中一些非强制的自定义事件处理
 */
class WebSocketEvent
{
	/**
	 * 握手事件
	 */
	public function onHandShake(\swoole_http_request $request,\swoole_http_response $response)
	{
		/** 此处自定义握手事件，返回false时终止握手 */
		if(!$this->customHandShake($request,$response)){
			$response->end();
			return false;
		};

		/** 此处是RFC规范中的WebSocket握手验证过程，必须执行，否则无法正确握手 */
		if($this->secWebsocketAccept($request,$response)){
			$response->end();
			return true;
		};

		$response->end();
		return false;
	}

	/**
	 * 自定义握手事件
	 */
	public function customHandShake(\swoole_http_request $request,\swoole_http_response $response): bool
	{
		/**
		 * 这里可以通过http request获取到相应数据，进行自定义验证即可
		 */
		$headers = $request->header;
		$cookie = $request->cookie;

		/**
		 * 如果不满足某些条件，则返回false
		 * if(){
		 * return false;
		 * }
		 */
		//var_dump($headers);


		return true;
	}

	 /**
     * RFC规范中的WebSocket握手验证过程
     * 以下内容必须强制使用
     *
     * @param \swoole_http_request  $request
     * @param \swoole_http_response $response
     * @return bool
     */
    protected function secWebsocketAccept(\swoole_http_request $request, \swoole_http_response $response): bool
    {
        // ws rfc 规范中约定的验证过程
        if (!isset($request->header['sec-websocket-key'])) {
            // 需要 Sec-WebSocket-Key 如果没有拒绝握手
            var_dump('shake fai1 3');
            return false;
        }
        if (0 === preg_match('#^[+/0-9A-Za-z]{21}[AQgw]==$#', $request->header['sec-websocket-key'])
            || 16 !== strlen(base64_decode($request->header['sec-websocket-key']))
        ) {
            //不接受握手
            var_dump('shake fai1 4');
            return false;
        }

        $key = base64_encode(sha1($request->header['sec-websocket-key'] . '258EAFA5-E914-47DA-95CA-C5AB0DC85B11', true));
        $headers = array(
            'Upgrade'               => 'websocket',
            'Connection'            => 'Upgrade',
            'Sec-WebSocket-Accept'  => $key,
            'Sec-WebSocket-Version' => '13',
            'KeepAlive'             => 'off',
        );

        if (isset($request->header['sec-websocket-protocol'])) {
            $headers['Sec-WebSocket-Protocol'] = $request->header['sec-websocket-protocol'];
        }

        // 发送验证后的header
        foreach ($headers as $key => $val) {
            $response->header($key, $val);
        }

        // 接受握手 还需要101状态码以切换状态
        $response->status(101);
        var_dump('shake success at fd :' . $request->fd);
        return true;
    }

    /**
    * 关闭事件
    */
    public function onClose(\swoole_server $server,int $fd,int $reactorId){
    	/**
    	 * @var array $info
    	 */
    	$info = $server->getClientInfo($fd);


    	if($info && $info['websocket_status'] === WEBSOCKET_STATUS_FRAME)
    	{
    		if($reactorId < 0){
    			echo "服务关闭\n";
    		}
    	}
    }
}
?>