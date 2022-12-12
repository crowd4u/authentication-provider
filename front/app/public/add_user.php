<?php

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

        // 認可リクエスト
        $client = new Client(['cookies' => true]);
        try {
            //TODO 認可エンドポイントのパラメーターを追加
            $res = $client->request("GET", 'http://n4u-api/auth');
        } catch (\Guzzle\Http\Exception\ClientErrorResponseException $e) {
            trigger_error($e->getMessage(), E_USER_ERROR);
        }
        $cookies = $client->getConfig('cookies');
        $cookie = $res->getHeader('Set-Cookie');


        //TODO API Request
        $this->smarty->display("confirm.tpl");
    }


}