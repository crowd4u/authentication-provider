<?php
require_once(__DIR__ . "/vendor/autoload.php");

/**
 * @param Smarty &$smarty
 * @return void
 */
function SmartyHelper(Smarty &$smarty)
{
    $smarty->setTemplateDir(__DIR__ . "/smarty/template");
    $smarty->setCompileDir(__DIR__ . "/smarty/template_compile");
    $smarty->setCacheDir(__DIR__ . "/smarty/template_c");
}
