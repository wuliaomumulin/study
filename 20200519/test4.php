<?php
//go关键词创建一个协程，可以简单的理解为创建了一个线程

Swoole\Runtime::enableCoroutine();//此行代码后，文件操作，sleep，Mysqli，PDO，streams等都变成异步IO，见'一键协程化'章节
$s = microtime(true);
//Co/run()见'协程容器'章节
Co\run(function(){
// i just want to sleep...
    for ($c = 100; $c--;) {
        go(function () {
            for ($n = 100; $n--;) {
                usleep(1000);
            }
        });
    }

    // 10k file read and write
    for ($c = 100; $c--;) {
        go(function () use ($c) {
            $tmp_filename = "/tmp/test-{$c}.php";
            for ($n = 100; $n--;) {
                $self = file_get_contents(__FILE__);
                file_put_contents($tmp_filename, $self);
                assert(file_get_contents($tmp_filename) === $self);
            }
            unlink($tmp_filename);
        });
    }
    // 10k pdo and mysqli read
    for ($c = 50; $c--;) {
        go(function () {
            $pdo = new PDO('mysql:host=127.0.0.1;dbname=test;charset=utf8', 'root', '123456');
            $statement = $pdo->prepare('SELECT * FROM `yaf_user`');
            for ($n = 100; $n--;) {
                $statement->execute();
                assert(count($statement->fetchAll()) > 0);
            }
        });
    }
    for ($c = 50; $c--;) {
        go(function () {
            $mysqli = new Mysqli('127.0.0.1', 'root', '123456', 'test');
            $statement = $mysqli->prepare('SELECT `id` FROM `yaf_user`');
            for ($n = 100; $n--;) {
                $statement->bind_result($id);
                $statement->execute();
                $statement->fetch();
                assert($id > 0);
            }
        });
    }


     // php_stream tcp server & client with 12.8k requests in single process
    function tcp_pack(string $data): string
    {
        return pack('n', strlen($data)) . $data;
    }

    function tcp_length(string $head): int
    {
        return unpack('n', $head)[1];
    }

    go(function () {
        $ctx = stream_context_create(['socket' => ['so_reuseaddr' => true, 'backlog' => 128]]);
        $socket = stream_socket_server(
            'tcp://0.0.0.0:9502',
            $errno, $errstr, STREAM_SERVER_BIND | STREAM_SERVER_LISTEN, $ctx
        );
        if (!$socket) {
            echo "$errstr ($errno)\n";
        } else {
            $i = 0;
            while ($conn = stream_socket_accept($socket, 1)) {
                stream_set_timeout($conn, 5);
                for ($n = 10; $n--;) {
                    $data = fread($conn, tcp_length(fread($conn, 2)));
                    echo $data.PHP_EOL;
                    assert($data === "你好，TCP服务 #{$n}!");
                    fwrite($conn, tcp_pack("你好，TCP用户 #{$n}!"));
                }
                if (++$i === 1) {
                    fclose($socket);
                    break;
                }
            }
        }
    });
    for ($c = 1; $c--;) {
        go(function () {
            $fp = stream_socket_client("tcp://127.0.0.1:9502", $errno, $errstr, 1);
            if (!$fp) {
                echo "$errstr ($errno)\n";
            } else {
                stream_set_timeout($fp, 5);
                for ($n = 10; $n--;) {
                    fwrite($fp, tcp_pack("你好，TCP服务 #{$n}!"));
                    $data = fread($fp, tcp_length(fread($fp, 2)));
                    echo $data.PHP_EOL; 
                    assert($data === "你好，TCP用户 #{$n}!");
                }
                fclose($fp);
            }
       });
    }

    // udp server & client with 12.8k requests in single process
    go(function(){
        /**
        * AF_INET (IPV4)
        * SOCK_DGRAM (udp),SOCK_STREAM(tcp)
        */
        $socket = new Swoole\Coroutine\Socket(AF_INET,SOCK_DGRAM,0);
        $socket->bind('127.0.0.1',9503);
        $client_map = array();
        for($c=1;$c--;){
            for($n=10;$n--;){
                $recv = $socket->recvfrom($peer);
                $client_uid = "{$peer['address']}:{$peer['port']}"; 
                $id = $client_map[$client_uid] = ($client_map[$client_uid]??-1)+1;
                $socket->sendto($peer['address'],$peer['port'],"UDP服务:你好,{$id}");
                assert($recv === "UDP客户:你好,{$n}");//断言返回bool值
                echo $recv.PHP_EOL;
            }
        }
        $socket->close();
    });

    for($c=1;$c--;){
        go(function(){
            $fp = stream_socket_client('udp://127.0.0.1:9503',$errno,$errstr,1);
            if(!$fp){
                echo '$errstr({$errno})';
            }else{
                for($n=10;$n--;){
                    fwrite($fp,"UDP客户:你好{$n};");
                    $recv = fread($fp, 1024);
                    list($address,$port) = explode(':',stream_socket_get_name($fp,true));//there
                    assert($recv === "UDP服务:你好,{$n}");//断言返回bool值
                    echo $recv.PHP_EOL;
                }
                fclose($fp);
            }

        });
    }

   
});
echo 'use ' . (microtime(true) - $s) . ' s'.PHP_EOL;

/*class test{
	public function test_mysqli(){
		$mysqli = new Mysqli('127.0.0.1','root','123456','test');
		$statement = $mysqli->prepare('select id,username from yaf_user');
		$statement->bind_result($id,$username);
		$statement->execute();
		while($statement->fetch()){
		 	print('id='.$id.',username='.$username.PHP_EOL);
		}
		$statement->close();
		$mysqli->close();
	}
}*/

?>