# passpi
> the password hashing API of your _dreams_

[![GoDoc](https://godoc.org/github.com/t3h2mas/passpi?status.svg)](https://godoc.org/github.com/t3h2mas/passpi)
---

## usage
download the project

`go get github.com/t3h2mas/passpi`

go to the project directory

`cd $GOPATH/src/github.com/t3h2mas/passpi`

build the project

`go build`

run the project (starts on :8080)

`./passpi`

### optional settings
using environment variables

**change the listening address**

`ADDR=':1337' ./passpi`

## run tests on package(s)
`go test -v ./...`