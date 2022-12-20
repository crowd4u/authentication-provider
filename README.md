# Authentication Provider

これは2022年のデータベース技術の課題用に制作した認可サーバーになります。

## How to run

### download the source

```bash
git clone git@github.com:crowd4u/authentication-provider.git
cd authentication-provider
```

### how to build

```bash
docker-compose build
docker-compose up -d
```

### how to access

http://localhost/public/index.php でトップページにアクセスすることが出来ます。


## about

golang/ディレクトリ配下は認可サーバーの実装が含まれています。
front/ディレクトリ配下にはテストアプリケーションの実装が含まれています。
