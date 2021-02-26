<?php 
/**
 * 定义异常错误类
 */
namespace App\Exception;

use EasySwoole\Http\Request;
use EasySwoole\Http\Response;

class ExceptionHeader{
	public static function handle(\Throwable $exception,Request $request,Response $response){
		var_dump($exception->getMessage());
	} 
}
?>