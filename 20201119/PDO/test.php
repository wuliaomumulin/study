<?php
/*header("Content-type:text/html;charset=utf-8");
//mysql数据库连接方式
$conn = @mysql_connect("localhost",'root','root') or die('数据库连接错误');
mysql_select_db("test",$conn);
mysql_query('set names utf8');
print('数据库连接成功');

//mysqli连接
$conn1 = mysqli_connect("localhost",'root','root','test');
if(!$conn1) print('数据库连接失败'.mysqli_connect_error());
else print('数据库连接成功');

//PDO连接
try{
	$pdo = new PDO("mysql:host=localhost;dbname=test","root",'root');
}catch(PDDException $e){
	print('数据库连接失败');
}
print('数据库连接成功');*/

require_once 'PDO.class.php';
//$db = DB::getInstance("182.92.148.54","root",'logis2016','plana','utf8');
$db = DB::getInstance("127.0.0.1","root",'root','test','utf8');

//$citys  = $db->query('select id,city from citys','All');

//获取最大值
/*$getMaxValue = $db->getMaxValue('citys','id','id<150');
echo $getMaxValue;*/

//$update = $db->update('test1',array('name'=>'张三','sex'=>1,'address'=>'合肥'),'id=1');
//print($update);
//带有主键的更新方案,将采取替换原数据:
/*$update1 = $db->update('test1',array('id'=>1,'name'=>'张三','sex'=>1,'address'=>'西安'),'',1);
print($update1);*/

//插入
/*$insert = $db->insert('test1',array('name'=>'张佳佳','sex'=>1,'address'=>'西红门'),1);
print($insert);*/

//如果replace没有主键，则执行插入操作数据,返回值1;如果有主键，则替换返回值2
/*$replace = $db->replace('test1',array('id'=>3,'name'=>'李四','sex'=>1,'address'=>'天通苑'));
print($replace);*/

/*$delete = $db->delete('test1','id=9',1);
print($delete);*/

//统计字段数
/*$getCount = $db->getCount('citys','id','id<150');
echo $getCount;*/

//获取表引擎
echo $db->getTableEngine('test','test1');

//事物处理
/*$arraySql = array('insert into test1 values(12,"张薪薪",0,"物资学院")','replace into test1 values(3,"刘铁锁",1,"东二旗")');
print($db->execTransaction($arraySql));*/

//销毁数据库对象连接
print($db->destruct());

//$result=$db->query('plana','bizs');

/*$str = 'values("';
foreach($result as $arr){
	if(is_array($arr)) $str.=implode('","',$arr);
}

$str.='")';
echo $str;*/
