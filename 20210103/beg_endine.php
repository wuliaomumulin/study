<?php


/**
* 判断机器字节序是大端小端
* true:大端
* false:小端
*/
function beg_endian(){
 return (pack('L',1) === pack('N',1));
}

function hex($str){
   return bin2hex($str);
}


function unhex($str){
   $arr = str_split($str,2);
   $val = '';
   foreach ($arr as $v) {
        $val .= hex2bin((int)$v);
   }
   return $val;
 }
/*
sql脚本
CREATE TABLE `test_bin` (
  `bin_id` binary(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

## 测试用例
/*
        不是大端就是小端，每两个字符为一组,第一位为高位，即大端，否则小端
*/
var_dump($a = pack('H*','16D2E6F67D343B87549A1E3964B6449A'));
var_dump(unpack('H*',$a)[1]);

$b = pack('H*','17D8BACFFEE4E3B4E533EF9C08C00EC5');
$c = pack('H*','085D1C4AD4D2C84157AC9FB855EE5993');

//ar_dump($b,$c);exit();

$pdo = new PDO('mysql:host=19.19.19.70;dbname=test;charset=utf8','root','123456'                                                                                               );
$pdo->setAttribute(PDO::ATTR_DEFAULT_FETCH_MODE,PDO::FETCH_ASSOC);//设置筛选模式

$sql = "insert into `test_bin` values('{$a}'),('{$b}'),('{$c}')";

$num = $pdo->exec($sql);

var_dump($num);
?>
