<?php
/* Smarty version 4.3.0, created on 2022-12-19 12:06:08
  from '/var/www/html/smarty/template/index.tpl' */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '4.3.0',
  'unifunc' => 'content_63a053b0a40cc4_18978637',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    '2c757288fe731bb15b4c79326b8ab2c9374dbd4b' => 
    array (
      0 => '/var/www/html/smarty/template/index.tpl',
      1 => 1671450539,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
    'file:./parts/head.tpl' => 1,
  ),
),false)) {
function content_63a053b0a40cc4_18978637 (Smarty_Internal_Template $_smarty_tpl) {
?><!DOCTYPE html>
<html>
<?php $_smarty_tpl->_subTemplateRender('file:./parts/head.tpl', $_smarty_tpl->cache_id, $_smarty_tpl->compile_id, 0, $_smarty_tpl->cache_lifetime, array(), 0, false);
?>

<body class="container p-4">

<h1>Hello World!</h1>
<h2>About</h2>

<p>このページはとあるプロジェクトのために作成されたAuthentication Providerのテストページです。</p>
<a href="./login.php">ログインページへ</a>
<a href="./crud.php">トークンを使ってめちゃくちゃしてみる</a>
</body>
</html>
<?php }
}
