<?php 
namespace App\HttpController;

use EasySwoole\Easyswoole\Trigger;
use EasySwoole\Http\AbstractInterface\Controller;
use EasySwoole\Http\Message\Status;
use EasySwoole\VerifyCode\Conf;

class User extends Controller{
	function index(){
		//验证码
		$code = new \EasySwoole\VerifyCode\VerifyCode(new Conf());
		$this->response()->withHeader('Content-Type','image/png');
		$this->response()->write($code->DrawCode()->getImageByte());
		//$this->writejson(200,[],'success');
	}
	function getBase64(){
        $config = new Conf();
        $code = new \EasySwoole\VerifyCode\VerifyCode($config);
        $this->response()->write($code->DrawCode()->getImageCode());
    }
    
}

?>