# Migrations

- Make sure you have installed this migration tool https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


## Create migrations

```
migrate create -ext sql -dir src/db/migrations table_name
```

## Runing migrations
```
migrate -database postgresql://postgres:password@localhost:5432/db_go_app\?sslmode=disable -path src/db/migrations up
```
