<?php
/* Smarty version 4.3.0, created on 2022-12-19 12:06:15
  from '/var/www/html/smarty/template/confirm.tpl' */

/* @var Smarty_Internal_Template $_smarty_tpl */
if ($_smarty_tpl->_decodeProperties($_smarty_tpl, array (
  'version' => '4.3.0',
  'unifunc' => 'content_63a053b70a8151_14448934',
  'has_nocache_code' => false,
  'file_dependency' => 
  array (
    'e3763c01ee3c8ead4e71dd247626bec8e829709a' => 
    array (
      0 => '/var/www/html/smarty/template/confirm.tpl',
      1 => 1671450500,
      2 => 'file',
    ),
  ),
  'includes' => 
  array (
    'file:./parts/head.tpl' => 1,
  ),
),false)) {
function content_63a053b70a8151_14448934 (Smarty_Internal_Template $_smarty_tpl) {
?><html>
<?php $_smarty_tpl->_subTemplateRender('file:./parts/head.tpl', $_smarty_tpl->cache_id, $_smarty_tpl->compile_id, 0, $_smarty_tpl->cache_lifetime, array(), 0, false);
?>

<?php echo '<script'; ?>
>
    function copyToClipboard() {
        // コピー対象をJavaScript上で変数として定義する
        var copyTarget = document.getElementById("copyTarget");

        // コピー対象のテキストを選択する
        copyTarget.select();

        // 選択しているテキストをクリップボードにコピーする
        document.execCommand("Copy");

        // コピーをお知らせする
        alert("コピーできました！ : " + copyTarget.value);
    }
<?php echo '</script'; ?>
>
<body class="container p-4">
<div>
<h2>アクセストークン（有効期限は1日です）</h2>
<div class="d-flex flex-row mb-3">
    <input id="copyTarget" class="form-control" type="text" value=<?php echo $_smarty_tpl->tpl_vars['access_token']->value;?>
 readonly>
    <button onclick="copyToClipboard()" type="submit" class="btn btn-secondary">Copy</button>
</div>
    <a href="../../public/index.php">トップページへ戻る</a>
</div>
</body>
</html><?php }
}
