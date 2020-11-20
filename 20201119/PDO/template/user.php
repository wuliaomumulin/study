<?php
/**
* 定义一个表名称
*/
$table = [
 	'table' => [
 		//定义一个前缀
 		'prefix' =>  'user_',
 		//定义一个唯一值，用于确定表
 		//'unique' =>  date('Y-m-d',strtotime("-1 day")),
 		'unique' =>  date('Y-m-d'),
 	],
 	//定义一个统一的语法结构体，来包含表格式
 	'content' => "(
		  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
		  `username` varchar(50) NOT NULL COMMENT '用户名',
		  `password` varchar(100) DEFAULT NULL COMMENT '密码',
		  `salt` varchar(20) DEFAULT NULL COMMENT '盐',
		  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
		  `mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
		  `status` tinyint(4) DEFAULT NULL COMMENT '状态  0：禁用   1：正常',
		  `create_user_id` bigint(20) DEFAULT NULL COMMENT '创建者ID',
		  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
		  PRIMARY KEY (`user_id`),
		  UNIQUE KEY `username` (`username`)
		)"
 ];






 /*
 * 返回一个可以执行的sql建表语句
 * 表不存在，则创建;
 */

return [
	'table' =>  $table['table']['prefix'].$table['table']['unique'],
	'sql' => 'CREATE TABLE IF NOT EXISTS `'.$table['table']['prefix'].$table['table']['unique'].'` '.$table['content'],
];
?>