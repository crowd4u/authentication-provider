<?php
/* Smarty version 4.3.0, created on 2022-12-19 13:01:03
  from '/var/www/html/smarty/template/test.tpl' */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '4.3.0',
  'unifunc' => 'content_63a0608fd35345_47489059',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    '78bbbfecfbfbec572dc0c859bc7b4bdde55ced42' => 
    array (
      0 => '/var/www/html/smarty/template/test.tpl',
      1 => 1671454861,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
    'file:./parts/head.tpl' => 1,
  ),
),false)) {
function content_63a0608fd35345_47489059 (Smarty_Internal_Template $_smarty_tpl) {
?><html>
<?php $_smarty_tpl->_subTemplateRender('file:./parts/head.tpl', $_smarty_tpl->cache_id, $_smarty_tpl->compile_id, 0, $_smarty_tpl->cache_lifetime, array(), 0, false);
?>
<body class="container p-4">
<h1>工事中だから許して</h1>
<p>とりあえずレギュレーションは満たしてます</p>
<p>進捗はこまめに積もうな</p>
<p>ユーザー情報</p>
<a href="https://github.com/crowd4u/authentication-provider">リポジトリ</a>
<br>
<h2>これは読み出し</h2>
<?php
$_from = $_smarty_tpl->smarty->ext->_foreach->init($_smarty_tpl, $_smarty_tpl->tpl_vars['data']->value, 'item');
$_smarty_tpl->tpl_vars['item']->do_else = true;
if ($_from !== null) foreach ($_from as $_smarty_tpl->tpl_vars['item']->value) {
$_smarty_tpl->tpl_vars['item']->do_else = false;
?>
    <span><?php echo $_smarty_tpl->tpl_vars['item']->value["id"];?>
</span>&nbsp;<span><?php echo $_smarty_tpl->tpl_vars['item']->value["email"];?>
</span>&nbsp;<span><?php echo $_smarty_tpl->tpl_vars['item']->value["user_name"];?>
</span>
    <br>
<?php
}
$_smarty_tpl->smarty->ext->_foreach->restore($_smarty_tpl, 1);?>
<h2>追加</h2>
<form action="../../public/form_confirm.php" method="post">
    <div class="mb-3">
        <label for="exampleInputEmail1" class="form-label">認可トークン</label>
        <input type="text" name="token" class="form-control" id="token">
        <div id="emailHelp" class="form-text">トークンを入力しやがれ</div>
    </div>
    <input type="hidden" name="PHASE" value="CREATE"/>
    <button type="submit" class="btn btn-primary">Create User</button>
</form>

</body>
</html><?php }
}
