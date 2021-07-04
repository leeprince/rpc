<?php

/**
php swoole 作为服务端 接收 go jsonrpc 客户端的rpc请求
*/

//创建Server对象，监听 127.0.0.1:9501 端口
$server = new Swoole\Server('127.0.0.1', 12345);

//监听连接进入事件
$server->on('Connect', function ($server, $fd) {
    echo "Client: Connect.\n";
});

//监听数据接收事件
$server->on('Receive', function ($server, $fd, $reactor_id, $data) {
    var_dump('接收 go jsonrpc 客户端的rpc请求的参数', $data);
    /** 按照 go jsonrpc 的格式返回给go jsonrpc 客户端 */
    // 同步请求回复
    $result = [
        "id" => 0,
        "result" => 1000,
        "error" => null
    ];
    $returnSyncData = json_encode($result);
    $server->send($fd, $returnSyncData);

    // 异步请求回复
    $result = [
        "id" => 1,
        "result" => 1000,
        "error" => null
    ];
    $returnAsyncData = json_encode($result);
    $server->send($fd, $returnAsyncData);
});

//监听连接关闭事件
$server->on('Close', function ($server, $fd) {
    echo "Client: Close.\n";
});

//启动服务器
$server->start();