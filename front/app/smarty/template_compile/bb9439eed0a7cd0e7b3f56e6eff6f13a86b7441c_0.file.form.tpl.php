<?php
/* Smarty version 4.3.0, created on 2022-12-19 12:06:10
  from '/var/www/html/smarty/template/form.tpl' */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '4.3.0',
  'unifunc' => 'content_63a053b281be85_24144143',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    'bb9439eed0a7cd0e7b3f56e6eff6f13a86b7441c' => 
    array (
      0 => '/var/www/html/smarty/template/form.tpl',
      1 => 1671449852,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
    'file:./parts/head.tpl' => 1,
  ),
),false)) {
function content_63a053b281be85_24144143 (Smarty_Internal_Template $_smarty_tpl) {
?><html>
<?php $_smarty_tpl->_subTemplateRender('file:./parts/head.tpl', $_smarty_tpl->cache_id, $_smarty_tpl->compile_id, 0, $_smarty_tpl->cache_lifetime, array(), 0, false);
?>

<body class="container p-4">

<h1>ログイン</h1>
<p>現在はテスト用に任意のアドレスとパスワードを入力するとトークンが発行されます</p>

<form action="../../public/form_confirm.php" method="post">
    <div class="mb-3">
        <label for="exampleInputEmail1" class="form-label">Email address</label>
        <input type="email" name="id" class="form-control" id="id" aria-describedby="emailHelp">
        <div id="emailHelp" class="form-text">We'll never share your email with anyone else.</div>
    </div>
    <div class="mb-3">
        <label for="exampleInputPassword1" class="form-label">Password</label>
        <input type="password" class="form-control" id="password" name="password">
    </div>
    <input type="hidden" name="PHASE" value="DONE"/>
    <button type="submit" class="btn btn-primary">Login</button>
</form>
</body>
</html><?php }
}
