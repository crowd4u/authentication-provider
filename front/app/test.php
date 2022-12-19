<?php


use GuzzleHttp\Client;

require_once(__DIR__ . "/vendor/autoload.php");
// 認可リクエスト
$client = new Client(['cookies' => true]);
try {
//TODO 認可エンドポイントのパラメーターを追加
    $res = $client->get('http://api:8080/auth', ["query" => [
        "client_id" => "example-client-id-1",
        "scope" => "hoge",
        "state" => "hoge",
        "redirect_url" => "http://api:8081",
    ]]);
    var_dump($res->getResponse());
} catch (\GuzzleHttp\Exception\ClientException $e) {
    trigger_error($e->getMessage(), E_USER_WARNING);
    echo $e->getMessage();
}