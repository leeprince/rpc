<?php

/**
php swoole 作为客户端发送请求到 go jsonrpc 服务端
*/

use Swoole\Coroutine\Client;
use function Swoole\Coroutine\run;

run(function () {
    $client = new Client(SWOOLE_SOCK_TCP);
    if (!$client->connect('127.0.0.1', 12345, 0.5))
    {
        echo "connect failed. Error: {$client->errCode}\n";
    }

    /** 按照 go jsonrpc 的格式发送给 go jsonrpc 服务端 */
    $sendData = [
        "method" => "arith.Multiply",
        "params" => [
            [
                'A' => 7,
                'B' => 8
            ]
        ],
        "id" => 0
    ];
    $sendData = json_encode($sendData);
    var_dump($sendData);
    $sendBool = $client->send($sendData);
    if (!$sendBool) {
        die("send failed.");
    }

    $recvData = $client->recv();
    if (empty($recvData)) {
        die("recv failed.");
    }
    var_dump("接收 go jsonrpc 服务端的响应", $recvData);

    $client->close();
});