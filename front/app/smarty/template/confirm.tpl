<html>
{include file='./parts/head.tpl'}

<script>
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
</script>
<body class="container p-4">
<div>
<h2>アクセストークン（有効期限は1日です）</h2>
<div class="d-flex flex-row mb-3">
    <input id="copyTarget" class="form-control" type="text" value={$access_token} readonly>
    <button onclick="copyToClipboard()" type="submit" class="btn btn-secondary">Copy</button>
</div>
    <a href="../../public/index.php">トップページへ戻る</a>
</div>
</body>
</html>