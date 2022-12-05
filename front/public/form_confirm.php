<?php
require_once(__DIR__ . "/add_user.php");

$form = new AddUserForm("ユーザーを追加");
switch ($_POST['PHASE']) {
    case "DONE":
        try {
            $form->doneForm();
        } catch (SmartyException $e) {
            echo "Error:" . $e->getMessage();
        }
        break;
    default:
        try {
            $form->showForm();
        } catch (SmartyException $e) {
            echo "Error:" . $e->getMessage();
        }
        break;
}

