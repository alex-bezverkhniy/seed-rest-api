## Generate docs and spec
```
swag init -d ./cmd/seed-rest-api,./internal/user
```

## Run app
```
make start
```

## Stop app
```
make stop
```

## Run DB only
```
make db
```



### Grand access to marinadb root for dev only

get all grants:

```
SELECT User, Host FROM mysql.user WHERE Host <> 'localhost';
```
or
```
SELECT User, Host FROM mysql.user;
```
change grant
```
GRANT ALL PRIVILEGES ON *.* TO 'root'@'192.168.1.13' IDENTIFIED BY ''; 
```