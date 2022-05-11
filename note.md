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

- รันโปรแกรม

```
$ go run main.go
```

- รันตรวจสอบ race condition

```
$ go run -race main.go
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
// coding
go run main.go
```

## Go Testing

```
go test .
go test ./...   // เทสลึกลงไปทุกชั้น
go test -v      // report แบบละเอียดขึ้น
```
