# NextCrowd4u Authenticatoin Provider

## How to use(Local)

How to build and run the containers using docker compose.

```shell
docker compose build
docker compose up -d
```

How to make a new migration file.

```shell
./migrate.sh <filename>
```

How to reset container.

```shell
docker compose down --rmi all --volumes --remove-orphans
```
 
 ## test

 
create session

```bash
curl -c cookie.txt --dump-header -  "localhost:8081/auth?client_id=example-user-id-1&scope=hoge&state=hoge&redirect_url=http://localhost:8081"
```

 post auth check
 
 ```bash
 curl -X POST -H "Content-Type: application/json" -d '{"username":"hoge","password":"password","client_id":"example-user-id-1}' -b cookie.txt -c cookie.txt --dump-header - localhost:8081/auth
 ```

get jwt

```bash
curl -b cookie.txt --dump-header - "localhost:8081/token?client_id=example-user-id-1&grant_type=authorization_code&client_secret=secret&scope=hoge&redirect_uri=http://localhost:8081&state=hoge&code=e1ef3eb4-5c95-4a10-829b-2f04167d1e8d" 
```