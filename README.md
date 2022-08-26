# GO JWT

## Project Initialization 

- first we have to make `go.mod` file which is like `package.json` for Node

```sh
    go mod init edr3x/gin-gorm
```

## Now we install necessary packages

```sh
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
go get github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v4
```

## Now to Run the program

```sh
CompileDaemon -command="./go-jwt"
```

> Here `go-jwt` is the module name we set up during the initialization of the project

## `.env` file 

```.env
    PORT=5050
    DB_URL ="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
```
