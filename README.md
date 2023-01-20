# gobank

use golang to implement a bank service with postgresql as datastore

## config 

use github.com/joho/godotenv/autoload as config loader

which with following setup in the configuration .env

| property name  | description    |
|----------------|----------------|
| DB_USER        | database user  |
| DB_PASSWD      | database password |
| DB_HOST        | database hostname |
| DB_PORT        | database port     |
| DB_NAME        | database name     |
| PORT           | server serve port |

### pg driver

use github.com/lib/pq as driver library

### Makefile

for better run and test 

use makefile for build and test
```makefile
build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test -v ./...
```