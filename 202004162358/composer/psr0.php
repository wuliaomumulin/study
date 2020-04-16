<?php
function autoload($className){
	$className = ltrim($className,'\\');
	$fileName = '';
	$namespace = '';
	if($lastNsPos = strrpos($className,'\\')){
		$namespace = substr($className,0,$lastNsPos);//命名空间
		$className = substr($className,$lastNsPos+1);//类名
		$filename = str_replace('\\',DIRECTORY_SEPARATOR,$namespace).DIRECTORY_SEPARATOR;
	}
	$filename .= str_replace('_',DIRECTORY_SEPARATOR,$className).'.php';

	require $filename;
}

spl_autoload_register('autoload');


//测试
new \app\mvc\controller\admin\Index();
?>