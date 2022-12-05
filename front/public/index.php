<?php
require_once(__DIR__ . "/../vendor/autoload.php");
require_once(__DIR__ . "/../SmartyHelper.php");

//Smartyの読み込み
$smarty = new Smarty();
//Smartyの初期化
SmartyHelper($smarty);

try {
    $smarty->display("index.tpl");
} catch (SmartyException $e) {
    trigger_error("Failed to display", E_USER_ERROR);
}
