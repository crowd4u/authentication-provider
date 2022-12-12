<?php

require_once(__DIR__ . "/add_user.php");

$form = new AddUserForm("ユーザーを追加");
try {
    $form->showForm();
} catch (SmartyException $e) {
    echo "Error" . $e->getMessage();
}
