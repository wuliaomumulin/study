<?php
//定时器
Swoole\Timer::tick(1000,function(){
    echo date('H:i:s')."\n";
});