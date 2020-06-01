<?php

/*
* 约瑟夫环问题
*
*/
$s = 0;
$n=100;$m=3;
for($i=1;$i<=$n;$i++){
        $s = ($s+$m)%$i;
        //echo $s.PHP_EOL;
}
echo ($s+1).PHP_EOL;
