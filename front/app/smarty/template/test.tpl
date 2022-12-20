<html>
{include file='./parts/head.tpl'}
<body class="container p-4">
<h1>工事中だから許して</h1>
<p>とりあえずレギュレーションは満たしてます</p>
<p>進捗はこまめに積もうな</p>
<p>ユーザー情報</p>
<a href="https://github.com/crowd4u/authentication-provider">リポジトリ</a>
<br>
<h2>これは読み出し</h2>
{foreach from=$data item=$item}
    <span>{$item["id"]}</span>&nbsp;<span>{$item["email"]}</span>&nbsp;<span>{$item["user_name"]}</span>
    <br>
{/foreach}
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
</html>