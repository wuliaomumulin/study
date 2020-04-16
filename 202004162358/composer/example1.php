<?php
$class = 'app\view\news\Index';

//顶级命名空间路径映射
$vendor_map = [
	'app' =>  __DIR__,
];

$vendor = substr($class,0,stripos($class,'\\'));//取出顶级命名空间[app]
$vendor_dir = $vendor_map[$vendor];//文件基目录
$rel_path = dirname(substr($class,strlen($vendor)));//相对路劲
$file_name = basename($class).'.php';//文件名Index.php

//输出文件所在路径
//DIRECTORY_SEPARATOR文件分隔符，windows为\或/，linux为/
echo $vendor_dir.$rel_path.DIRECTORY_SEPARATOR.$file_name;
?>