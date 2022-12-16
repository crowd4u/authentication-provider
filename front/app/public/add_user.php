<?php

use Guzzle\Http\Client;

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
            $res = $client->get('http://localhost:8080/auth');
        } catch (\Guzzle\Http\Exception\ClientErrorResponseException $e) {
            trigger_error($e->getMessage(), E_USER_WARNING);
        }

        // 認証プロバイダーへ情報を送信
        $code = "";
        $parseUrl = "";
        try {
            $res = $client->post('http://localhost:8080', [], ['forms_params' => [
                "user_id" => "example-user-id-1",
                "password" => "password",
                "client_id" => "example-client-id-1",
            ]]);
            $locationHeader = $res->getHeader("Location");
            trigger_error($locationHeader, E_USER_NOTICE);
            $parseUrl = parse_url($locationHeader);
            $code = $parseUrl["code"];
        } catch (\Guzzle\Http\Exception\ClientErrorResponseException $e) {
            trigger_error($e->getMessage(), E_USER_WARNING);
        }

        //トークンを取得
        try {
            $res = $client->get('http://localhost:8081/token', [
                "query" => [
                    "code" => $code,
                    "client_id" => "example-client-id-1",
                    "grant_type" => "authorization_code",
                    "client_secret" => "secret",
                    "scope" => "hoge",
                    "redirect_uri" => $parseUrl,
                ]
            ]);
            $body = $res->getResponseBody();
            $this->smarty->assign("access_token", $body["token"]);
        } catch (Exception $e) {
            trigger_error($e->getMessage(), E_USER_WARNING);
        }
        //TODO API Request
        $this->smarty->display("confirm . tpl");
    }
}
