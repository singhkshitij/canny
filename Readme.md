# Canny project
A project aimed at creation of rules to notify users of right time to invest in markets

## Setup
1. Download dependencies
```shell
go mod download
or
go mod ./...
```

2. Start server
```shell
make server
```

3. Start reloading server
   Install dependency (one time)
```shell
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
Start live reload server
```shell
make dev-server
```

4. Update swagger docs
```shell
make swagger
```