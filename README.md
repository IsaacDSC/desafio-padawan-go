# DESAFIO PADAWAN GO

### Started only database mysql using docker-compose
```sh
  docker-compose up database -d
```

### Copy schema mysql and create tables in database;
```sh
cat src/infra/database/sqlc/schema.sql
```

### start application;
```sh
go run ./src/cmd/main.go
```

### start tests application;
```sh
go test ./...
```

### Test To API
*http-client => external/http-client/internal-api.http*
