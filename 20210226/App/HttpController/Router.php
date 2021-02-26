<?php
namespace App\HttpController;

use EasySwoole\Http\AbstractInterface\AbstractRouter;
use FastRoute\RouteCollector;
use EasySwoole\Http\Request;
use EasySwoole\Http\Response;

class Router extends AbstractRouter{
	function initialize(RouteCollector $routeCollector){
		//示例代码
		//$routeCollector->get('user','index.html');
		$routeCollector->get('/WebSocket/index','WebSocket/index');
		$routeCollector->post('user','User');
		$routeCollector->get('rpc','Rpc/index');

		$routeCollector->get('/',function(Request $request,Response $response){
			$response->write('this router index');
		});

		$routeCollector->get('/test',function(Request $request,Response $response){
			$response->write('this router test');
			return '/a';//重新定位到/a方法
		});
		$routeCollector->get('/user/{id:\d+}',function(Request $request,Response $response){
			$response->write("this is router user,your id is {$request->getQueryParam('id')},请求头:{$request}");
			var_export($request->getHeaders());//获得头信息

			
			return false;//不再往下请求，结束本次响应
		});

		
	}
}

?>