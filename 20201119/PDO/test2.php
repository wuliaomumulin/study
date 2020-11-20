<?php


require_once 'TimeLine.class.php';

$template = require_once 'template/user.php';
//exit($user);


$db = new TimeLine("127.0.0.1","root",'phpcj','test','utf8');


/*$data = [
		'username' => mt_rand(10,1000).'liyb',
		'password' => mt_rand(10,1000).'password',
		'salt' => mt_rand(10,1000).'salt',
		'email' => mt_rand(10,1000).'email',
		'mobile' => mt_rand(10,1000).'mobile',
		'status' => mt_rand(0,1),
		'create_user_id' => mt_rand(0,100),
		'create_time' => '2020-11-19 12:00:00',
	];

$count = $db->batchInsert($data,$template,'create_time');



echo $count;*/


$data = [
	[
		//'user_id' => 2,
		'username' => mt_rand(10,1000).'liyb',
		'password' => mt_rand(10,1000).'password',
		'salt' => mt_rand(10,1000).'salt',
		'email' => mt_rand(10,1000).'email',
		'mobile' => mt_rand(10,1000).'mobile',
		'status' => mt_rand(0,1),
		'create_user_id' => mt_rand(0,100),
		'create_time' => '2020-11-19 12:00:00',
	],
	[
		//'user_id' => 3,
		'username' => mt_rand(10,1000).'liyb',
		'password' => mt_rand(10,1000).'password',
		'salt' => mt_rand(10,1000).'salt',
		'email' => mt_rand(10,1000).'email',
		'mobile' => mt_rand(10,1000).'mobile',
		'status' => mt_rand(0,1),
		'create_user_id' => mt_rand(0,100),
		'create_time' => '2020-11-20 12:00:00',
	],
	[
		//'user_id' => 4,
		'username' => mt_rand(10,1000).'liyb',
		'password' => mt_rand(10,1000).'password',
		'salt' => mt_rand(10,1000).'salt',
		'email' => mt_rand(10,1000).'email',
		'mobile' => mt_rand(10,1000).'mobile',
		'status' => mt_rand(0,1),
		'create_user_id' => mt_rand(0,100),
		'create_time' => '2020-11-21 12:00:00',
	],
];

$count = $db->batchInsert($data,$template,'create_time');
echo $count;

/*$tables = $db->getTables('sys_user%');
$data = $db->batchQuery($tables,'*',true);
var_dump($data);*/





//var_dump($b);
print($db->destruct());





/*class A{

	public function __construct(){
		echo 'init';
	}
	
	public function aa(){
		return __METHOD__;
	}

}

class B extends A{



	public function bb(){
		return __METHOD__;
	}

}

$B = new B();
echo $B->aa();
echo $B->bb();*/
?>