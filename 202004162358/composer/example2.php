<?php
//示例中的代码其实就是 ThinkPHP 自动加载器源码的精简版，它是 ThinkPHP 5 能实现惰性加载的关键。
include 'Loader.php';

spl_autoload_register('Loader::autoload');

new \app\mvc\controller\admin\Index();
?>