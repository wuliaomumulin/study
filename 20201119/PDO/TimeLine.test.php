<?php

require_once 'PDO.class.php';


/**
* 将mysql变为时序数据库

*/
class TimeLine extends DB{


	/**
	* 
	* 和Es对应的查询索引
	* @param $like 字符串 topology_%
	* % 代表多个字符 _ 代码单个字符
	* @param $return * topology_edg_ 
	*/
	public function getTables(string $like){
		$ret = [];
		$strSql = 'show tables like "'.$like.'"';
        $arrayTableInfo=$this->query($strSql);
        foreach($arrayTableInfo as $arr) $ret[] = array_shift($arr);
        return $ret;
	}

	/**
	* 批量查询
	* @param $tables Array 一维数组，存储表名称
	* @param $fields String 查询字段
	* @param $hidden_table Bool 隐藏数据表名称
	*/
	public function batchQuery(Array $tables,String $fields = '*',$hidden_table=false){
		$ret = [];
		foreach($tables as $table){
			$strSql = "select {$fields} from {$table}";
			$cur = $this->query($strSql);

			if($hidden_table){

				foreach ($cur as $arr) $ret[] = $arr;					

			}else{
				
				$ret[$table] = $cur;
			
			}

		}

		return $ret;

	}
}



//$db = new DB("127.0.0.1","root",'phpcj','test','utf8');
$db = DB::getInstance("127.0.0.1","root",'phpcj','test','utf8');
//$tables = $db->getTables('select * from tb_user');
$tables = $db->getTables('sys_user%');
$data = $db->batchQuery($tables,'*',true);
var_dump($data);

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