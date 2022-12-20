<html>
{include file='./parts/head.tpl'}

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
</html>