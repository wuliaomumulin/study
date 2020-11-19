<?php
/**
 * PDO php class
 * author lin <wuliaomumulin@163.com>
 * version 1.0
 */
	class DB{
		protected static $_instance = null;
		protected $dbName = '';
		protected $dsn;
		protected $dbh;
		/**
		 * 构造函数
		 */
		private function __construct($dbHost,$dbUser,$dbPasswd,$dbName,$dbCharset){
			try{
				$this->dsn = 'mysql:host='.$dbHost.';dbname='.$dbName;
				$this->dbh = new PDO($this->dsn,$dbUser,$dbPasswd);
				$this->dbh->exec('set character_set_connection='.$dbCharset.', character_set_results='.$dbCharset.', character_set_client=binary');
			}catch(PDOException $e){
				$this->outputError($e->getMessage());
			}
		}
		/**
		 * 防止克隆
		 */
		private function __clone(){}
		/**
		 * 实例化一个数据库
		 * @param  [string] $dbName    
		 * @param  [string] $dbUser    
		 * @param  [string] $dbPasswd  
		 * @param  [string] $dbName    
		 * @param  [string] $dbCharset 
		 * @return [object] 数据库实例  
		 */        
		public function getInstance($dbHost,$dbUser,$dbPasswd,$dbName,$dbCharset){
			if(self::$_instance==null){
				self::$_instance = new self($dbHost,$dbUser,$dbPasswd,$dbName,$dbCharset);
			}
			return self::$_instance;
		}
		/**
		 * [query description]
		 * @param  [string]  $sqlstr    [查询字符串]
		 * @param  string  $querymode [查询模式，All:全部|Row：单条]
		 * @param  boolean $debug     [是否返回sql语句，默认false]
		 * @return [array]    $result         [数据]
		 */
		public function query($sqlstr,$querymode='All',$debug=false){
				if($debug==true) $this->debug($sqlstr);
				$recordset = $this->dbh->query($sqlstr);
				$this->getPDOError();
				if($recordset){
					$recordset->setFetchMode(PDO::FETCH_ASSOC);
					if($querymode=='All'){
						$result = $recordset->fetchAll();
					}elseif($querymode=='Row'){
						$result = $recordset->fetch();
					}
				}else{
					$result = '';
				}
				return $result;
		}
		/**
		 * [update 更新]
		 * @param  [type] $table          [表名]
		 * @param  [type] $arrayDataValue [数据]
		 * @param  string $where          [条件]
		 * @param  [type] $debug          [错误输出]
		 * @return [type]                 [description]
		 */
		public function update($table,$arrayDataValue,$where='',$debug=false){
			$this->checkFields($table,$arrayDataValue);
			if($where){
				$strSql = '';
				foreach($arrayDataValue as $key => $val){
					$strSql.=','.$key.'="'.$val.'"';
				}
				$strSql = substr($strSql,1);
				$strSql = 'update '.$table.' set '.$strSql.' where '.$where;
			}else{
				$strSql = 'replace into '.$table.'('.implode(',',array_keys($arrayDataValue)).') values("'.implode('","',$arrayDataValue).'")';
			}
			$result = $this->execSql($strSql,$debug);
			return $result;
		}
		/**
		 * [insert 插入]
		 * @param  [type]  $table          [表名]
		 * @param  [type]  $arrayDataValue [数据]
		 * @param  boolean $debug          [调试]
		 * @return [type]                  [description]
		 */
		public function insert($table,$arrayDataValue,$debug=false){
			$this->checkFields($table,$arrayDataValue);
			$strSql = 'insert into '.$table.'('.implode(',',array_keys($arrayDataValue)).') values("'.implode('","',$arrayDataValue).'")';
			$result = $this->execSql($strSql,$debug);
			return $result;
		}
		/**
		 * [replace 覆盖,需有主键ID]
		 * @param  [type]  $table     
		 * @param  [type]  $arrayDataValue 
		 * @param  boolean $debug       
		 * @return [type]                  
		 */
		public function replace($table,$arrayDataValue,$debug=false){
			$this->checkFields($table,$arrayDataValue);
			$strSql = 'replace into '.$table.'('.implode(',',array_keys($arrayDataValue)).') values("'.implode('","',$arrayDataValue).'")';
			$result = $this->execSql($strSql,$debug);
			return $result;
		}
		/**
		 * [delete 删除]
		 * @param  [string]  $table 
		 * @param  string  $where 
		 * @param  boolean $debug 
		 * @return [type]         
		 */
		public function delete($table,$where='',$debug=false){
			if($where == ''){
				$this->outputError('where is Null');
			}else{
				$strSql = 'delete from '.$table.' where '.$where;
				$result = $this->execSql($strSql,$debug);
				return $result;
			}
		} 
		/**
		 * [execSql 执行sql]
		 * @param  [type]  $strSql 
		 * @param  boolean $debug  
		 * @return [type]   
		 */
		public function execSql($strSql,$debug=false){
			if($debug==true) $this->debug($strSql);
			$result = $this->dbh->exec($strSql);
			$this->getPDOError();
			return $result;
		}
		/**
		 * [getMaxValue 获取最大值]
		 * @param  [type]  $table     
		 * @param  [type]  $fields_name 
		 * @param  string  $where      
		 * @param  boolean $debug     
		 * @return [type]               
		 */
		public function getMaxValue($table,$fields_name,$where='',$debug=false){
			$strSql = 'select max('.$fields_name.') as MAX_VALUE from '.$table;
			if($where!='') $strSql.=' where '.$where;
			if($debug==true) $this->debug($strSql);
			$arrTemp = $this->query($strSql,'Row');
			$maxValue = $arrTemp['MAX_VALUE'];
			if($maxValue==''||$maxValue==Null){
				$maxValue = 0;
			}
			return $maxValue;
		}
		/*获取字段总数*/
		public function getCount($table,$fields_name,$where='',$debug=false){
			$strSql = 'select count('.$fields_name.') as NUM from '.$table;
			if($where!='') $strSql.=' where '.$where;
			if($debug==true) $this->debug($strSql);
			$arrTemp = $this->query($strSql,'Row');
			return $arrTemp['NUM'];
		}
		/*获得表引擎*/
		public function getTableEngine($dbName,$tableName){
			$strSql = 'show table status from '.$dbName.' where name="'.$tableName.'"';
			$arrayTableInfo=$this->query($strSql);
			$this->getPDOError();
			return $arrayTableInfo[0]['Engine'];
		}
		/*开始事务*/
		private function beginTransaction(){
			$this->dbh->beginTransaction();
		}
		/*提交事务*/
		private function commit(){
			$this->dbh->commit();
		}
		/*火棍事务*/
		private function rollback(){
			$this->dbh->rollback();
		}
		/*执行事务*/
		public function execTransaction($arraySql){
			$retval = 1;
			$this->beginTransaction();
			foreach($arraySql as $strSql){
				if($this->execSql($strSql)==0) $retval=0;
			}
			if($retval==0){
				$this->rollback();
				return false;
			}else{
				$this->commit();
				return true;
			}
		}		
		/**
		 * [checkFields 检查指定字段是否在指定数据表中存在]
		 * @param  [string] $table          [表名]
		 * @param  [array]  $arrayDataValue [数据：字段名：数据]
		 */
		private function checkFields($table,$arrayFields){
			$fields = $this->getFields($table);
			foreach ($arrayFields as $key => $value) {
				if(!in_array($key,$fields)){
					$this->outputError('unkown columns'.$key.' in fields list');
				}
			}
		}
		/**
		 * [getFields 获得表字段]
		 * @param  [string] $table [表名称]
		 * @return [array]  $fields [表字段数组]
		 */
		private function getFields($table){
			$fields = array();
			$recordset = $this->dbh->query('show columns from '.$table);
			$this->getPDOError();
			$recordset->setFetchMode(PDO::FETCH_ASSOC);
			$result = $recordset->fetchAll();
			foreach ($result as $a) {
				$fields[]=$a['Field'];
			}
			return $fields;
		}
		/**
		 * 捕捉PDO错误异常信息
		 * @return [type] [description]
		 */
		private function getPDOError(){
			if($this->dbh->errorCode()!='00000'){
				$arrayError = $this->dbh->errorInfo(); 
				$this->outputError($arrayError[2]);
			}
		}
		/**
		 * [outputError 输出错误信息]
		 * @param  [type] $strErrMsg [错误信息]
		 * @return [type]            [description]
		 */
    private function outputError($strErrMsg)
    {
        throw new Exception('MySQL Error: '.$strErrMsg);
    }
		/**
		 * debug sql语句错误信息
		 * @param mixed $debuginfo
		 */
		private function debug($debuginfo){
			var_dump($debuginfo);
			exit();
		}
		/**
		 * [destruct 关闭数据库连接]
		 * 
		 */
		public function destruct(){
			$this->dbh = null;
		}
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
?>



