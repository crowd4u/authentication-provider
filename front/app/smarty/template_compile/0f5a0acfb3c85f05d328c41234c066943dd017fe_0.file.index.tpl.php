<?php
/* Smarty version 4.3.0, created on 2022-12-12 21:37:28
  from '/usr/local/apache2/test/app/smarty/template/index.tpl' */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '4.3.0',
  'unifunc' => 'content_63972088c46174_32815869',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    '0f5a0acfb3c85f05d328c41234c066943dd017fe' => 
    array (
      0 => '/usr/local/apache2/test/app/smarty/template/index.tpl',
      1 => 1670840351,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
    'file:./parts/head.tpl' => 1,
  ),
),false)) {
function content_63972088c46174_32815869 (Smarty_Internal_Template $_smarty_tpl) {
?><!DOCTYPE html>
<html>
<head>
    <?php $_smarty_tpl->_subTemplateRender('file:./parts/head.tpl', $_smarty_tpl->cache_id, $_smarty_tpl->compile_id, 0, $_smarty_tpl->cache_lifetime, array(), 0, false);
?>

</head>
<body>
Hello World!
</body>
</html>
<?php }
}
