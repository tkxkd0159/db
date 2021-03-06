# db

## Server
change session secret when it is hijacked. so we can invalidate previous sessions
```sh
kill -15 <app_pid>
```

## MongoDB
```sh
docker run -d --name jsmongo \
-p 27017:27017
-e MONGO_INITDB_ROOT_USERNAME=jsadmin \
-e MONGO_INITDB_ROOT_PASSWORD=secret \
mongo

docker exec -it <container_name> mongo -u "jsadmin" -p "secret"

> db.createUser({ user: "ljs", pwd: "secret", roles: [{role: "readWrite", db: "myapp"}, "readWrite"] })
```

## PostgreSQL
```sh
docker run -d -p 5432:5432 --name jspg -e POSTGRES_PASSWORD=secret postgres
docker exec -it <container_name> psql -U postgres
```

## Redis
```sh
docker run --name jsredis -p 6379:6379 -d redis
docker exec -it jsredis redis-cli
```