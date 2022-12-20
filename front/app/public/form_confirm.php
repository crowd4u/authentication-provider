<?php
require_once(__DIR__ . "/add_user.php");
require_once(__DIR__ . "/../vendor/autoload.php");
require_once (__DIR__. "/class/TestDbConnect.php");


switch ($_POST['PHASE']) {
    case "DONE":
        try {
            $form = new AddUserForm("ユーザーを追加");
            $form->doneForm();
        } catch (SmartyException $e) {
            echo "Error:" . $e->getMessage();
        }
        break;
    case "CREATE":
        $instance = TestDbConnect::getInstance();
        $instance->addUser();
        $instance->showPage();
        break;
    default:
        try {
            $form = new AddUserForm("ユーザーを追加");
            $form->showForm();
        } catch (SmartyException $e) {
            echo "Error:" . $e->getMessage();
        }
        break;
}

