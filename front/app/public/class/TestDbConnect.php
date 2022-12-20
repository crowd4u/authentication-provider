<?php

use GuzzleHttp\Client;
use GuzzleHttp\Exception\ClientException;

require_once(__DIR__ . "/../../vendor/autoload.php");
require_once(__DIR__ . "/../../SmartyHelper.php");

class TestDbConnect
{
    private $smarty;
    /**
     * @var Client
     */
    private $client;
    private $db;
    private static $instance;


    public function __construct()
    {
        $this->smarty = new Smarty();
        SmartyHelper($this->smarty);
        $this->client = new Client(['cookies' => true]);
        try {
            $this->db = new PDO('mysql:dbname=s2113591;host=localhost;charset=utf8', 's2113591', 'tsukuba');
        } catch (PDOException $e) {
            trigger_error($e, E_USER_WARNING);
        }
        try {
            //TODO 認可エンドポイントのパラメーターを追加
            $res = $this->client->get('https://api.digital-future.jp/auth', ["query" => [
                "client_id" => "example-client-id-1",
                "scope" => "hoge",
                "state" => "hoge",
                "redirect_url" => "https://api.digital-future.jp/public/form_confirm.php",
            ], 'allow_redirects' => false]);
        } catch (ClientException $e) {
            trigger_error($e->getMessage(), E_USER_WARNING);
        }
    }

    public static function getInstance()
    {
        if (!isset(self::$instance)) {
            self::$instance = new TestDbConnect();
        }
        return self::$instance;
    }

    public function addUser()
    {
        $token = $_POST['token'];
        //認可リクエスト
        try {
            //TODO 認可エンドポイントのパラメーターを追加
            $res = $this->client->get('https://api.digital-future.jp/token/check', ["query" => [
                "token" => $token,
            ]]);
        } catch (ClientException $e) {

        }

        if ($res->getStatusCode() != 200) {
            trigger_error("error, failed to insert data", E_USER_ERROR);
        }

        $statement = "insert into users(id,email,user_name,hash_password,created_at) values(?,?,?,?,CURRENT_TIMESTAMP)";
        $stmt = $this->db->prepare($statement);
        $id = uniqid();
        $stmt->execute(array($id, $id . "@example.com", "john","password"));
    }

    private
    function loadData()
    {
        $sth = $this->db->query('select id,email,user_name from users');
        $entry = $sth->fetchAll(PDO::FETCH_ASSOC);
        $this->smarty->assign("data", $entry);
    }

    public
    function showPage()
    {
        $this->loadData();
        $this->smarty->display("test.tpl");
    }
}
