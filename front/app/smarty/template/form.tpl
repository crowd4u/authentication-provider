<html>
{include file='./parts/head.tpl'}

<body>
<h1>ユーザーを追加</h1>
<form action="../public/form_confirm.php" method="post">
    <lable>Email</lable>
    <input id="id" name="id" type="text" value={$user_id}>
    <lable>Password</lable>
    <input id="password" name="password" type="text" value={$password}>
    <input type="hidden" name="PHASE" value="DONE"/>
    <input class="button_text" type="submit" name="submit" value="ログイン"/>
</form>
</body>
</html>