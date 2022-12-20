<?php

require_once(__DIR__ . "/../vendor/autoload.php");
require_once(__DIR__ . "/../SmartyHelper.php");
require_once (__DIR__. "/class/TestDbConnect.php");



$instance = TestDbConnect::getInstance();
$instance->showPage();