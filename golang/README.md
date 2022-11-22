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

 query
 
 ```bash
 curl -X POST -H "Content-Type: application/json" -d '{"username":"hoge","password":"password","client_id":"example-user-id-1}' localhost:8081/auth
 ```
