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
# 1. Docker
docker run -d -p 6666:5432 --name jspg -e PGDATA=/var/lib/postgresql/data/pgdata -v pgdata:/var/lib/postgresql/data -e POSTGRES_USER=ljs -e POSTGRES_PASSWORD=secret postgres
docker exec -it <container_name> psql -U postgres --dbname <dbname>

# 2. Local
psql -h 127.0.0.1 -p 5432 -U <username> -d <dbname> -W

# Import
pg_restore -U ljs -d dvdrental ~/sda/dbsample/dvdrental.tar
```

## Redis
```sh
docker run --name jsredis -p 6379:6379 -d redis
docker exec -it jsredis redis-cli
```