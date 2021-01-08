<?php
/**
* 判断机器字节序是大端小端
* true:大端
* false:小端
*/
function beg_endian(){
 return (pack('L',1) === pack('N',1));
}
?>
