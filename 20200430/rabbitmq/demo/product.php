<?php

header('Content-Type:text/html;charset=utf8;');

$params = array(
    'exchangeName' => 'myexchange',
    'queueName' => 'liyb_queue',
    'routeKey' => 'myroute',
);

$connectConfig = array(
    'host' => 'localhost',
    'port' => 5672,
    'login' => 'admin',
    'password' => 'admin',
    'vhost' => '/'
);

//var_dump(extension_loaded('amqp')); 判断是否加载amqp扩展

//exit();
try {
    $conn = new AMQPConnection($connectConfig);
    $conn->connect();
    if (!$conn->isConnected()) {
        //die('Conexiune esuata');
        //TODO 记录日志
        echo 'rabbit-mq 连接错误:', json_encode($connectConfig);
        exit();
    }
    $channel = new AMQPChannel($conn);
    if (!$channel->isConnected()) {
        // die('Connection through channel failed');
        //TODO 记录日志
        echo 'rabbit-mq Connection through channel failed:', json_encode($connectConfig);
        exit();
    }
    $exchange = new AMQPExchange($channel);
    $exchange->setFlags(AMQP_DURABLE);//持久化
    $exchange->setName($params['exchangeName']?:'');
    $exchange->setType(AMQP_EX_TYPE_DIRECT); //direct类型
    $exchange->declareExchange();

    //$channel->startTransaction();

    $queue = new AMQPQueue($channel);
    $queue->setName($params['queueName']?:'');
    $queue->setFlags(AMQP_DURABLE);
    $queue->declareQueue();

    //绑定
    $queue->bind($params['exchangeName'], $params['routeKey']);
} catch(Exception $e) {

}


$num = mt_rand(100, 500);

//生成消息
for($i = $num; $i <= $num+5; $i++)
{
    $exchange->publish("this is {$i} message..", $params['routeKey'], AMQP_MANDATORY, array('delivery_mode'=>2));
}
