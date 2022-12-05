<?php

require_once(__DIR__ . "/../vendor/autoload.php");
require_once(__DIR__ . "/../SmartyHelper.php");

class AddUserForm
{
    public $user_id;
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
        $this->smarty->display("form.tpl");
    }

    /**
     * @throws SmartyException
     */
    public function doneForm()
    {
        $this->smarty->display("confirm.tpl");
    }


}