<?php 
namespace App\HttpController;

use Easyswoole\Easyswoole\Trigger;
use EasySwoole\Http\AbstractInterface\Controller;
use Easyswoole\Http\Message\Status;

class Http extends Controller{
	

	/**
	 * http服务错误与异常
	 * 1、控制器级别的异常处理
	 * 2、全局异常处理
	 *
	 * 怎么使用和源码处理异常?
	 */
	
	//当请求结束的时候，我们使用gc方法对私有属性的回收,如session
	protected function onException(\Throwable $throwable): void
	{
		var_dump($throwable->getMessage());
	}


	function index(){
		//直接抛出异常
		throw new Exception('控制器异常');	
	}
	
}

?>