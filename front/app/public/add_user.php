<?php

use GuzzleHttp\Client;
use GuzzleHttp\Exception\ClientException;


require_once(__DIR__ . "/../vendor/autoload.php");
require_once(__DIR__ . "/../SmartyHelper.php");

class AddUserForm
{
    public $user_id;
    public $password;
    public $name;
    public $smarty;

    /**
     * @param string $name
     */
    public function __construct(string $name)
    {
        $this->name = $name;
        $this->user_id = "";
        $this->smarty = new Smarty();
        SmartyHelper($this->smarty);
    }

    /**
     * @return void
     * @throws SmartyException
     */
    public function showForm()
    {
        $this->smarty->assign("title", $this->name);
        $this->smarty->assign("user_id", $this->user_id);
        $this->smarty->assign("password", $this->password);
        $this->smarty->display("form.tpl");
    }

    /**
     * @throws SmartyException
     */
    public function doneForm()
    {
        session_start();
        //TODO 安全に値を取る方法を考える
        $id = $_POST['id'];
        $password = $_POST['password'];
        $this->smarty->assign("access_token", "Failed to fetch token");

        // 認可リクエスト
        $client = new Client(['cookies' => true]);
        try {
            //TODO 認可エンドポイントのパラメーターを追加
            $res = $client->get('https://api.digital-future.jp/auth', ["query" => [
                "client_id" => "example-client-id-1",
                "scope" => "hoge",
                "state" => "hoge",
                "redirect_url" => "https://api.digital-future.jp/public/form_confirm.php",
            ]]);
        } catch (ClientException $e) {
            trigger_error($e->getMessage(), E_USER_WARNING);
        }

        // 認証プロバイダーへ情報を送信
        $code = "";
        $url = "";
        $parseUrl = "";
        try {
            $res = $client->post('https://api.digital-future.jp/auth', ['forms_params' => [
                "user_id" => "example-user-id-1",
                "password" => "password",
                "client_id" => "example-client-id-1",
            ], 'allow_redirects' => false]);
            $locationHeader = $res->getHeader("Location");
            $url = $locationHeader[0];
            $parseUrl = parse_url($url);
            $outputs = [];
            parse_str($parseUrl['query'], $output);
            $code = $output["code"];
        } catch (ClientException $e) {
            trigger_error($e->getMessage(), E_USER_WARNING);
        }

        //トークンを取得
        try {
            $res = $client->get('https://api.digital-future.jp/token', [
                "query" => [
                    "code" => $code,
                    "client_id" => "example-client-id-1",
                    "grant_type" => "authorization_code",
                    "client_secret" => "secret",
                    "scope" => "hoge",
                    "redirect_uri" => $url,
                ]
            ]);
            $body = json_decode($res->getBody()->getContents(), true);
            $this->smarty->assign("access_token", $body["access_token"]);
        } catch (Exception $e) {
            trigger_error($e->getMessage(), E_USER_WARNING);
        }
        //TODO API Request
        $this->smarty->display("confirm.tpl");
    }
}
