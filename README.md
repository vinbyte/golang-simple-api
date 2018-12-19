1. Make sure Go already installed in your machine and setup your [Go Workspace](https://golang.org/doc/code.html#Workspaces)
2. Install [Govendor](https://github.com/kardianos/govendor) `go get github.com/kardianos/govendor`
3. Go to your $GOPATH/src/. Run `go get github.com/501army/golang-simple-api`
4. Go to inside $GOPATH/src/github.com/501army/golang-simple-api
5. Run `govendor sync` to pull all package needed
6. You're ready to go. Run `go run main.go`. Now your application is running in localhost:2323.
7. Try to access `/v1` a message will appear. Try to access `v1/name` my name will appear :D