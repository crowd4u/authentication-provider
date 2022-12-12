<?php
/* Smarty version 4.3.0, created on 2022-12-12 13:22:00
  from '/var/www/html/smarty/template/form.tpl' */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '4.3.0',
  'unifunc' => 'content_63972af8630215_77577659',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    'bb9439eed0a7cd0e7b3f56e6eff6f13a86b7441c' => 
    array (
      0 => '/var/www/html/smarty/template/form.tpl',
      1 => 1670851314,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
    'file:./parts/head.tpl' => 1,
  ),
),false)) {
function content_63972af8630215_77577659 (Smarty_Internal_Template $_smarty_tpl) {
?><html>
<?php $_smarty_tpl->_subTemplateRender('file:./parts/head.tpl', $_smarty_tpl->cache_id, $_smarty_tpl->compile_id, 0, $_smarty_tpl->cache_lifetime, array(), 0, false);
?>

<body>
<h1>ユーザーを追加</h1>
<form action="../public/form_confirm.php" method="post">
    <lable>Email</lable>
    <input id="id" name="id" type="text" value=<?php echo $_smarty_tpl->tpl_vars['user_id']->value;?>
>
    <lable>Password</lable>
    <input id="password" name="password" type="text" value=<?php echo $_smarty_tpl->tpl_vars['password']->value;?>
>
    <input type="hidden" name="PHASE" value="DONE"/>
    <input class="button_text" type="submit" name="submit" value="ログイン"/>
</form>
</body>
</html><?php }
}
