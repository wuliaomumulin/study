<?php
/***
* decsr:统计网卡流量脚本
* @liyb
*
* 一、前提:开启服务前，请保证redis扩展和shell_exec函数处于开启状态

* 二、服务使用
* 		php network.php 开启服务
* 		php network.php& 以后台运行的方式开启服务,该条命令会返回pid，关闭服务的时候使用
*       ps -ef|grep network.php 查找该进程的pid
* 		kill -9 pid 来关闭服务

*/
try{
	$NetworkBytes = new NetworkBytes();
	$NetworkBytes->handle();

}catch(\Exception $e){
	echo $e->getMessage();
}






//3、存入redis
/*
lrange network-bytes 0 10 #指定范围获取列表值
rpush network-bytes 100 #向列表压入数据
lpushx network-bytes 105 #向列表头部压入数据
rpop network-bytes #弹出最后一个元素
llen network-bytes #获取列表长度
*/


/***
* decsr:统计网卡流量脚本
 @liyb
*/
class NetworkBytes{
	//redis句柄
	public $redis = Null;
	//设置抓取网卡名称	
	public $networkName = 'ens33';	
	//操作redis-table的名称
	public $table = 'network-bytes';
	//设置队列长度,60个
	public $listLen = 30;

	//初始化上一秒值为0
	static private $prev = 0;
	//顺序,正序asc|倒叙desc
	static private $order = 'asc'; 

	//@params $environ product|develop
	public function __construct($environ='product'){
		self::connect(self::config($environ));
	}
	/*
	   配置信息
	   @params $environ product|develop
	   @return Array;		
	*/
	private static function config($environ){
		//初始化返回值
		$config = [];
		//定义配置文件位置
		$filename =  dirname(dirname(dirname(__FILE__))).'/conf/application.ini';
		$str = `awk '/redis.config/' {$filename}`;
		
		//获取最后一块的配置
		$ip = '((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}';
		$preg = "/redis\.config\.server=\'({$ip})\'/i";
		preg_match_all($preg,$str,$res);
		$config['hostname'] = ($environ=='product') ? $res[1][0] : $res[1][1];

		$preg = "/redis\.config\.port=\'(\d+)\'/i";
		preg_match_all($preg,$str,$res);
		$config['port'] = ($environ=='product') ? $res[1][0] : $res[1][1];
		return $config;
	}
	/*
	 	初始化连接句柄
	*/
	private function connect($config){
		$redis = new Redis();
		$redis->connect($config['hostname'],$config['port']);
		$this->redis = $redis;
	}
	/*
	  业务操作
	*/
	public function handle(){
		//常驻进程
		while(True){			
			//沉睡1s，根据服务器性能进行配置，如果性能好的话，去掉即可
			sleep(1);

			//保存网卡流量
			self::procducer();
		}
	}
	/**
	*	统计网络流量脚本
	*/
	private function procducer(){
		$str = `ifconfig {$this->networkName}`;//指定网卡名称
		$preg = '/TX packets \d+ +bytes (\d+)/i';//比特位
		preg_match($preg,$str,$res);
		$current = intval($res[1]);

		switch (static::$order){
			case 'asc':
				if($this->redis->llen($this->table) == $this->listLen){
					$this->redis->lpop($this->table);
				}
				$this->redis->rpush($this->table,$current-static::$prev);
				break;
			case 'desc':
				if($this->redis->llen($this->table) == $this->listLen){
					$this->redis->rpop($this->table);
				}
					$signal = $this->redis->lpushx($this->table,$current-static::$prev);
				//如果是第一次操作，我们创建一个空链表
				if($signal === 0){
					$this->redis->lpush($this->table,$current);
				}
				break;
			default:
				# code...
				break;
		}

		//将当前值设为过期，就是上一秒值
		static::$prev = $current;

	}

}
?>
