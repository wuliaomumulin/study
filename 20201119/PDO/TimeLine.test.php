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
	public function batchQuery(Array $tables,String $fields = '*',Bool $hidden_table=false){
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

	/**
	* 批量插入
	* @param $data Array 二维数组或一维数组，key=>val结构
	* @param $template Array 表名和表sql
	* @param $standard String 判断标准
	*/
	public function batchInsert(Array $data,Array $template,String $standard){

		if(is_int($this->execSql($template['sql']))){

			//$this->debug = true;


			//判断是否一维数组

			if(count($data) == count($data,1)){
				//一维数组
				return $this->insert($template['table'],$data);

			}else{
				//多维数组				
				return $this->insertCluster($template['table'],$data);
				//return $this->replaceCluster($template['table'],$data);
			}
		}

	}

	/**
	* 解析当前表名称
	* @param $standard String 判断标准
	*/
	private function parseTableName(String $standard){
		/*if(is_numeric($standard)){
			exit('数字');
		}else{
			exit('string');
		}*/
	}


}

?>
