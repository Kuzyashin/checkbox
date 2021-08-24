# Checkbox test task

### Local run

1) Set up env vars
```
POSTGRES_PASSWORD=checkbox
POSTGRES_USER=checkbox
POSTGRES_DB=checkbox


SERVER_HOST=0.0.0.0
SERVER_PORT=3000
TOMTOM_APIKEY=FOOBAR
DATABASE_USERNAME=checkbox
DATABASE_PASSWORD=checkbox
DATABASE_HOST=postgres
DATABASE_PORT=5432
DATABASE_DBNAME=checkbox
GIN_MODE=release
```
2) `make run_local`

### Docker run
1) Create and fill `.env` file in `build` dir from template `.default.env`
2) `make docker_run` or `make docker_run_detach` for detached run 


### Documentation

Swagger docs can be accessed on http://localhost:${SERVER_PORT} - 
where ${SERVER_PORT} - Env var
