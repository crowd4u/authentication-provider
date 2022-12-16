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
curl -c cookie.txt --dump-header -  "localhost:8081/auth?client_id=example-client-id-1&scope=hoge&state=hoge&redirect_url=http://localhost:8081"
```

 post auth check
 
 ```bash
 curl -X POST -H "Content-Type: application/json" -d '{"user_id":"example-user-id-1","password":"password","client_id":"example-client-id-1}' -b cookie.txt -c cookie.txt --dump-header - localhost:8081/auth
 ```

get jwt

```bash
curl -b cookie.txt --dump-header - "localhost:8081/token?client_id=example-client-id-1&grant_type=authorization_code&client_secret=secret&scope=hoge&redirect_uri=http://localhost:8081&state=hoge&code=179e9fd5-44b9-4c88-bc64-0cf80fdd79d6" 
```