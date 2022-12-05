<html>
<body>
<form action="../public/form_confirm.php" method="post">
    <input id="name" name="name" type="text" value={$user_id}>
    <input type="hidden" name="PHASE" value="DONE"/>
    <input class="button_text" type="submit" name="submit" value="表示"/>
</form>
</body>
</html>