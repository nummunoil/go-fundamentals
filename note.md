## Installation

download https://go.dev/ -> install -> test command in terminal

```
$ go env
```

command

- ลบแคชที่เกิดจากการ build

```
$ go clean --cache
```

- ลบ dependency ต่างๆ ออกจากเครื่อง

```
$ go clean --modcache
```

## Go Hello

```
mkdir hello
cd hello
// go mod init hello
go mod init github.com/nummunoil/hello
ls // go.mod
code .
touch main.go
```

Run

```
$ go run main.go
```
