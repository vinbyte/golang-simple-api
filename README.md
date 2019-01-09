1. Make sure Go already installed in your machine and setup your [Go Workspace](https://golang.org/doc/code.html#Workspaces)
2. Install [Govendor](https://github.com/kardianos/govendor) `go get github.com/kardianos/govendor`
3. Go to your $GOPATH/src/. Run `go get github.com/501army/golang-simple-api`
4. Go to inside $GOPATH/src/github.com/501army/golang-simple-api
5. Run `govendor sync` to pull all package needed
6. Adjust your config.json file
7. Create your database `golangtes` or change it in config with your own name.
8. You're ready to go. Run `go run main.go`.
9. Now your application is running in your specific port in config (default 2323). Try to access `/v1` a message will appear.

Route list:
* `v1/name` : will produce json include my name :D
* `v1/peoples` : will produce json all people data in table

Depedencies :
- [Golang ORM](https://github.com/jinzhu/gorm)
- [Mysql Driver](https://github.com/go-sql-driver/mysql)
- [Data Faker](https://github.com/bxcodec/faker)