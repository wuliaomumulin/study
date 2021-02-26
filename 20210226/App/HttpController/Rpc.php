<?php 
namespace App\HttpController;

use Easyswoole\Easyswoole\Trigger;
use EasySwoole\Http\AbstractInterface\Controller;
use Easyswoole\Http\Message\Status;

class Rpc extends Controller{
	
	private $userid;//一经赋值，数据存入进程当中
	

	//代理__contruct()，我们一般不要使用该方法，因为底层由该方法实现
	protected function onRequest(?string $action): ?bool
	{
		return parent::onRequest($action);
	}
	//当请求结束的时候，我们使用gc方法对私有属性的回收,如session的
	protected function gc(){
		$this->userid = null;//直接赋值null，不要使用unset，因为unset是给变量赋值引用的null
		parent::gc();
	}

	//代理__destruct()
	//void：代表没有返回值
	//?:代表形参可为null
	protected function afterAction(?string $actionname): void{
		parent::afterAction($actionname);
	}
	//方法找不到
	protected function actionNotFound(?string $action){
		$this->response()->withStatus(Status::CODE_NOT_FOUND);
        $this->response()->write('action not found');
	}

	function index(){
		$this->writejson(200,[],'success');
		return '/test';
	}
	function test(){
		$this->response()->write('this is test');
		return '/test2';
	}
	function test2(){
		$this->response()->write('this is test2');
		return true;
	}
}

?>