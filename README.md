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

## MySQL
```sh
docker run --name jsmysql -e MYSQL_ROOT_PASSWORD=secret -p 3306:3306 -d mysql:8.0.31
```

## PostgreSQL
```sh
# 1. Docker
docker run --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > postgresql.conf # get the default config
docker run -p 5432:5432 --name jspg --mount type=bind,src="$(pwd)/postgresql.conf",dst=/etc/postgresql/postgresql.conf -e POSTGRES_PASSWORD=secret postgres -c 'config_file=/etc/postgresql/postgresql.conf' # conf file mode : 644
docker exec -it <container_name> psql -U ljs --dbname <dbname>

# 2. Local
psql -h 127.0.0.1 -p 5432 -U <username> -d <dbname>

# Import
pg_restore -U ljs -d dvdrental ~/sda/dbsample/dvdrental.tar
```

## Redis
```sh
docker run --name jsredis -p 6379:6379 -d redis
docker exec -it jsredis redis-cli
```
