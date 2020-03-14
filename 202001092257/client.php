<?php

$string = 'hello,world!';

$sock = stream_socket_client('unix:///var/tmp/test_domain_socket', $errno, $errstr);

fwrite($sock, $string."\r\n");

echo fread($sock, 4096)."\n";

fclose($sock);

?>
