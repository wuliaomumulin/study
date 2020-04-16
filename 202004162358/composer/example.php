<?php

//可以传入多个参数，逗号分隔，多个相同的类名最先找到的最先匹配
spl_autoload_extensions('.class.php,.php');

//设置autoload寻找php定义的类文件的目录，PATH_SEPARATOR常量分隔符，linux(:),windows(;)
$path = get_include_path() . PATH_SEPARATOR . "libs/";
set_include_path($path);

//autoload注册
spl_autoload_register();

new Test();
new User();
echo $path;
?>
