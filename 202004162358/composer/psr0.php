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
/*
composer命名自动加载更新命名空间技巧：
composer dump-autoload # 更新autoload_static.php、autoload_psr4.php、composer.json
composer dump-autoload -o # classmap，（-o 等同于 --optimize）
composer dump-autoload -a # 权威classmap
composer dump-autoload --apcu # APCu cache优化类文件
 */
?>