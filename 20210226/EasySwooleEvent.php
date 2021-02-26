<?php
namespace EasySwoole\EasySwoole;

use EasySwoole\Socket\Dispatcher;
use App\Websocket\WebSocketParser;
use App\Websocket\WebSocketEvent;
use App\Exception\ExceptionHeadler;

use EasySwoole\Component\Di;
use EasySwoole\EasySwoole\Swoole\EventRegister;
use EasySwoole\EasySwoole\AbstractInterface\Event;
use EasySwoole\Http\Request;
use EasySwoole\Http\Response;

class EasySwooleEvent implements Event
{

    public static function initialize()
    {
        // TODO: Implement initialize() method.
        date_default_timezone_set('Asia/Shanghai');
        // 在全局方法中注册一个类
        Di::getInstance()->set(SysConst::HTTP_EXCEPTION_HANDLER,[ExceptionHeadler::class,'handle']);
    }

    public static function mainServerCreate(EventRegister $register)
    {

        /**
         * 单独端口开启TCP服务器
         */
        /*$server = ServerManager::getInstance()->getSwooleServer();

        $subPort1 = $server->addListener('0.0.0.0',9502,SWOOLE_TCP);
        $subPort1->set(
            ['open_length_check'=>false]//不验证数据包
        );
        $subPort1->on('connect',function(\swoole_server $server,int $fd,int $reactor_id){
            echo "fd:{$fd} 已连接\n";
            $str = "连接成功";
            $server->send($fd,$str);
        });
        $subPort1->on('close',function(\swolle_server $server,int $fd,int $reactor_id){
            echo "fd:{$fd}已关闭\n";
        });
        $subPort1->on('receive',function(\swoole_server $server,int $fd,int $reactor_id,string $data){
            echo "fd:{$fd}发送消息:{$data}\n";
        });*/

        /**
         * websocket控制器
         */
        //创建一个Dispatcher配置
        $conf = new \EasySwoole\Socket\Config();
        //设置Dispatcher为WebSocket模式
        $conf->setType(\EasySwoole\Socket\Config::WEB_SOCKET);
        //设置解析器对象
        $conf->setParser(new WebSocketParser());
        //创建Dispatcher对象，并注入config对象
        $dispatch = new Dispatcher($conf);
        //给server注册相关事件，在WebSocket模式下，on message 事件必须注册并且交给dispatcher对象
        $register->set(EventRegister::onMessage,function(\swoole_websocket_server $server,\swoole_websocket_frame $frame) use($dispatch){
            $dispatch->dispatch($server,$frame->data,$frame);
        });
        //自定义握手事件
        $webSocketEvent = new WebSocketEvent();
        $register->set(EventRegister::onHandShake,function(\swoole_http_request $request,\swoole_http_response $response) use($webSocketEvent){
            $webSocketEvent->onHandShake($request,$response);
        });
        //自定义关闭事件
        $register->set(EventRegister::onClose,function(\swoole_server $server,int $fd,int $reactorId) use($webSocketEvent){
            $webSocketEvent->onClose($server,$fd,$reactorId);
        });
        // TODO: Implement mainServerCreate() method.
    }

    public static function onRequest(Request $request, Response $response): bool
    {
        // TODO: Implement onRequest() method.
        return true;
    }

    public static function afterRequest(Request $request, Response $response): void
    {
        // TODO: Implement afterAction() method.
    }
}